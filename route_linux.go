package route

import (
	"net"
	"strings"

	"github.com/easykube/util"
	"github.com/vishvananda/netlink"
)

//添加路由
func RouteAdd(route *Route) error {
	r := &netlink.Route{}
	r.Dst = route.Dst
	r.Gw = route.Gw
	r.Src = route.Src
	r.LinkIndex = route.LinkIndex
	r.Scope = netlink.Scope(route.Scope)

	err := netlink.RouteAdd(r)
	if err != nil && strings.Contains(err.Error(), "file exists") {
		return nil
	}
	return err
}

//删除路由
func RouteDel(route *Route) error {
	r := &netlink.Route{}
	r.Dst = route.Dst
	r.Gw = route.Gw
	r.Src = route.Src
	r.LinkIndex = route.LinkIndex
	r.Scope = netlink.Scope(route.Scope)

	return netlink.RouteDel(r)
}

//获取路由
func RouteGet(destination net.IP) ([]Route, error) {
	rs, err := netlink.RouteGet(destination)
	if err != nil {
		return nil, err
	}
	list := make([]Route, len(rs))
	for i, r := range rs {
		r0 := Route{}
		r0.Dst = r.Dst
		r0.Gw = r.Gw
		r0.Src = r.Src
		r0.LinkIndex = r.LinkIndex
		r0.Scope = Scope(r.Scope)
		list[i] = r0

	}
	return list, nil
}

//显示所有路由
func RouteList() ([]Route, error) {
	rs, err := netlink.RouteList(nil, 0)
	if err != nil {
		return nil, err
	}
	list := make([]Route, len(rs))
	for i, r := range rs {
		r0 := Route{}
		r0.Dst = r.Dst
		r0.Gw = r.Gw
		r0.Src = r.Src
		r0.LinkIndex = r.LinkIndex
		r0.Scope = Scope(r.Scope)
		list[i] = r0

	}
	return list, nil

}

//初始化路由器设置
//eth网卡名称
func InitRouter(eth string) {
	//out, err := util.ExecCmdLine("echo 1>/proc/sys/net/ipv4/ip_forward ")
	//if err != nil {
	//	util.Log(err)
	//} else {
	//	println(out)
	//}
	/**
	list, err := netlink.LinkList
	if err != nil {
		println(err)
		return
	}
	eth := ""
	for _,link range list{
		addrs, err2 := netlink.AddrList(link, FAMILY_ALL)
	    if err2 != nil {
		   continue
	    }
		for _,addr range addrs{
			if addr.IP.String()=ip{
				eth= link.name
			}
		}
	}
	if eth==""{
		return
	}
	**/
	util.ExecCmdLine("iptables -t nat -A POSTROUTING -o " + eth + " -j MASQUERADE")
	util.ExecCmdLine("iptables -A FORWARD -i " + eth + " -j ACCEPT")

}
