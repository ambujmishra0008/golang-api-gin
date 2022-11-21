package auth

import (
	b64 "encoding/base64"
	"encoding/json"
	"log"
	"os"
)

var USER_LIST map[string]string

func init() {
	file, err := os.Open("auth/users.json")
	if err != nil {
		log.Println("error in opening users file", err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&USER_LIST)
	if err != nil {
		log.Println("error in decoding user file", err)
	}
}

func Validate(user string, password string, hasAuth bool) bool {
	if !hasAuth || user == "" || password == "" {
		return false
	}
	//user = b64.StdEncoding.EncodeToString([]byte(user))
	password = b64.StdEncoding.EncodeToString([]byte(password))
	if USER_LIST[user] == password {
		return true
	}
	return false

}

// var normal_p = "Ambuj"
// 	encode := b64.StdEncoding.EncodeToString([]byte(normal_p))
// 	decode, err := b64.StdEncoding.DecodeString(encode)
// 	fmt.Println(string(decode), err)
