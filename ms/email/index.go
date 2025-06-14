package email

import (
	"log"

	"gopkg.in/gomail.v2"
)

func SendOrganizationEmail(to string, orgTitle string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "vikram@ethereal.email")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "New Organization Created")
	m.SetBody("text/plain", "Your organization '"+orgTitle+"' was created successfully!")

	d := gomail.NewDialer("smtp.ethereal.email", 587, "al93@ethereal.email", "wYYRv58ZTCEZEsHYeC")

	if err := d.DialAndSend(m); err != nil {
		log.Printf("Email send failed: %v", err)
		return err
	}
	log.Println("Ethereal test email sent successfully")
	return nil
}
