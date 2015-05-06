package gnbackend

type User struct {
	Id                 int64
	Uname, Rname, Mail string
}

func NewUser() *User {
	return &User{
		Id:    -1,
		Uname: "",
		Rname: "",
		Mail:  "",
	}
}

func Register(name, rname, mail, pwd string) (int64, error) {
	return -1, nil
}
