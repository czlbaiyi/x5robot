package robot

import(
	"x5robot/event"
	"x5robot/x5base"
	"fmt"
)

func Init(){
	x5base.Register(&event.CEventVersionRequestResult{},handlerGetVersion)
}

func handlerGetVersion(r *Robot,i interface{}) {
	packet := i.(event.CEventVersionRequestResult)
	fmt.Println(packet)
}