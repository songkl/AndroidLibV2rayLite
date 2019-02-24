package CoreI

import (
	"strconv"

	"v2ray.com/core"
)

type Status struct {
	IsRunning       bool
	VpnSupportnodup bool
	PackageName     string
	DomainName      string

	Vpoint core.Server
}

func CheckVersion() int {
	return 16
}

func (v *Status) GetDataDir() string {
	return v.PackageName
}

func (v *Status) GetApp(name string) string {
	return v.PackageName + name
}

func (v *Status) GetTun2socksArgs(fd int) []string {
	return []string{"--netif-ipaddr",
		"26.26.26.2",
		"--netif-netmask",
		"255.255.255.252",
		"--socks-server-addr",
		"127.0.0.1:10808",
		"--tunfd",
		strconv.Itoa(fd),
		"--tunmtu",
		"1500",
		"--sock-path",
		"/dev/null",
		"--loglevel",
		"1",
		"--netif-ip6addr",
		"fe80:2626::2",
		"--enable-udprelay"}
}

func (v *Status) GetDomainNameList() []string {
	var dynaArr []string
	if v.DomainName != "" {
		dynaArr = append(dynaArr, v.DomainName)
	}
	return dynaArr
}

func (v *Status) GetVPNSetupArg() string {
	return "m,1500 a,26.26.26.1,30 a,fe80:2626::1,126 r,0.0.0.0,0 r,::,0 d,26.26.26.2"
}
