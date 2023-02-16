package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func PublicIP() string {
	resp, err := http.Get("https://ifconfig.co/ip")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Public IP address:", string(ip))
	return string(ip)
}
