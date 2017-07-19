package x5base

// ReqRepHead ...
type ReqRepHead struct {
	Serial   int32
	SeqOrAck int32
}

// NetMessage ...
type NetMessage struct {
	ReqRepHead
	CLSID    int32
	RecvTime int64
}
