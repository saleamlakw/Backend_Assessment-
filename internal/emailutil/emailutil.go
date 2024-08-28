package emailutil

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func SendVerificationEmail(recipientEmail string, VerificationToken string) error {
	// Email configuration
	from := os.Getenv("SENDER_EMAIL")
	password := os.Getenv("SENDER_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	subject := "Subject: Account Verification\n"
	mime := "MIME-Version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	url := fmt.Sprintf("http://localhost:8080/verify-email/%v", VerificationToken)
	body := Emailtemplate(url)
	message := []byte(subject + mime + "\n" + body)
	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{recipientEmail}, message)
	if err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}

	return nil

}

// func SendOtpVerificationEmail(recipientEmail string, otp string) error {
// 	// Email configuration
// 	from := os.Getenv("SENDER_EMAIL")
// 	password := os.Getenv("SENDER_PASSWORD")
// 	smtpHost := os.Getenv("SMTP_HOST")
// 	smtpPort := os.Getenv("SMTP_PORT")

// 	subject := "Subject: Account Verification\n"
// 	mime := "MIME-Version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
// 	body := OTPEmailTemplate(otp)
// 	message := []byte(subject + mime + "\n" + body)
// 	auth := smtp.PlainAuth("", from, password, smtpHost)

// 	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{recipientEmail}, message)
// 	if err != nil {
// 		log.Fatalf("Failed to send email: %v", err)
// 	}

// 	return nil

// }