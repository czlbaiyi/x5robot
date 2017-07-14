package event

import(
	"bytes"
	"encoding/binary"
)

// MobileClientInfo ...
type MobileClientInfo struct {
	ClientVersion string `x5tag:"client_version"`
	ClientPlatform int32 `x5tag:"client_platform"`
	LoginSource int32 `x5tag:"login_source"`
	DeviceInfo string `x5tag:"device_info"`
	OpenID string `x5tag:"open_id"`
	OpenKey string `x5tag:"open_key"`
	Pf string  `x5tag:"pf"`
	NetInfo string `x5tag:"net_info"`
	DeviceSystem string `x5tag:"device_system"`
	Location string `x5tag:"location"`
	ClientMemorySize int32 `x5tag:"client_memory_size"`
	ClientStorageTotal int64 `x5tag:"client_storage_total"`
	ClinetStorageFree int64 `x5tag:"clinet_storage_free"`
	ClientResVersion  string `x5tag:"client_res_version"`
	CarrierType int32 `x5tag:"carrier_type"`
	UnityLogin int32 `x5tag:"unity_login"`
	RegChannelID string `x5tag:"reg_channel_id"`
	WakeupSourcePrivilege int32 `x5tag:"wakeup_source_privilege"`
	Pfkey string  `x5tag:"pfkey"`
	PayToken string `x5tag:"pay_token"`
	CPUInfo string `x5tag:"cpu_info"`
}

// CEventVersionRequest ...
type CEventVersionRequest struct {
	DefaultServerID int32 `x5tag:"default_server_id"`
	ClientInfo MobileClientInfo `x5tag:"client_info"`
	CLSID int32 `x5tag:"CLSID"`
}

// Encode ...
func (p *CEventVersionRequest) Encode() []byte{
	allBuffs := bytes.NewBuffer([]byte{})
	binary.Write(allBuffs, binary.LittleEndian, p.DefaultServerID)
	
	return nil;
}
    