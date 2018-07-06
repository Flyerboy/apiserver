package model

type User struct {
	Id int	`json:"id"`
	Name string	`json:"name"`
	Email string `json:"email"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) GetUserById(id int) User {
	var user User
	GetDB().First(&user, id)
	return user
}

func (u *User) GetUserList(start, limit int) []User {
	var users []User
	GetDB().Model(User{}).Select("id,name,email").Offset(start).Limit(limit).Find(&users)
	return users
}
