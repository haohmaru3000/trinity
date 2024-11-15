package user_model

type User struct {
	Id       int    `json:"id" gorm:"column:id;"`
	Email    string `json:"email" gorm:"column:email;"`
	FullName string `json:"full_name" gorm:"column:full_name;"`
}
