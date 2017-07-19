package x5base

var headerLength int32 = 4

// NetPacket ...
type NetPacket struct {
	HeadLength int32  `x5tag:"HeadLength"`
	BodyLength int32  `x5tag:"BodyLength"`
	HeadBuffer []byte `x5tag:"headBuffer"`
	BodyBuffer []byte `x5tag:"bodyBuffer"`
}

// Decode ...
func (p *NetPacket) Decode() []byte {
	return nil
}

// Encode ...
func (p *NetPacket) Encode() []byte {
	return nil
}
