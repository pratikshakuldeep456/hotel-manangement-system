package hms

type Guest struct {
	Id      int
	Name    string
	Email   string
	PhoneNo string
}

func NewUser(name, email string) *Guest {
	return &Guest{
		Name:  name,
		Email: email,
	}
}
