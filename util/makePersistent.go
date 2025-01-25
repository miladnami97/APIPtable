package util

import (
	"log"
	"os/exec"
)

func SaveIptablesRules() error {
	cmd := exec.Command("sudo", "bash", "-c", "iptables-save > /etc/iptables/rules.v4")
	err := cmd.Run()
	if err != nil {
		log.Println("ERROR - SaveIptablesRules:", err)
		return err
	}
	log.Println("INFO - iptables rules saved successfully")
	return nil
}
