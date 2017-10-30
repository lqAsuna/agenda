package entity

//User .
type User struct {
	Name     string
	Password string
	Email    string
	Phone    string
}

//GetUser
func GetUser(N string, pass string, E string, P string) User {
	us := User()
	us.Name = N
	us.Password = pass
	us.Email = E
	us.Phone = P
	return us
}

//GetName
func (u *User) GetName() string {
	return u.Name
}

//SetName
func (u *User) SetName(N string) {
	u.Name = N
}

//GetPassword
func (u *User) GetPassword() string {
	return u.Password
}

//SetPassword
func (u *User) SetPassword(P string) {
	u.Password = P
}

//GetEmail
func (u *User) GetEmail() string {
	return u.Email
}

//SetEmail
func (u *User) SetEmail (E string) {
	u.Email = E
}

//GetPhone
func (u *User) GetPhone() string {
	return u.Phone
}

//SetName
func (u *User) SetPhone(Ph string) {
	u.Phone = Ph
}
