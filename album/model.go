package album

import (
	"gorm.io/gorm"
)

type Album struct {
	gorm.Model
	ID     string `gorm:"unique_index"`
	Title  string
	Artist string
	Price  float64
}
