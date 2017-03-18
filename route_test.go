package route

import (
	"net"
	"testing"
)

func Test_Route(t *testing.T) {
	r := &Route{}
	r.LinkIndex = 1
	r.Dst = &net.IPNet{}
	r.Dst.IP = net.ParseIP("192.168.2.0")
	r.Dst.Mask = net.IPv4Mask(255, 255, 255, 0)
	r.Gw = net.ParseIP("192.168.0.222")
	println(r.String())
	RouteAdd(r)

	list, err := RouteList()
	if err != nil {
		println(err)

	} else {
		println("list")
		for index, v := range list {
			print(index)
			print(" ")
			println(v.String())

		}
	}
	println("get")

	list, err = RouteGet(net.ParseIP("192.168.2.0"))

	if err != nil {
		println(err)

	} else {
		println("list")
		for index, v := range list {
			print(index)
			print(" ")
			println(v.String())

		}
	}

	err = RouteDel(r)
	if err != nil {
		println(err)
	}

}
