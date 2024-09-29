package dns

import (
	"fmt"
	"net"
)

var unresolvedIP = "unresolvedIP"

func Resolve(host string, port string) (string, error) {
	ip, err := net.LookupHost(host)
	if err != nil {
		return unresolvedIP, err
	}
	addr := fmt.Sprintf("%v:%v", ip, port)
	return addr, nil
}
