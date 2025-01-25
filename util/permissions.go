package util

import (
	"github.com/coreos/go-iptables/iptables"
)

func CheckIptablesPermission() (bool, error) {
	ip, err := iptables.New()
	if err != nil {
		return false, err
	}

	_, err = ip.ListChains("filter")
	if err != nil {
		return false, err

	}

	return true, nil
}
