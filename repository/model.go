package repository

type User struct {
	UserID   string `gorm:"column:user_id"`
	UserName string `gorm:"column:user_name"`
	Email    string `gorm:"column:email"`
	IsGuest  bool   `gorm:"column:is_guest"`
	UserType string `gorm:"column:user_type"`
}

func (u *User) TableName() string {
	return "tbl_user"
}
