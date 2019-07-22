package util

import (
	"io/ioutil"
	"log"
	"net/http"
)

// HTTPGet Http Get请求
func HTTPGet(url string) (body []byte) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	defer resp.Body.Close()

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return body

}

// GetReqIP 从*http.Request 获取请求来源IP
// 先从 header 的 Remote_addr 获取 如果为空
// 从 *http.Request.RemoteAddr 获取
func GetReqIP(r *http.Request) string {
	ip := r.Header.Get("Remote_addr")
	if ip == "" {
		ip = r.RemoteAddr
	}
	return ip
}
