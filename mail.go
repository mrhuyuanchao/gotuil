package goutil

import (
	"fmt"

	"github.com/go-gomail/gomail"
)

// Mail 邮件
type Mail struct {
	From        *AddressInfo
	To          []*AddressInfo
	Cc          []*AddressInfo
	Bcc         []*AddressInfo
	Subject     string
	Body        string
	ContentType string
	Attach      []string
	Host        string
	Port        int
	UserName    string
	Password    string
}

// AddressInfo 邮件地址
type AddressInfo struct {
	Address string
	Name    string
}

// SendMail 发送邮件
func SendMail(mailContent *Mail) error {
	mail := gomail.NewMessage()
	mail.SetAddressHeader("From", mailContent.From.Address, mailContent.From.Name)
	setAddress(mailContent.To, "To", mail)
	setAddress(mailContent.Cc, "Cc", mail)
	setAddress(mailContent.Bcc, "Bcc", mail)
	mail.SetHeader("Subject", mailContent.Subject)
	mail.SetBody(mailContent.ContentType, mailContent.Body)
	attachFile(mailContent.Attach, mail)
	dialer := gomail.NewPlainDialer(mailContent.Host, mailContent.Port, mailContent.UserName, mailContent.Password)
	if err := dialer.DialAndSend(mail); err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func setAddress(list []*AddressInfo, filed string, mail *gomail.Message) {
	if len(list) <= 0 {
		return
	}
	array := make([]string, 0)
	for _, item := range list {
		str := mail.FormatAddress(item.Address, item.Name)
		array = append(array, str)
	}
	mail.SetHeader(filed, array...)
}

func attachFile(fileList []string, mail *gomail.Message) {
	if len(fileList) <= 0 {
		return
	}
	for _, item := range fileList {
		mail.Attach(item)
	}
}
