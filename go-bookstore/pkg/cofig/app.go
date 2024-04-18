package cofig

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzu/gorm"
)

var(
	db * gorm.DB

)

func Connect(){
	d,err := gorm.Open("mysql","kumar")
}