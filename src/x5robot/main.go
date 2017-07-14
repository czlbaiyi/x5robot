package main

import (
	"fmt"
	"x5robot/x5base"
	"unsafe"
)

func main(){
	fmt.Println("Hello World!");

	netMessage := &x5base.NetMessage{}
	netMessage.CLSID = 151
	netMessage.Serial = 100
	netMessage.SeqOrAck = 0
	allbyte := netMessage.Serialize()
	fmt.Println(allbyte)
	
	a := x5base.GetCRC32([]byte{'1'})
	_ = a

	i := int(1)
	fmt.Println(unsafe.Sizeof(i)) // 4	
}