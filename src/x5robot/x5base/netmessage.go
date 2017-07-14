package x5base

import(
	"bytes"
	"encoding/binary"
)

// ReqRepHead ...
type ReqRepHead struct {
	Serial int32 `x5tag:"serial"`
	SeqOrAck int32 `x5tag:"seq_or_ack"`
}

// NetMessage ...
type NetMessage struct{
	ReqRepHead
	CLSID int32 `x5tag:"CLSID"`
	RecvTime int64 `x5tag:"RecvTime"`
}

// Serialize ...
func (p *NetMessage) Serialize() []byte{
	//
	//1 p.CLSID
	allBuffs := bytes.NewBuffer([]byte{})
	binary.Write(allBuffs, binary.LittleEndian, p.CLSID)

	//2 p.Serial
	binary.Write(allBuffs, binary.LittleEndian, p.Serial)

	//3 p.SeqOrAck
	binary.Write(allBuffs, binary.LittleEndian, p.SeqOrAck)

	//4. mes

	return allBuffs.Bytes()
}