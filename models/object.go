package models

import (
	"strings"
	"math"
)

type SURL struct {
	Id        int64 `json:"-"`
	ShortCode string `orm:"-"`
	LongURL   string `orm:"size(1024)"`
}

func(surl *SURL)Save(){
	ormdb := NewOrm()
	ormdb.Insert(surl)
	surl.ShortCode = ShortCode(surl.Id)
}

func(surl *SURL)Get(){
	surl.Id = CodeId(surl.ShortCode)
	shortcode := surl.ShortCode
	ormdb := NewOrm()
	ormdb.Read(surl)
	surl.ShortCode = shortcode
}

func ShortCode(id int64)(code string){
	letters := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	weight := int64(len(letters))
	for id > 0{
		quotient := id / weight
		remainder := id - quotient * weight
		letter := string(letters[remainder])
		id = quotient
		code = letter + code
	}
	return
}

func CodeId(code string)(id int64){
	letters := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	weight := len(letters)
	for bit,letter := range(code){
		id += int64(strings.Index(letters,string(letter)))*int64(math.Pow(float64(weight),float64((len(code)-bit-1))))
	}
	return
}