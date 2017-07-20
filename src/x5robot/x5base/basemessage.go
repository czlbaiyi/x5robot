package x5base

//////////////////CEventMobileReqRepBase//////////////////
type CEventMobileReqRepBase struct {
	NetMessage `x5tag:"inherit"`
}

//////////////////CEventMobileRequest//////////////////
type CEventMobileRequest struct {
	CEventMobileReqRepBase `x5tag:"inherit"`
}

func (r *CEventMobileRequest) IsSilent() bool {
	return false
}

//////////////////CEventIngame//////////////////
type CEventIngame struct {
	NetMessage `x5tag:"inherit"`
}

func (r *CEventIngame) IsSilent() bool {
	return true
}

func (r *CEventIngame) IsNeedResend() bool {
	return false
}

//////////////////ServerPushMessage//////////////////
type ServerPushMessage struct {
	NetMessage `x5tag:"inherit"`
}
