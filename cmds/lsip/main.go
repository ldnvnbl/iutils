package main

import (
	"fmt"
	"os"

	"github.com/ldnvnbl/iutils/lsip"
)

func main() {
	ips, err := lsip.ListIPInfo()
	if err != nil {
		fmt.Fprintf(os.Stderr, "list ip info failed: %v", err)
		return
	}

	for _, i := range ips {
		i.Output(os.Stdout)
	}
}
