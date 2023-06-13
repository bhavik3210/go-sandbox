package basics

import "fmt"

func DemoInterfaces() {
	PrintHeader("Interfaces")

	// printer is interface type but user is being assigned to that type? How?
	var p printer = user{username: "user", id: 23}
	/*
		why does this assignment work?
		because user is fulfilling the contract of printer interface by providing a member method that has Print() method onto it.

		In other words, user struct is "implementing" the interface via member method so a little less loose in golang than other traditional languages but the idea is similar.
	*/
	fmt.Println(p.Print())

	// to access underlying type of user from interface printer, convert interface back to concrete type
	u := p.(user)
	fmt.Println(u)

	// you can also assign multiple types to same p
	// then use switch v := p.(type) to pull out specific type and take action on it  within in the switch statement
}

type printer interface {
	Print() string
}

func (u user) Print() string {
	return fmt.Sprintf("%v [%v]\n", u.username, u.id)
}
