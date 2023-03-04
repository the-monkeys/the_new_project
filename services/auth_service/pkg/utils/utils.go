package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/google/uuid"
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

func GetUUID() string {
	uuid := uuid.New()
	id := uuid.ID()

	return strconv.Itoa(int(id))
}
