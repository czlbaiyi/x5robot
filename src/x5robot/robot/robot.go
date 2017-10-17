package robot

import (
	"net"
	"x5robot/x5base"
	)

type Robot struct {
	NetPacket          chan interface{}
	conn               net.Conn
	msgbuf             []byte
}

func (r *Robot) Init(){
	x5base.Re
}

func (r *Robot) Update(){
	select {
		case np := <-r.NetPacket:
		
	}
}