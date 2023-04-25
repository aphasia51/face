package forms

type UserForm struct {
	UserName string `gorm:"username" form:"username" json:"username" binding:"required"`
	PassWord string `gorm:"password" form:"password" json:"password" binding:"required,max=20"`
}
