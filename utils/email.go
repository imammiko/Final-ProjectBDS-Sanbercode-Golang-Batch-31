package utils

import (
	"fmt"
	"log"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendingEmail(sendTo string, subject string, message string) {
	CONFIG_SMTP_HOST := Getenv("CONFIG_SMTP_HOST", "smtp.gmail.com")
	i, _ := strconv.Atoi(Getenv("CONFIG_SMTP_PORT", "587"))
	CONFIG_SMTP_PORT := i
	CONFIG_SENDER_NAME := Getenv("CONFIG_SENDER_NAME", "")
	CONFIG_AUTH_EMAIL := Getenv("CONFIG_AUTH_EMAIL", "")
	CONFIG_AUTH_PASSWORD := Getenv("CONFIG_AUTH_PASSWORD", "")
	fmt.Println(CONFIG_SENDER_NAME, "masuk", Getenv("API_SECRET", "gagla"))
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_AUTH_EMAIL)
	mailer.SetHeader("To", sendTo)
	// mailer.SetAddressHeader("Cc", "tralalala@gmail.com", "Tra Lala La")
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", message)
	// mailer.Attach("./sample.png")

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}
