package entity

type Pager struct {
	PagerID     uint64 `gorm:"primary_key:auto_increment" json:"id"`
	PagerName   string `gorm:"type:varchar(255)" json:"pagername"`
	PagerStatus bool   `gorm:"type:varchar(255)" json:"xstatus"`
	UserID      uint64 `gorm:"not null" json:"-"`
	User        User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	SSID        string `gorm:"type:varchar(255)" json:"ssid"`
	SSIDPass    string `gorm:"type:varchar(255)" json:"ssidPass"`
}
