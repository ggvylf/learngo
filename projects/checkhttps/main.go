package main

import (
	"crypto/tls"
	"fmt"
	"strings"
	"time"

	"gitee.com/pdudo/SampleHttp"
)

func sslcheck(domainName string) (before time.Time, after time.Time, err error) {
	if !strings.HasSuffix(domainName, ":443") {
		domainName = domainName + ":443"
	}
	if strings.HasPrefix(domainName, "https://") {
		domainName = strings.SplitN(domainName, "https://", 2)[1]
	}

	conn, err := tls.Dial("tcp", domainName, nil)

	if err != nil {
		return time.Now(), time.Now(), err
	}

	before = conn.ConnectionState().PeerCertificates[0].NotBefore
	after = conn.ConnectionState().PeerCertificates[0].NotAfter

	return
}

func main() {
	SampleHttp.Route("get", "/sslcheck", func(info *SampleHttp.HttpInfo) {
		domainName := info.RequestData["domain_name"]

		var toClient string
		createDate, expireDate, err := sslcheck(domainName)
		if err != nil {
			toClient = fmt.Sprintf("{\"code\": \"1\",\"data\": {\"msg\": \"%s\"}}", err)
		} else {
			fmt.Println(createDate, expireDate, err)
			toClient = fmt.Sprintf("{\"code\": \"0\",\"data\": {\"domainName\": \"%s\","+
				"\"createDate\": \"%s\","+
				"\"expireDate\":\"%s\"}}"+
				"", domainName,
				createDate.Format("2006-01-02 15:04:05"),
				expireDate.Format("2006-01-02 15:04:05"))
		}

		info.SetResponseHeader("Content-Type", "application/json")

		info.Write([]byte(toClient))
	})
	SampleHttp.StartServer("0.0.0.0:8082")
}
