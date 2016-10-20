package models

import (
	// "fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// "math/rand"
)

type Project struct {
	gorm.Model
	Name   string `json:"name"`
	UserID uint   `gorm:"index"`
	Token  string `gorm:"index"`
}

func (p *Project) GenerateToken() {

}

// Something I found on StackOverflow... should look into actual RNG in Go
// func randSeq(n int) string {
// 	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
//     b := make([]rune, n)
//     for i := range b {
//         b[i] = letters[rand.Intn(len(letters))]
//     }
//     return string(b)
// }
