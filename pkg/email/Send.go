package email

import "fmt"

type User struct{ Name, Email string }

func SendEmail(u *User) string {
	return fmt.Sprintf("Sending email for %s for email addres %s", u.Name, u.Email)
}
