package hms

type Guest struct {
	Id      int
	Name    string
	Email   string
	PhoneNo string
}

func NewGuest(name, email, phone string) *Guest {
	return &Guest{
		Id:    GenerateId(),
		Name:  name,
		Email: email,

		PhoneNo: phone,
	}
}
