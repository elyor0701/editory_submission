package helper

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"gopkg.in/mail.v2"
	"io/ioutil"
	"log"
	net_http "net/http"
	"net/smtp"
	"strings"

	"github.com/pkg/errors"
)

const (
	// server we are authorized to send email through
	host     = "smtp.gmail.com"
	hostPort = ":587"

	// user we are authorizing as  old="gehwhgelispgqoql"  new="xkiaqodjfuielsug"
	from                     string = "ucode.udevs.io@gmail.com"
	defaultPassword          string = "xkiaqodjfuielsug"
	verificationTemplateHtml string = `<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Email Verification</title>
			</head>
			<body style="font-family: Arial, sans-serif;">
			
				<div>
					<h1>Hello, {{first_name}}!</h1>
					<p>Welcome to our platform. To complete your registration, please click the link below to verify your email address:</p>
			
					<p><a href="{{verification_link}}" target="_blank">Verify Email</a></p>
			
					<p>If you didn't sign up for an account, please ignore this email.</p>
			
					<p>Best regards,<br>Editorypress Submission System</p>
				</div>
			
			</body>
			</html>
`
	emailVerificationPage string = ""
)

type EmailInfo struct {
	Username string
	Password string
}

type SendMessageByEmail struct {
	From    EmailInfo
	To      string
	Subject string
	Message string
}

func GetGoogleUserInfo(accessToken string) (map[string]interface{}, error) {
	resp, err := net_http.Get("https://www.googleapis.com/oauth2/v3/userinfo?access_token=" + accessToken)
	// fmt.Println("Request to https://www.googleapis.com/oauth2/v3/userinfo?access_token= " + accessToken)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	userInfo := make(map[string]interface{})

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}

func SendEmail(subject, to, link, token string) error {
	message := `
		You can update your password using the following url
   
	   ` + link + "?token=" + token

	auth := smtp.PlainAuth("", from, defaultPassword, host)
	msg := "To: \"" + to + "\" <" + to + ">\n" +
		"From: \"" + from + "\" <" + from + ">\n" +
		"Subject: " + subject + "\n" +
		message + "\n"

	if err := smtp.SendMail(host+hostPort, auth, from, []string{to}, []byte(msg)); err != nil {
		return errors.Wrap(err, "error while sending message to email")
	}

	return nil
}

func SendCodeToEmail(subject, to, code string, email string, password string) error {

	log.Printf("---SendCodeEmail---> email: %s, code: %s", to, code)

	message := `
		Код для подтверждения: ` + code

	// if email == "" {
	// 	email = from
	// }
	// if password == "" {
	// 	password = defaultPassword
	// }

	auth := smtp.PlainAuth("", email, password, host)

	msg := "To: \"" + to + "\" <" + to + ">\n" +
		"From: \"" + email + "\" <" + email + ">\n" +
		"Subject: " + subject + "\n" +
		message + "\n"

	if err := smtp.SendMail(host+hostPort, auth, from, []string{to}, []byte(msg)); err != nil {
		return errors.Wrap(err, "error while sending message to email")
	}

	return nil
}

func GoMessageSend(m SendMessageByEmail) error {
	// Sender's email information
	sender := mail.NewMessage()
	sender.SetHeader("From", m.From.Username)
	sender.SetHeader("To", m.To)
	sender.SetHeader("Subject", m.Subject)
	sender.SetBody("text/html", m.Message)

	// Set up the email server configuration
	d := mail.NewDialer("smtp.gmail.com", 587, m.From.Username, m.From.Password)

	// Uncomment the line below if you are using TLS
	// d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email
	if err := d.DialAndSend(sender); err != nil {
		return err
	}

	return nil
}

func MakeEmailMessage(m map[string]string) string {
	messageTemp := verificationTemplateHtml
	for key, val := range m {
		messageTemp = strings.ReplaceAll(messageTemp,
			fmt.Sprintf("{{%s}}", key),
			val,
		)
	}

	return messageTemp
}
