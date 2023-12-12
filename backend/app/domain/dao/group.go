package dao

import "time"

type Group struct {
	ID        int       `gorm:"column:id; primary_key; not null" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	CreatedAt time.Time `gorm:"->:false;column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"->:false;column:updated_at" json:"-"`
}
