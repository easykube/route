package route

import (
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/easyKube/util"
)

//添加路由
func RouteAdd(route *Route) error {
	bmask := route.Dst.Mask
	mask := fmt.Sprintf("%d.%d.%d.%d", bmask[0], bmask[1], bmask[2], bmask[3])
	args := "add " + route.Dst.IP.String() + " mask " + mask + " " + route.Gw.String()
	out, err := util.ExecCmd("route", strings.Split(args, " ")...)
	if err != nil {
		return err
	}

	if !strings.Contains(out, "操作完成") {
		if strings.Contains(out, "对象已存在") {
			return nil
		}
		return errors.New("添加路由失败：" + args + "  " + out)
	}
	return err
}

//删除路由
func RouteDel(route *Route) error {
	bmask := route.Dst.Mask
	mask := fmt.Sprintf("%d.%d.%d.%d", bmask[0], bmask[1], bmask[2], bmask[3])
	args := "delete " + route.Dst.IP.String() + " mask " + mask + " " + route.Gw.String()
	out, err := util.ExecCmd("route", strings.Split(args, " ")...)
	if err != nil {
		return err
	}

	if !strings.Contains(out, "操作完成") {
		if strings.Contains(out, "路由删除失败: 找不到元素") {
			return nil
		}
		return errors.New("删除路由失败：" + args + "  " + out)
	}
	return err
}

func parseRoute(out string) []Route {

	lines := strings.Split(out, "\n")
	iBegin := -1
	routes := []Route{}
	for index, line := range lines {
		if strings.Contains(line, "活动路由") || strings.Contains(line, "永久路由") {
			iBegin = index
		} else {
			if iBegin > 0 {
				if index > iBegin+1 {
					words := SplitSpace(line)
					if len(words) >= 3 && net.ParseIP(words[0]) != nil && net.ParseIP(words[1]) != nil && net.ParseIP(words[2]) != nil {
						route := Route{}
						route.LinkIndex = index
						route.Dst = &net.IPNet{}
						route.Dst.IP = net.ParseIP(words[0])
						route.Dst.Mask = net.IPv4Mask(255, 255, 255, 0)

						route.Gw = net.ParseIP(words[2])
						routes = append(routes, route)
					}

				}

			}

		}
	}

	return routes

}

//获取路由
func RouteGet(destination net.IP) ([]Route, error) {
	args := "print " + destination.String()
	out, err := util.ExecCmd("route", strings.Split(args, " ")...)
	if err != nil {
		return nil, err
	}
	return parseRoute(out), nil
}

//按空格拆分并压缩
func SplitSpace(s string) []string {
	words := strings.Split(s, " ")
	ret := []string{}
	for _, w := range words {
		w = strings.TrimSpace(w)
		if w != "" {
			ret = append(ret, w)
		}

	}
	return ret

}

//显示路由列表
func RouteList() ([]Route, error) {
	out, err := util.ExecCmd("route", "print")
	if err != nil {
		return nil, err
	}

	return parseRoute(out), nil
}

//初始化路由器设置
func InitRouter(ip string) {

}
