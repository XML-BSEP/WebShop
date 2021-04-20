package usecase

import (
	"bytes"
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"html/template"
	"net/smtp"
	"os"
	"time"
)

type MailForActivation struct {
	Email	string `json:"email"`
	Token	string `json:"token"`
	Expires	string `json::"expires"`
}

const (
	emailHost = "smtp.gmail.com"
	emailFrom = "duke.strategic@gmail.com"
	emailPassword = "Duke123456"
	emailPort = "587"
	adr = emailHost + ":" + emailPort
)

const (
	tokenExpiresIn = 100000000000000 * 3600 * 5
)

type TokenHandler struct {
	PrivateKey *rsa.PrivateKey
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func NewToken(subject string, email string) (string, error) {

	createdAt := time.Now().Unix()
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, Claims{
		StandardClaims: jwt.StandardClaims{
			Subject:   subject,
			ExpiresAt: createdAt + tokenExpiresIn,
			IssuedAt:  createdAt,
			NotBefore: createdAt,
		},
		Email: email,
	})
	return t.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
}

// Verifikuje token i vraca subjekt i mail
func (e TokenHandler) Verify(token string) (string, string, error) {
	parsed, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return &e.PrivateKey.PublicKey, nil
	})

	if err != nil {
		return "", "", err
	}

	if claims, ok := parsed.Claims.(*Claims); ok && parsed.Valid {
		return claims.StandardClaims.Subject, claims.Email, nil
	}
	return "", "", errors.New("invalid token")
}





func SendMail(subjectMail string, subjectName string, verCode string) error {

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

	/*
	msg := []byte(
		"To: " + subjectMail + "\r\n" +
			"Subject: " + "Email verification by Duke Strategic Techologies" + "\r\n" +
			"Dear " + subjectName + ",\nWe just need to verify your email address before you can access DukeStrategic\n " +
			"\nVerify your email address " + verCode +
			"\n\nThanks! ,\nDuke Strategic Technologies")*/

	to := []string{subjectMail}

	/*
	if err := smtp.SendMail(adr, emailAuth, emailFrom, to, msg); err != nil {
		return  err
	}
	*/

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

func tokenLink(subject string, mail string) string{


	token, err := NewToken(subject, mail)

	if err != nil{
		return fmt.Sprintf("Error token generate")
	}

	return fmt.Sprintf("http://localhost:443/activateMail/%s", token)
}