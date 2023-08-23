package systemLogic

import (
	"bytes"
	"github.com/dstgo/wilson/app/conf"
	"github.com/dstgo/wilson/app/core/resp"
	"github.com/dstgo/wilson/assets"
	"github.com/jordan-wright/email"
	"text/template"
)

func NewEmailLogic(cfg *conf.AppConf, pool *email.Pool) EmailLogic {
	return EmailLogic{
		pool: pool,
		conf: cfg.EmailConf,
	}
}

type EmailLogic struct {
	pool *email.Pool
	conf *conf.EmailConf
}

func (e EmailLogic) SendHtmlTemplateMail(mail *email.Email, tmpl string, data map[string]any) error {
	// parse template
	tmp, err := template.ParseFS(assets.Fs, tmpl)
	if err != nil {
		return resp.FileSystemErr(err)
	}
	buf := bytes.NewBuffer([]byte{})

	// execute template
	if err = tmp.Execute(buf, data); err != nil {
		return resp.ProgramErr(err)
	}

	// set content
	mail.HTML = buf.Bytes()

	// send email
	return e.SendMail(mail)
}

func (e EmailLogic) SendTextMail(mail *email.Email, text string) error {
	mail.Text = []byte(text)
	return e.SendMail(mail)
}

func (e EmailLogic) SendMail(mail *email.Email) error {
	if err := e.pool.Send(mail, e.conf.SendTimeout); err != nil {
		return resp.NetworkErr(err)
	}
	return nil
}
