package models

import "time"

// NetworkInfo network info such as subnet, local network, public ip,
// and isp, and lat and long
type NetworkInfo struct {
	LocalSubnet      string `json:"local_subnet"`
	PublicAddress    string `json:"public_address"`
	InternetProvider string `json:"internet_provider"`
	Lat              string `json:"lat"`
	Long             string `json:"long"`
}

type SpeedTestInfo struct {
	Latency time.Duration `json:"latency"`
	DLSpeed float64       `json:"dl_speed"`
	ULSpeed float64       `json:"ul_speed"`
	Server  string        `json:"server"`
	Host    string        `json:"host"`
}

type MtrTarget struct {
	Address string `json:"address"`
	Result  string `json:"result"`
}

type IcmpTarget struct {
	Address string `json:"address"`
	Result  struct {
		ElapsedMilliseconds int64 `json:"elapsed_ms"`
	} `json:"result"`
}
