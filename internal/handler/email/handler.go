package email

import (
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/pkg/httpx"
	"github.com/dstgo/wilson/internal/pkg/locale"
	resp2 "github.com/dstgo/wilson/internal/pkg/resp"
	"github.com/dstgo/wilson/internal/pkg/valid"
	email2 "github.com/dstgo/wilson/internal/types/api/email"
	"github.com/dstgo/wilson/internal/types/code"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/jordan-wright/email"
	"path"
	"strings"
)

var EmailProviderSet = wire.NewSet(
	NewSender,
	NewEmailCodeCache,
	NewEmailHandler,
)

func NewEmailHandler(cfg *conf.AppConf, emailLogic Sender, codeCache CodeCache) EmailHandler {
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
	codeCache    CodeCache
}

// SendCodeEmail
//
//	@Summary		SendCodeEmail
//	@Description	auth code email api
//	@Tags			email
//	@Accept			json
//	@Produce		json
//	@Param			email	query	string	true	"email"
//	@Router			/email/code [GET]
func (e EmailHandler) SendCodeEmail(ctx *gin.Context) {
	emailReq := new(email2.Email)
	if err := valid.BindAndResp(ctx, valid.Query(emailReq)); err != nil {
		return
	}

	// generate authcode
	newUUUID := uuid.NewString()
	authcode := strings.ToUpper(strings.Split(newUUUID, "-")[0])

	// store in redis
	if err := e.codeCache.Set(ctx, authcode, emailReq.Email, e.cfg.Expire()); err != nil {
		resp2.InternalErr(ctx).
			Code(code.DatabaseError).
			MsgI18n("email.sendFail").
			Error(resp2.DataBaseErr(err)).Send()
		return
	}

	// create new email body
	ee := email.NewEmail()
	ee.From = e.cfg.User
	ee.To = append(ee.To, emailReq.Email)
	ee.Subject = locale.GetWithCtx(ctx, "email.codeSubject")

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
		resp2.Fail(ctx).Code(code.EmailSendFailed).MsgI18n("email.sendFail").Error(err).Send()
	} else {
		resp2.Ok(ctx).Code(code.EmailSendOk).MsgI18n("email.sendOk").Send()
	}
}
