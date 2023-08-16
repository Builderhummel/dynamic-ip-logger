package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func getPublicIpv4FromDns() string {
	ip, err := exec.Command("dig", "+short", "-4", os.Getenv("DIL_DYNDNS"), "@resolver1.opendns.com").Output()
	_ = err
	return string(ip)
}

// func getPublicIpv6FromDns() string {
// 	ip, err := exec.Command("dig", "+short", "AAAA", "-6", os.Getenv("DIL_DYNDNS"), "@resolver1.opendns.com").Output()
// 	_ = err
// 	return string(ip)
// }

//function should save the ip, the current date and time in Unix timestamp in a csv file
func saveIpInCsv(ip string) {
	filename := os.Getenv("DIL_FILE")
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		//file does not exist
		//create file and add header
		file, err := os.Create(filename)
		_ = err
		defer file.Close()

		_, _ = file.WriteString("ip,unixtime\n")
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	_ = err
	defer file.Close()

	//remove newline from ip
	ip = ip[:len(ip)-1]

	//now ip and unix timestamp in csv format
	now := time.Now().Unix()
	csv := fmt.Sprintf("%s,%d\n", ip, now)
	_, _ = file.WriteString(csv)
}

func main() {
	saveIpInCsv(getPublicIpv4FromDns())
	//Run successfully message with timestamp
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "Run successfully")
}
