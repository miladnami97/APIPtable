package util

import (
	"github.com/coreos/go-iptables/iptables"
)

func AddFirewallRule(ipsetName string, dstPort string, action string) error {
	ipt, err := iptables.New()

	if err != nil {
		return err
	}
	err = ipt.AppendUnique("filter", "INPUT", "-p", "tcp", "--dport", dstPort, "-m", "set", "--match-set", ipsetName, "src", "-j", action)

	if err != nil {
		return err
	}

	return nil
}

func DeleteFirewallRule(ipsetName string, dstPort string, action string) error {
	ipt, err := iptables.New()
	if err != nil {
		return err
	}

	err = ipt.Delete("filter", "INPUT", "-p", "tcp", "--dport", dstPort, "-m", "set", "--match-set", ipsetName, "src", "-j", action)
	if err != nil {
		return err
	}
	return nil
}

func AddIPToSet(setName, ip string) error {
	ipt, err := iptables.New()
	if err != nil {
		return err
	}

	err = ipt.AppendUnique("filter", "INPUT", "-m", "set", "--match-set", setName, "src", "-j", "DROP")
	if err != nil {
		return err
	}

	return nil
}

func DeleteIPFromSet(setName, ip string) error {
	ipt, err := iptables.New()
	if err != nil {
		return err
	}

	err = ipt.Delete("filter", "INPUT", "-m", "set", "--match-set", setName, "src", "-j", "DROP")
	if err != nil {
		return err
	}

	return nil
}
