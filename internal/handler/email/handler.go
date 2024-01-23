package email

import (
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/core/resp"
	"github.com/dstgo/wilson/internal/data/cache"
	"github.com/dstgo/wilson/internal/pkg/locale"
	"github.com/dstgo/wilson/internal/types/code"
	emailType "github.com/dstgo/wilson/internal/types/email"
	"github.com/dstgo/wilson/pkg/ginx/bind"
	"github.com/dstgo/wilson/pkg/ginx/httpx"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/jordan-wright/email"
	"path"
	"strings"
)

var EmailProviderSet = wire.NewSet(
	NewSender,
	cache.EmailCodeCacheProvider,
	NewEmailHandler,
)

func NewEmailHandler(cfg *conf.WilsonConf, emailLogic Sender, codeCache cache.RedisEmailCodeCache) EmailHandler {
	return EmailHandler{
		EmailLogic:   emailLogic,
		cfg:          cfg.EmailConf,
		fallbackLang: cfg.LocaleConf.Locale,
		codeCache:    codeCache,
	}
}

type EmailHandler struct {
	EmailLogic   Sender
	cfg          *conf.EmailConf
	fallbackLang string
	codeCache    cache.RedisEmailCodeCache
}

// SendCodeEmail
// @Summary      SendCodeEmail
// @Description  auth code email api
// @Tags         email
// @Accept       json
// @Produce      json
// @Param        email	query	string	true	"email"
// @Success      200  {object}  types.Response
// @Router       /email/code [GET]
func (e EmailHandler) SendCodeEmail(ctx *gin.Context) {
	emailReq := new(emailType.SendCodeEmailOption)
	if err := bind.Binds(ctx, bind.Query(emailReq)); err != nil {
		return
	}

	// generate authcode
	newUUUID := uuid.NewString()
	authcode := strings.ToUpper(strings.Split(newUUUID, "-")[0])

	// store in redis
	if err := e.codeCache.Set(ctx, authcode, emailReq.Email, e.cfg.Expire()); err != nil {
		resp.InternalFailed(ctx).Error(emailType.ErrSendFailed).Send()
		return
	}

	// create new email body
	ee := email.NewEmail()
	ee.From = e.cfg.User
	ee.To = append(ee.To, emailReq.Email)
	ee.Subject = locale.GetWithLang(httpx.GetFirstAcceptLanguage(ctx), "email.code.subject")

	// judge language
	language := httpx.GetFirstAcceptLanguage(ctx)
	if language == "" {
		language = e.fallbackLang
	}

	codeTmpl := path.Join("email", language, "code.html")

	// send html email
	err := e.EmailLogic.SendHtmlTemplateMail(ee, codeTmpl, map[string]any{
		"to":     emailReq.Email,
		"code":   authcode,
		"expire": e.cfg.Exp,
	})

	if err != nil {
		resp.Fail(ctx).Code(code.EmailSendFailed).MsgI18n("email.sendFail").Error(err).Send()
	} else {
		resp.Ok(ctx).Code(code.EmailSendOk).MsgI18n("email.sendOk").Send()
	}
}
