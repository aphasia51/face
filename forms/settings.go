package forms

type Settings struct {
	Start    int `gorm:"start"    form:"start"    json:"start"    binding:"required"`
	Duration int `gorm:"duration" form:"duration" json:"duration" binding:"required"`
}
