package route

import (
	"fmt"
	"net"
)

// Scope is an enum representing a route scope.
type Scope uint8

// Route represents a netlink route. A route is associated with a link,
// has a destination network, an optional source ip, and optional
// gateway. Advanced route parameters and non-main routing tables are
// currently not supported.
type Route struct {
	LinkIndex int
	Scope     Scope
	Dst       *net.IPNet
	Src       net.IP
	Gw        net.IP
}

func (r Route) String() string {
	return fmt.Sprintf("{Interface index: %d Dst: %s Src: %s Gw: %s}", r.LinkIndex, r.Dst,
		r.Src, r.Gw)
}

func RouteEquals(a, b *Route) bool {
	if a == nil || b == nil {
		return false
	}
	if a.Gw.String() == b.Gw.String() {
		if a.Dst == b.Dst {
			return true
		} else if a.Dst != nil && b.Dst != nil && a.Dst.String() == b.Dst.String() {
			return true
		}
	}
	return false
}
