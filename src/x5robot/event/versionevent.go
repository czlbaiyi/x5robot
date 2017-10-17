package event

import (
	"x5robot/x5base"
)
/////////////////////////////CEventVersionRequest///////////////////////////////
// MobileClientInfo ...
type MobileClientInfo struct {
	ClientVersion         string `x5tag:"client_version"`
	ClientPlatform        int32  `x5tag:"client_platform"`
	LoginSource           int32  `x5tag:"login_source"`
	DeviceInfo            string `x5tag:"device_info"`
	OpenID                string `x5tag:"open_id"`
	OpenKey               string `x5tag:"open_key"`
	Pf                    string `x5tag:"pf"`
	NetInfo               string `x5tag:"net_info"`
	DeviceSystem          string `x5tag:"device_system"`
	Location              string `x5tag:"location"`
	ClientMemorySize      int32  `x5tag:"client_memory_size"`
	ClientStorageTotal    int64  `x5tag:"client_storage_total"`
	ClinetStorageFree     int64  `x5tag:"clinet_storage_free"`
	ClientResVersion      string `x5tag:"client_res_version"`
	CarrierType           int32  `x5tag:"carrier_type"`
	UnityLogin            int32  `x5tag:"unity_login"` // 0,非unity登陆  //1. unity登陆
	RegChannelID          string `x5tag:"reg_channel_id"`
	WakeupSourcePrivilege int32  `x5tag:"wakeup_source_privilege"`
	Pfkey                 string `x5tag:"pfkey"`
	PayToken              string `x5tag:"pay_token"`
	CPUInfo               string `x5tag:"cpu_info"`
}

// CEventVersionRequest ...
type CEventVersionRequest struct {
	DefaultServerID            int32             `x5tag:"default_server_id"`
	ClientInfo                 MobileClientInfo  `x5tag:"client_info"`
	DefaultServerIDList        []int32           `x5tag:"3"`
	x5base.CEventMobileRequest `x5tag:"inherit"`
}

/////////////////////////////CEventVersionRequestResult///////////////////////////////
type ResourceVersionInfo struct {
	Version string `x5tag:version`
	Vin_incr_ver string `x5tag:min_incr_ver`
	Incr_url string `x5tag:incr_url`
	Incr_size int32 `x5tag:incr_size`
	Incr_md5 string `x5tag:incr_md5`
	All_url string `x5tag:all_url`
	All_size int32 `x5tag:all_size`
	All_md5 string `x5tag:all_md5`
	All_version string `x5tag:all_version`
}

type AppVersionInfo struct {
	Version string `x5tag:version`
	Min_ignore_version string `x5tag:min_ignore_version`
	Is_force_update bool `x5tag:is_force_update`
	All_url string `x5tag:all_url`
	All_size int32 `x5tag:all_size`
	All_md5 string `x5tag:all_md5`
}

type HostInfo struct {
	Ip string `x5tag:ip`
	Port int32 `x5tag:port`
	Type int32 `x5tag:type`
}

type RegionStatusInfo struct {
	M_id int32 `x5tag:m_id`
	M_name string `x5tag:m_name`
	Server_status int32 `x5tag:server_status`
	Online_players int32 `x5tag:online_players`
	M_zone_id int32 `x5tag:m_zone_id`
	M_is_open bool `x5tag:m_is_open`
	Target_gateway_list []HostInfo `x5tag:target_gateway_list`
}

type UpdateNotice struct {
	Notice string `x5tag:1`
	Title string `x5tag:2`
	Pic_urls []string `x5tag:3`
}

type MobileNetToken struct {
	Type int16 `x5tag:1`
	Serial int16 `x5tag:2`
	Token_id int64 `x5tag:3`
}

type RoleSimpleShowInfo struct {
	M_server_id int32 `x5tag:1`
	M_level int32 `x5tag:2`
	M_sex int32 `x5tag:3`
	M_month_login_count int32 `x5tag:4`
	M_loading_pic string `x5tag:5`
	M_zone_id int32 `x5tag:6`
}

type CEventVersionRequestResult struct {
	NRet int32 `x5tag:nRet`
	Res_infos []ResourceVersionInfo `x5tag:res_infos`
	App_infos []AppVersionInfo `x5tag:app_infos`
	M_special_status RegionStatusInfo `x5tag:4`
	M_update_notice UpdateNotice `x5tag:5`
	M_notice string `x5tag:6`
	Token MobileNetToken `x5tag:7`
	M_carrier_type int32 `x5tag:8`
	Zone_name_map map[int32]string `x5tag:9`
	M_select_role_info RoleSimpleShowInfo `x5tag:10`

	x5base.CEventMobileRequest `x5tag:"inherit"`
}
