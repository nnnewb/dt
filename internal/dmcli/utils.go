package dmcli

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/metadata"
)

// GenGID 生成 gid ，格式 dm-<IP>-<时间戳哈希>
func genGID() (string, error) {
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

			return fmt.Sprintf("dm-%s-%d", hex.EncodeToString(r), time.Now().UnixMilli()), nil
		}
	}

	return "", errors.New("no valid net interface found for generate GID")
}

func getOrCreateGID(ctx context.Context) (string, error) {
	var gid string
	if md, ok := metadata.FromIncomingContext(ctx); ok && len(md.Get("dm-gid")) != 0 {
		gid = md.Get("dm-gid")[0]
	} else {
		// TODO 调用 dm 的 create-global-transaction
	}

	return gid, nil
}