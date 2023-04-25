package forms

type Record struct {
	Id       int    `gorm:"id"        form:"id"        json:"id"        binding:"required"`
	UserName string `gorm:"user_name" form:"user_name" json:"user_name" binding:"required"`
	Start    int    `gorm:"start"     form:"start"     json:"start"     binding:"required"`
	Duration int    `gorm:"duration"  form:"duration"  json:"duration"  binding:"required"`
	Date     string `gorm:"date"      form:"date"      json:"date"      binding:"required"`
}
