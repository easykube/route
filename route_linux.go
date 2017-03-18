package route

import (
	"net"

	"github.com/EasyKube/EasyRouter/util"
	"github.com/vishvananda/netlink"
)

//添加路由
func RouteAdd(route *Route) error {
	err := netlink.RouteAdd(route)
	if err != nil && Strings.Container(err.String(), "file exists") {
		return nil
	}
	return err
}

//删除路由
func RouteDel(route *Route) error {
	return netlink.RouteDel(route)
}

//获取路由
func RouteGet(destination net.IP) ([]Route, error) {
	return netlink.RouteGet(destination)
}

//显示所有路由
func RouteList() ([]Route, error) {
	return netlink.RouteList()
}

//初始化路由器设置
func InitRouter(ip string) {
	util.ExecCmdLine("echo 1>/proc/sys/net/ipv4/ip_forward ")
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
	util.ExecCmdLine("iptables -t nat -A POSTROUTING -o " + eth + " -j MASQUERADE")
	util.ExecCmdLine("iptables -A FORWARD -i " + eth + " -j ACCEPT")

}
