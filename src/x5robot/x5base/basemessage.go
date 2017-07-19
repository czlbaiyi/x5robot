package x5base

//////////////////CEventMobileReqRepBase//////////////////
type CEventMobileReqRepBase struct {
	NetMessage
}

//////////////////CEventMobileRequest//////////////////
type CEventMobileRequest struct {
	CEventMobileReqRepBase
}

func (r *CEventMobileRequest) IsSilent() bool {
	return false
}

//////////////////CEventIngame//////////////////
type CEventIngame struct {
	NetMessage
}

func (r *CEventIngame) IsSilent() bool {
	return true
}

func (r *CEventIngame) IsNeedResend() bool {
	return false
}

//////////////////ServerPushMessage//////////////////
type ServerPushMessage struct {
	NetMessage
}
