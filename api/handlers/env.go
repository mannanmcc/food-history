package handlers

import "github.com/jinzhu/gorm"

//Env - holds the environment data
type Env struct {
	Db *gorm.DB
}
