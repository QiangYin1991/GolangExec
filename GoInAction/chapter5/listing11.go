package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending user email to %s <%s>\n",
		u.name,
		u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	bill := user{name: "bill", email: "bill@email.com"}
	bill.notify()

	lisa := &user{name: "lisa", email: "lisa@email.com"}
	lisa.notify()

	bill.changeEmail("bill@newDomain.com")
	lisa.changeEmail("lisa@newDomain.com")

	bill.notify()
	lisa.notify()
}
