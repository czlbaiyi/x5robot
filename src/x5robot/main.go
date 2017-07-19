package main

import (
	"fmt"
	"x5robot/x5base"
	"x5robot/event"
)

func main() {
	fmt.Println("Hello World!")

	versionReq := &event.CEventVersionRequest{}
	versionReq.CLSID = 151
	versionReq.Serial = 100
	versionReq.SeqOrAck = 0
	allbyte := x5base.Serialize(versionReq)
	fmt.Println(allbyte)

	a := x5base.GetCRC32([]byte{'1'})
	fmt.Println(a)
}
