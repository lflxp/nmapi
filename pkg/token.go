package pkg

import (
	"errors"
	"time"

	"github.com/astaxie/beego"
)

var TokenData map[string]*Token

// token过期时间
const expireTime int64 = 3600

func init() {
	TokenData = map[string]*Token{}
}

type Token struct {
	Token    string    `json:"token"`
	Username string    `json:"username"`
	Create   time.Time `json:"create"`
}

func (this *Token) IsExpire() bool {
	var result bool
	this.Create.Unix()
	if time.Now().Unix()-this.Create.Unix() > expireTime {
		result = false
	} else {
		result = true
	}
	return result
}

func HasToken(key string) bool {
	rs := false
	if _, ok := TokenData[key]; ok {
		//存在
		rs = true
	}
	return rs
}

func GetToken(key string) *Token {
	beego.Critical("GETTOKEN", key, TokenData)
	if data, ok := TokenData[key]; ok {
		beego.Critical("111111", data, ok)
		return data
	}
	return nil
}

func DeleteToken(token string) error {
	ok := HasToken(token)
	if ok {
		delete(TokenData, token)
		return nil
	}
	return errors.New("no such token")
}

func NewToken(username string) *Token {
	data := &Token{
		Token:    GetRandomSalt(),
		Username: username,
		Create:   time.Now(),
	}
	TokenData[data.Token] = data
	return data
}
