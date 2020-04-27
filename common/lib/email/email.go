package email

import (
	"bytes"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
	"strings"
	"time"
)

type AuthMail struct {
	AppName  string
	Host     string
	Port     string
	From     *mail.Address
	To       []*mail.Address
	Username string
	Password string
}

func New(appname, host, port, from, password string, to []string) (*AuthMail, error) {
	conn, err := net.DialTimeout("tcp", host+":"+port, 3*time.Second)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	sender, err := mail.ParseAddress(from)
	if err != nil {
		return nil, err
	}
	var receiverList []*mail.Address

	for _, v := range to {
		receiver, err := mail.ParseAddress(v)
		if err != nil {
			return nil, err
		}
		receiverList = append(receiverList, receiver)
	}
	return &AuthMail{
		AppName:  appname,
		Host:     host,
		Port:     port,
		From:     sender,
		To:       receiverList,
		Username: strings.Split(from, "@")[0],
		Password: password,
	}, nil
}

func (m *AuthMail) Send(title, content string) error {
	auth := smtp.PlainAuth("", m.Username, m.Password, m.Host)

	var toList []string
	for _, v := range m.To {
		toList = append(toList, v.Address)
	}

	message := createMessage(m.AppName, title, content)

	return smtp.SendMail(m.Host+":"+m.Port, auth, m.From.Address, toList, message.Bytes())
}

func createMessage(appname, title, content string) *bytes.Buffer {
	format := "2006-01-02 15:04:05"
	dataTime := time.Unix(time.Now().Unix(), 0).Format(format)
	body := dataTime + " - " + title
	subject := appname
	contents := fmt.Sprintf("Subject: %s\r\n\r\n%s\r\n\r\n%s", subject, body, content)
	message := bytes.NewBufferString(contents)
	return message
}
