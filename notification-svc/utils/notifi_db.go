package utils

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
	"github.com/jordan-wright/email"
)

func init() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}
}

func EmailUser(to string, orderID string, amount float64, status string) error {

	fmt.Println("SMTP_EMAIL:", os.Getenv("SMTP_EMAIL"))
	fmt.Println("SMTP_PASSWORD:", os.Getenv("SMTP_PASSWORD"))
	fmt.Println("SMTP_HOST:", os.Getenv("SMTP_HOST"))
	fmt.Println("SMTP_PORT:", os.Getenv("SMTP_PORT"))

	e := email.NewEmail()
	e.From = fmt.Sprintf("Order Notifier <%s>", os.Getenv("SMTP_EMAIL"))
	e.To = []string{to}
	e.Subject = "Order Update: " + orderID
	e.Text = []byte(fmt.Sprintf(
		"Hi there,\n\nYour order (%s) worth $%.2f has a new status: %s.\n\nThank you!",
		orderID, amount, status))

	auth := smtp.PlainAuth("", os.Getenv("SMTP_EMAIL"), os.Getenv("SMTP_PASSWORD"), os.Getenv("SMTP_HOST"))
	err := e.Send(os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"), auth)
	if err != nil {
		log.Println("Email sending failed:", err)
		return err
	}

	log.Println("Email sent to:", to)
	return nil
}
