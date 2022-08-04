package main

import (
	"log"

	"github.com/bamchoh/pasori"
)

var (
	VID uint16 = 0x054C // SONY
	PID uint16 = 0x06C3 // RC-S380
)

func main() {
	idm, err := pasori.GetID(VID, PID)
	if err != nil {
		panic(err)
	}
	log.Println(idm)
	//return idm
}
