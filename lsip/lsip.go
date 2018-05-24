package lsip

import (
	"fmt"
	"io"
	"net"
	"os"
)

type IPInfo struct {
	IPs           []string
	InterfaceName string
}

func (i IPInfo) Output(w io.Writer) {

	if w == nil {
		w = os.Stdout
	}

	for _, ip := range i.IPs {
		fmt.Fprintf(w, "%s %s\n", i.InterfaceName, ip)
	}
}

func ListIPInfo() (ret []IPInfo, err error) {
	ret = []IPInfo{}

	ifs, err := net.Interfaces()
	if err != nil {
		return
	}

	for _, f := range ifs {
		ipi := IPInfo{
			InterfaceName: f.Name,
		}
		addrs, err := f.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			ip, _, _ := net.ParseCIDR(addr.String())

			if ip != nil {
				ip = ip.To4()
			}

			if ip == nil {
				continue
			}
			ipi.IPs = append(ipi.IPs, ip.String())
		}
		ret = append(ret, ipi)
	}
	return
}
