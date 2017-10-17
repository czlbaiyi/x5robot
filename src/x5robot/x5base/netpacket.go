package x5base

import (
	"bytes"
	"encoding/binary"
)

var headerLength int32 = 4

// NetPacket ...
type NetPacket struct {
	HeadLength int32  `x5tag:"HeadLength"`
	BodyLength int32  `x5tag:"BodyLength"`
	HeadBuffer []byte `x5tag:"headBuffer"`
	BodyBuffer []byte `x5tag:"bodyBuffer"`
}

func GenSendPacket(netMsg interface{}) *NetPacket {
	allbyte := serialize(netMsg)

	packet := &NetPacket{}
	packet.BodyLength = int32(len(allbyte))
	packet.BodyBuffer = allbyte

	packet.HeadLength = headerLength
	lenBuf := bytes.NewBuffer([]byte{})
	binary.Write(lenBuf, binary.LittleEndian, &packet.BodyLength)
	packet.HeadBuffer = lenBuf.Bytes()

	return packet
}

// Decode ...
func (p *NetPacket) Decode() []byte {
	return nil
}

// Encode ...
func (p *NetPacket) Encode() []byte {
	return nil
}
