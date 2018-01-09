package main

import (
	"fmt"
	"io/ioutil"

	"github.com/howeyc/gopass"
	"github.com/spaco/spo/src/util/encrypt"
)

func main() {
	fmt.Printf("input password\n")
	key, err := gopass.GetPasswd()
	if err != nil {
		fmt.Printf("get pass error:%+v\n", err)
		panic(err)
	}
	//encryptMsg, _ := encrypt.Encrypt(key, "")
	//if err := ioutil.WriteFile("spo_secret.key", []byte(encryptMsg), 0666); err != nil {
	//	fmt.Printf("write error:%+v\n", err)
	//	panic(err)
	//}
	encryptMsg, err := ioutil.ReadFile("spo_secret.key")
	if err != nil {
		fmt.Printf("read error:%+v\n", err)
		panic(err)
	}

	msg, err := encrypt.Decrypt(key, string(encryptMsg))
	if err != nil {
		fmt.Printf("err:%+v\n", err)
		return
	}
	fmt.Println(msg)
}
