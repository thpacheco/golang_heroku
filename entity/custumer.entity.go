package entity

type Custumer struct {
	ID        int64  `gorm:"primary_key:auto_increment" json:"-"`
	Name      string `gorm:"type:varchar(100)" json:"-"`
	Email     string `gorm:"type:varchar(200)" json:"-"`
	Telephone string `gorm:"type:varchar(20)" json:"-"`
	DataStart string `gorm:"type:date" json:"-"`
}
