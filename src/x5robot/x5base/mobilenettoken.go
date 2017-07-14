package x5base

//MobileNetToken ...
type MobileNetToken struct{
	Type int16 `x5tag:"type"`
	Serial int16 `x5tag:"serial"`
	Tokenid int64 `x5tag:"token_id"`
}