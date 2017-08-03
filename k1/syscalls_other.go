//
//   date  : 2017-07-14
//   author: xjdrew
//

// +build !linux,!darwin,!windows

package k1

import (
	"errors"
	"net"
)

var errOS = errors.New("unsupported os")

func initTun(tun string, ipNet *net.IPNet, mtu int) error {
	return errOS
}

func addRoute(tun string, subnet *net.IPNet) error {
	return errOS
}

func fixDnsPort(ip net.IP) net.IP {
	return ip
}
