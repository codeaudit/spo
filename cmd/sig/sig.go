package main

import (
	"fmt"

	"github.com/spaco/spo/src/cipher"
)

func main() {
	pubkeyStr := "036b956f9b94c0e5657dba331425fd1f92b766cab592502e7ec424f7f894a2fc31"

	seckey := cipher.MustSecKeyFromHex("dcd61fd1777d82c279e5d856032df752884ef771d68872161bacb469ae119b4d")
	pubkey := cipher.PubKeyFromSecKey(seckey)
	hash, _ := cipher.SHA256FromHex(pubkeyStr)
	sig := cipher.SignHash(hash, seckey)

	fmt.Println(pubkey.Hex())
	fmt.Println(pubkeyStr)
	fmt.Println(sig.Hex())

}
