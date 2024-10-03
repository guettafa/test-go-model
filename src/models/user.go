package model

type IUser interface {
	GetLastName() string
	GetAge() int
	SetLastName(newLastName string)
	SetAge(newAge int)
}

type User struct {
	modelImpl
	lastName string
	age int
}

func NewUser(firstName string, lastName string, age int) User {
	u := User{
		lastName: lastName,
		age: age,
	}
	u.name = firstName
	return u
}

func (u *User) GetLastName() string {
	return u.lastName
}

func (u *User) SetLastName(newLastName string) {
	u.lastName = newLastName
}

