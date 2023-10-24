package email

import (
	"bytes"
	"github.com/dstgo/wilson/assets"
	"github.com/dstgo/wilson/internal/conf"
	"github.com/dstgo/wilson/internal/types/system"
	"github.com/jordan-wright/email"
	"text/template"
)

func NewSender(cfg *conf.AppConf) (Sender, func(), error) {
	emailConf := cfg.EmailConf
	// initialize email pool
	pool, err := email.NewPool(emailConf.Address(), emailConf.MaxPoolSize, emailConf.SmtpAuth())
	if err != nil {
		return Sender{}, nil, err
	}

	sender := Sender{
		pool: pool,
		conf: emailConf,
	}

	fn := func() {
		pool.Close()
	}

	return sender, fn, nil
}

type Sender struct {
	pool *email.Pool
	conf *conf.EmailConf
}

func (e Sender) SendHtmlTemplateMail(mail *email.Email, tmpl string, data map[string]any) error {
	// parse template
	tmp, err := template.ParseFS(assets.Fs, tmpl)
	if err != nil {
		return system.ErrFileSystem.Wrap(err)
	}
	buf := bytes.NewBuffer([]byte{})

	// execute template
	if err = tmp.Execute(buf, data); err != nil {
		return system.ErrProgram.Wrap(err)
	}

	// set content
	mail.HTML = buf.Bytes()

	// send email
	return e.SendMail(mail)
}

func (e Sender) SendTextMail(mail *email.Email, text string) error {
	mail.Text = []byte(text)
	return e.SendMail(mail)
}

func (e Sender) SendMail(mail *email.Email) error {
	if err := e.pool.Send(mail, e.conf.SendTimeout); err != nil {
		return system.ErrNetwork.Wrap(err)
	}
	return nil
}
