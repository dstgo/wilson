package whois

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
)

func WhoIs(ip string) string {
	ip = net.ParseIP(ip).String()
	if len(ip) == 0 {
		return ""
	}

	c := &http.Client{Timeout: 2 * time.Second}
	resp, err := c.Get(fmt.Sprintf("https://whois.pconline.com.cn/ip.jsp?ip=%slevel=2", ip))
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return ""
	}

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	bytes, err := simplifiedchinese.GBK.NewDecoder().Bytes(bs)
	if err != nil {
		return ""
	}
	return string(bytes)
}
