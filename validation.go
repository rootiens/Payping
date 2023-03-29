package main

import (
	"net/url"
	"os/exec"
)

func ReturnURLValidation(rurl string) bool {
	url, err := url.Parse(rurl)
	if err != nil || url.Scheme != "https" || url.Host == "" {
		return false
	}
	return true
}

func RefIDValidation(refid string) string {
	if len(refid) < 6 {
		uuid, _ := exec.Command("uuidgen").Output()
        refid = string(uuid)[:len(uuid)-1]
	}
	return refid
}
