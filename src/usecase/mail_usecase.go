package usecase

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"html/template"
	"log"
	"net/smtp"
)

type MailForActivation struct {
	Email	string `json:"email"`
	Token	string `json:"token"`
	Expires	string `json::"expires"`
}



func init_viper() {
	viper.SetConfigFile(`src/configurations/mailconfig.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}





func SendMail(subjectMail string, subjectName string, verCode string) error {

	init_viper()
	emailFrom := viper.GetString(`emailFrom`)
	emailPassword := viper.GetString(`emailPassword`)
	emailHost := viper.GetString(`emailHost`)
	adr := viper.GetString(`emailHost`) + ":" + viper.GetString(`emailPort`)

	emailAuth := smtp.PlainAuth("", emailFrom, emailPassword, emailHost)

	t, _  := template.ParseFiles("src/assets/mail_template/template.html")
	var body bytes.Buffer
	headers := "MIME-version: 1.0;\nContent-Type: text/html;"
	body.Write([]byte(fmt.Sprintf("Subject: Email verification by Duke Strategic Techologies\n%s\n\n", headers)))

	t.Execute(&body, struct {
		UserName string
		Code string
	}{
		UserName: subjectName,
		Code: verCode,
	})

	to := []string{subjectMail}


	if err := smtp.SendMail(adr, emailAuth, emailFrom, to, body.Bytes()); err != nil {
		return  err
	}



	return nil
}

func SendRestartPasswordMail(subjectMail string, verCode string) error {

	init_viper()
	emailFrom := viper.GetString(`emailFrom`)
	emailPassword := viper.GetString(`emailPassword`)
	emailHost := viper.GetString(`emailHost`)
	adr := viper.GetString(`emailHost`) + ":" + viper.GetString(`emailPort`)

	emailAuth := smtp.PlainAuth("", emailFrom, emailPassword, emailHost)

	t, _  := template.ParseFiles("src/assets/mail_template/template_reset.html")
	var body bytes.Buffer
	headers := "MIME-version: 1.0;\nContent-Type: text/html;"
	body.Write([]byte(fmt.Sprintf("Subject: Code for password reset by Duke Strategic Techologies\n%s\n\n", headers)))

	t.Execute(&body, struct {
		Code string
	}{
		Code: verCode,
	})

	to := []string{subjectMail}


	if err := smtp.SendMail(adr, emailAuth, emailFrom, to, body.Bytes()); err != nil {
		return  err
	}



	return nil
}




func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	body := buf.String()
	return body, nil
}
