package email

import (
	"fmt"
	"github.com/dstgo/wilson/app/conf"
	"github.com/dstgo/wilson/app/core/locale"
	"github.com/dstgo/wilson/app/core/resp"
	"github.com/dstgo/wilson/app/core/vax"
	"github.com/dstgo/wilson/app/data"
	"github.com/dstgo/wilson/app/pkg/httpx"
	"github.com/dstgo/wilson/app/types/code"
	"github.com/dstgo/wilson/app/types/request"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/jordan-wright/email"
	"path"
	"strings"
)

var EmailProviderSet = wire.NewSet(
	NewEmailLogic,
	NewEmailHandler,
)

func NewEmailHandler(cfg *conf.AppConf, emailLogic EmailLogic, datasource *data.DataSource) EmailHandler {
	return EmailHandler{
		EmailLogic: emailLogic,
		cfg:        cfg.EmailConf,
		redis:      datasource.Redis(),
	}
}

type EmailHandler struct {
	EmailLogic EmailLogic
	cfg        *conf.EmailConf
	redis      *redis.Client
}

// SendCodeEmail
//
//	@Summary		auth code email api
//	@Description	auth code email api
//	@Tags			email
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Param			email	query	string	true	"email"
//	@Router			/email/code [GET]
func (e EmailHandler) SendCodeEmail(ctx *gin.Context) {
	emailReq := new(request.Email)
	if err := vax.BindAndResp(ctx, vax.Query(emailReq)); err != nil {
		return
	}

	// generate authcode
	newUUUID := uuid.NewString()
	authcode := strings.ToUpper(strings.Split(newUUUID, "-")[0])

	// store in redis
	if err := e.redis.Set(ctx, fmt.Sprintf("email:code:%s", authcode), emailReq.Email, e.cfg.Expire()).Err(); err != nil {
		resp.Error(ctx).Code(code.DatabaseError).MsgI18n("email.sendFail").Error(resp.DataBaseErr(err)).Send()
		return
	}

	// create new email body
	ee := email.NewEmail()
	ee.From = e.cfg.User
	ee.To = append(ee.To, emailReq.Email)
	ee.Subject = locale.GetWithCtx(ctx, "email.codeSubject")

	// judge language
	language := httpx.GetFirstAcceptLanguage(ctx)
	if language != "zh-CN" {
		language = "en-US"
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
