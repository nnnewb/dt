package tm

import (
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

// GenGID 生成 gid ，格式 tm-<IP>-<时间戳哈希>
func GenGID() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		// 取第一个非本地环回(127.0.0.1) 同时有 IPv4 地址的网卡接口
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ip := ipnet.IP.To4().String()
			ns := strings.Split(ip, ".")
			r := []byte{}
			for _, n := range ns {
				result, err := strconv.Atoi(n)
				if err != nil {
					return "", err
				}
				r = append(r, byte(result))
			}

			return fmt.Sprintf("tm-gi-%s-%d", hex.EncodeToString(r), time.Now().UnixMilli()), nil
		}
	}

	return "", errors.New("no valid net interface found for generate GID")
}

// MustGenGID 生成 gid ，格式参见 GenGID ，失败则 panic
func MustGenGID() string {
	gid, err := GenGID()
	if err != nil {
		panic(err)
	}
	return gid
}

func GenBranchID(branch string) (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		// 取第一个非本地环回(127.0.0.1) 同时有 IPv4 地址的网卡接口
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ip := ipnet.IP.To4().String()
			ns := strings.Split(ip, ".")
			r := []byte{}
			for _, n := range ns {
				result, err := strconv.Atoi(n)
				if err != nil {
					return "", err
				}
				r = append(r, byte(result))
			}

			return fmt.Sprintf("tm-bi-%s-%d", hex.EncodeToString(r), time.Now().UnixMilli()), nil
		}
	}

	return "", errors.New("no valid net interface found for generate BranchID")
}

func MustGenBranchID(branch string) string {
	bi, err := GenBranchID(branch)
	if err != nil {
		panic(err)
	}
	return bi
}
