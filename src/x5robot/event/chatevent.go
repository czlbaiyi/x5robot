package event

type CEventMobileChatResult struct {
	Mret          int32  `x5tag:m_ret`
	NewContent    string `x5tag:new_content`
	ZoneID        int32  `x5tag:zone_id`
	IsPeerOffline int32  `x5tag:is_peer_offline`
}
