package util

import (
	"net"

	"github.com/vishvananda/netlink"
)

func CreateIPSetGroup(groupName string) error {

	options := netlink.IpsetCreateOptions{}
	err := netlink.IpsetCreate(groupName, "hash:ip", options)
	if err != nil {
		return err
	}

	return nil
}

func DeleteIPSetGroup(groupName string) error {

	err := netlink.IpsetDestroy(groupName)
	if err != nil {
		return err
	}

	return nil
}

func AddIPToIPSet(groupName string, ip string) error {

	entry := &netlink.IPSetEntry{
		IP: net.ParseIP(ip),
	}

	err := netlink.IpsetAdd(groupName, entry)
	if err != nil {
		return err
	}

	return nil
}

func RemoveIPFromIPSet(groupName string, ip string) error {

	entry := &netlink.IPSetEntry{
		IP: net.ParseIP(ip),
	}

	err := netlink.IpsetDel(groupName, entry)
	if err != nil {
		return err
	}

	return nil
}
