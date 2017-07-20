package main

import (
	"fmt"
	"x5robot/x5base"
	"x5robot/event"
	"reflect"
)

func main() {
	fmt.Println("Hello World!")

	versionReq := &event.CEventVersionRequest{}
	versionReq.CLSID = 151
	versionReq.Serial = 100
	versionReq.SeqOrAck = 0
	versionReq.DefaultServerID = 1
	versionReq.ClientInfo.CarrierType = 0
	versionReq.ClientInfo.ClientMemorySize = 16322
	versionReq.ClientInfo.ClientPlatform = 1
	versionReq.ClientInfo.ClientResVersion = "0.1"
	versionReq.ClientInfo.ClientStorageTotal = 0
	versionReq.ClientInfo.ClientVersion = "0.4.0"
	versionReq.ClientInfo.ClinetStorageFree = 0
	versionReq.ClientInfo.CPUInfo = `Intel(R) Core(TM) i7-4790 CPU @ 3.60GHz-3591-8`
	versionReq.ClientInfo.DeviceInfo = `All Series (ASUS)`
	versionReq.ClientInfo.DeviceSystem = `Windows 7 Service Pack 1 (6.1.7601) 64bit`
	versionReq.ClientInfo.Location = ""
	versionReq.ClientInfo.LoginSource = 2
	versionReq.ClientInfo.NetInfo = "WIFI"
	versionReq.ClientInfo.OpenID = "13093805"
	versionReq.ClientInfo.OpenKey = ""

	v1 := reflect.ValueOf(versionReq)
	fmt.Println("v1 ",v1)
	v2 := v1.Interface()
	fmt.Println("v2 ",v2)
	v3 := reflect.ValueOf(v2)
	fmt.Println("v3 ",v3)
	if v3.Kind() == reflect.Ptr {
		v3 = v3.Elem()
	}
	fmt.Println("v3 ",v3)
	v4 := v3.Interface()
	fmt.Println("v4 ",v4)

	v5 := reflect.TypeOf(v4)
	for i := 0; i < v5.NumField(); i++ {
		fmt.Println(v5.Field(i).Tag)
	}

	allbyte := x5base.Serialize(versionReq)
	fmt.Println(allbyte)

	a := x5base.GetCRC32([]byte{'1'})
	fmt.Println(a)
}
