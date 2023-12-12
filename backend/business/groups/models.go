package groups

import "time"

type Group struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	Name      string
}
