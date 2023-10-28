package authen

import (
	"context"
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/data/cache"
	"github.com/dstgo/wilson/internal/pkg/jwtx"
	"github.com/dstgo/wilson/internal/types/system"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

var (
	ErrTokenExpired          = errors.New("access token expired")
	ErrTokenNeedRefreshed    = errors.New("token need to be refreshed")
	ErrInvalidTokenPayload   = errors.New("invalid token payload")
	ErrInvalidToken          = errors.New("invalid token")
	ErrTokenMisMatch         = errors.New("token mismatch")
	ErrTokenExpirationExceed = errors.New("token expiration exceeded")
)

// Parser
// The Parser should verify if the request has been authenticated.
type Parser interface {
	Parse(ctx context.Context, token string) (Token, error)
}

// Issuer
// The Issuer should issue a new jwt token and return the token info
type Issuer interface {
	Issue(ctx context.Context, payload UserPayload) (Token, error)
}

// Refresher
// if refresh-token expired , Refresher will not refresh token
// else if access-token has expired after delay duration, Refresher will not refresh token
// else if access-token has expired before delay duration, Refresher will issue a new access-token
// else if access-token has not expired, Refresher will renewal the access-token expired time
type Refresher interface {
	Refresh(ctx context.Context, accessToken string, refreshToken string) (Token, error)
}

type Token struct {
	Access  Jwt
	Refresh Jwt
}

type Jwt struct {
	Tk      jwtx.Jwt
	Payload UserClaims
}

// UserPayload
// basic user info
type UserPayload struct {
	Username   string `json:"username"`
	UUID       string `json:"uuid"`
	Persistent bool   `json:"persistent"`
}

type UserClaims struct {
	UserPayload
	jwt.RegisteredClaims
}

func NewRefreshTokenAuthor(cfg *conf.JwtConf, accessCache cache.TokenCache, refreshCache cache.TokenCache) RefreshTokenAuthor {
	return RefreshTokenAuthor{
		accessCache:     accessCache,
		refreshCache:    refreshCache,
		delay:           cfg.Delay,
		accessTokenExp:  cfg.Exp,
		refreshTokenExp: cfg.RExp,
		sig:             cfg.Sig,
		isu:             cfg.Isu,
		method:          jwt.SigningMethodHS512,
	}
}

type RefreshTokenAuthor struct {
	accessCache  cache.TokenCache
	refreshCache cache.TokenCache

	delay           time.Duration
	accessTokenExp  time.Duration
	refreshTokenExp time.Duration

	sig    string
	isu    string
	method jwt.SigningMethod
}

func (r RefreshTokenAuthor) Issue(ctx context.Context, payload UserPayload) (Token, error) {
	var (
		tk          Token
		issueAt     = jwt.NewNumericDate(time.Now())
		expiredAt   = jwt.NewNumericDate(time.Now().Add(r.accessTokenExp))
		reExpiredAt = jwt.NewNumericDate(time.Now().Add(r.refreshTokenExp))
	)

	// access token id same as refresh token id
	tokenId := uuid.NewString()

	accClaims := UserClaims{
		UserPayload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    r.isu,
			Subject:   payload.UUID,
			ExpiresAt: expiredAt,
			IssuedAt:  issueAt,
			ID:        tokenId,
		},
	}

	// issue token
	accJwt, err := jwtx.NewJwt(r.sig, r.method, accClaims)
	if err != nil {
		return tk, err
	}
	tk.Access = Jwt{
		Tk:      accJwt,
		Payload: accClaims,
	}

	ttl := r.accessTokenExp
	if payload.Persistent {
		ttl += r.delay
	}

	// cache token
	// access token cache ttl = expiration duration + delay duration
	err = r.accessCache.Set(ctx, tokenId, tokenId, ttl)
	if err != nil {
		return tk, system.ErrDatabase.Wrap(err)
	}

	if !payload.Persistent {
		return tk, nil
	}

	reClaims := accClaims
	reClaims.ExpiresAt = reExpiredAt
	reJwt, err := jwtx.NewJwt(r.sig, r.method, reClaims)
	if err != nil {
		return tk, err
	}
	tk.Refresh = Jwt{
		Tk:      reJwt,
		Payload: reClaims,
	}

	err = r.refreshCache.Set(ctx, tokenId, tokenId, r.refreshTokenExp)
	if err != nil {
		return tk, system.ErrProgram.Wrap(err)
	}

	return tk, nil
}

func (r RefreshTokenAuthor) Parse(ctx context.Context, token string) (Token, error) {
	access, err := r.parseAccess(ctx, token)
	return Token{Access: access}, err
}

func (r RefreshTokenAuthor) parseAccess(ctx context.Context, token string) (Jwt, error) {
	var (
		tk         Jwt
		secret     = r.sig
		method     = r.method
		userClaims UserClaims
	)

	// try to parse token
	parsedJwt, err := jwtx.ParseJwt(token, secret, method, &UserClaims{})
	if claims, e := parsedJwt.Claims.(*UserClaims); e {
		userClaims = *claims
	} else {
		return tk, ErrInvalidTokenPayload
	}

	// check if is expired
	var expired bool
	if errors.Is(err, jwt.ErrTokenExpired) {
		expired = true
	} else if err != nil {
		return tk, err
	}

	if expired && !userClaims.Persistent {
		return tk, ErrTokenExpired
	}

	tk = Jwt{Tk: parsedJwt, Payload: userClaims}

	// check if is in cache
	_, e, err := r.accessCache.Get(ctx, userClaims.ID)
	if err == nil && !e {
		return tk, ErrTokenExpired
		// if token is expired, but still in cache which means that token need tobe refreshed
	} else if err == nil && e && expired {
		return tk, ErrTokenNeedRefreshed
	} else if err != nil {
		return tk, err
	}

	return tk, nil
}

func (r RefreshTokenAuthor) parseRefresh(ctx context.Context, token string) (Jwt, error) {
	var (
		tk         Jwt
		secret     = r.sig
		method     = r.method
		userClaims UserClaims
	)

	// try to parse token
	parsedJwt, err := jwtx.ParseJwt(token, secret, method, &UserClaims{})
	if claims, e := parsedJwt.Claims.(*UserClaims); e {
		userClaims = *claims
	} else {
		return tk, ErrInvalidTokenPayload
	}

	// check if is expired
	if errors.Is(err, jwt.ErrTokenExpired) {
		return tk, ErrTokenExpired
	} else if err != nil {
		return tk, err
	}

	tk.Tk = parsedJwt
	tk.Payload = userClaims

	// find in cache
	_, e, err := r.refreshCache.Get(ctx, userClaims.ID)
	if err == nil && !e {
		return tk, ErrTokenExpired
	} else if err != nil {
		return tk, system.ErrDatabase.Wrap(err)
	}

	return tk, nil
}

func (r RefreshTokenAuthor) Refresh(ctx context.Context, accessToken string, refreshToken string) (Token, error) {
	var (
		token Token
		renew bool
	)

	access, err := r.parseAccess(ctx, accessToken)
	// if no need to refresh, just return error
	if err != nil && !errors.Is(err, ErrTokenNeedRefreshed) {
		return token, err
		// if access token has not expired, renewal expiration time
	} else if err == nil {
		renew = true
	}

	refresh, err := r.parseRefresh(ctx, refreshToken)
	if err != nil {
		return token, err
	}

	// access token id should same as refresh token id
	if access.Payload.ID != refresh.Payload.ID {
		return token, ErrTokenMisMatch
	}

	var ttl time.Duration

	if renew {
		expiredAt := access.Payload.ExpiresAt.Time
		duration := expiredAt.Sub(time.Now())
		// renew 1/2 per-time
		increment := r.accessTokenExp / 2
		// renew cache ttl
		oldTTL, err := r.accessCache.TTL(ctx, access.Payload.ID)
		if err != nil {
			return token, system.ErrDatabase.Wrap(err)
		}

		// prevent access expiration-time becoming too large
		if duration < r.accessTokenExp*2 {
			access.Payload.ExpiresAt = jwt.NewNumericDate(access.Payload.ExpiresAt.Add(increment))
			ttl = oldTTL + increment
		} else {
			return token, ErrTokenExpirationExceed
		}
	} else {
		access.Payload.IssuedAt = jwt.NewNumericDate(time.Now())
		access.Payload.ExpiresAt = jwt.NewNumericDate(time.Now().Add(r.accessTokenExp))
		ttl = r.accessTokenExp + r.delay
	}

	// allocate a new token
	newJwt, err := jwtx.NewJwt(r.sig, r.method, access.Payload)
	if err != nil {
		return Token{}, err
	}

	//  store in cache
	err = r.accessCache.Set(ctx, access.Payload.ID, access.Payload.ID, ttl)
	if err != nil {
		return token, system.ErrDatabase.Wrap(err)
	}

	token.Access = Jwt{Tk: newJwt, Payload: access.Payload}
	token.Refresh = refresh
	token.Refresh.Tk.SignedJwt = refreshToken

	return token, nil
}
