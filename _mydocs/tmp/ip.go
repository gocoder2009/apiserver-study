package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 获取X-Real-IP头部字段值
		xRealIP := r.Header.Get("X-Real-IP")

		// 获取X-Forwarded-For头部字段值
		xForwardedFor := r.Header.Get("X-Forwarded-For")

		var clientIP string
		if xRealIP != "" {
			clientIP = xRealIP
		} else if xForwardedFor != "" {
			// 如果X-Forwarded-For存在，可能会有多个IP地址，取第一个即为真实IP
			clientIP = strings.Split(xForwardedFor, ", ")[0]
		} else {
			// 如果X-Real-IP和X-Forwarded-For都不存在，则直接使用RemoteAddr作为客户端IP
			clientIP = r.RemoteAddr
		}

		fmt.Fprintf(w, "真实客户端IP地址：%s\n", clientIP)
	})

	http.ListenAndServe(":8080", nil)
}
