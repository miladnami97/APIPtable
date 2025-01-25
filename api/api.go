package api

import (
	"log"
	"main/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FirewallRuleRequest represents a request for firewall rule operations.
type FirewallRuleRequest struct {
	IpsetName string `json:"ipset_name"`
	DstPort   string `json:"dst_port"`
	Action    string `json:"action"`
}

// IPSetGroupRequest represents a request for IPSet group operations.
type IPSetGroupRequest struct {
	GroupName string `json:"group_name"`
}

// IPRequest represents a request to manage IPs in IPSet groups.
type IPRequest struct {
	GroupName string `json:"group_name"`
	IP        string `json:"ip"`
}

// AddFirewallRule adds a firewall rule.
// @Summary Add a firewall rule
// @Description Adds a firewall rule based on the specified IP set, destination port, and action.
// @Tags Firewall
// @Accept json
// @Produce json
// // @Param request body FirewallRuleRequest true "Firewall rule details"
// @Success 200 {object} map[string]interface{} "Successful response"
// @Failure 400 {object} map[string]interface{} "Bad request response"
// @Failure 500 {object} map[string]interface{} "Internal server error response"
// @Router /firewall/rules [post]
func AddFirewallRule(c *gin.Context) {
	var req FirewallRuleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("ERROR - AddFirewallRule -", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Invalid request"})
		return
	}

	err := util.AddFirewallRule(req.IpsetName, req.DstPort, req.Action)
	if err != nil {
		log.Println("ERROR - AddFirewallRule -", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": "Internal Server Error"})
		return
	}
	log.Println("INFO - Iptable rule added for " + req.IpsetName + " port " + req.DstPort + " action " + req.Action)
	err = util.SaveIptablesRules()
	if err != nil {
		log.Println("ERROR - MakePersistent -", err)
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "rule added"})
}

// DeleteFirewallRules deletes a firewall rule.
// @Summary Delete a firewall rule
// @Description Deletes a firewall rule based on the specified IP set, destination port, and action.
// @Tags Firewall
// @Accept json
// @Produce json
// @Param request body FirewallRuleRequest true "Firewall rule details"
// @Success 200 {object} map[string]interface{} "Successful response"
// @Failure 400 {object} map[string]interface{} "Bad request response"
// @Failure 500 {object} map[string]interface{} "Internal server error response"
// @Router /firewall/rules [delete]
func DeleteFirewallRules(c *gin.Context) {

	var req FirewallRuleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("ERROR - DeleteFirewallRules -", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Invalid request"})
		return
	}

	err := util.DeleteFirewallRule(req.IpsetName, req.DstPort, req.Action)
	if err != nil {
		log.Println("ERROR - DeleteFirewallRules -", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": "Internal Server Error"})
		return
	}
	log.Println("INFO - Iptable rule deleted for ipset name: " + req.IpsetName + " port " + req.DstPort + " action " + req.Action)
	err = util.SaveIptablesRules()
	if err != nil {
		log.Println("ERROR - MakePersistent -", err)
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "rule deleted."})
}

// CreateIPSetGroup creates an IPSet group.
// @Summary Create an IPSet group
// @Description Creates an IPSet group with the specified group name.
// @Tags IPSet
// @Accept json
// @Produce json
// @Param request body IPSetGroupRequest true "IPSet group details"
// @Success 200 {object} map[string]interface{} "Successful response"
// @Failure 400 {object} map[string]interface{} "Bad request response"
// @Failure 500 {object} map[string]interface{} "Internal server error response"
// @Router /firewall/ipset [post]
func CreateIPSetGroup(c *gin.Context) {

	var req IPSetGroupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("ERROR - CreateIPSetGroup -", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Invalid request"})
		return
	}

	err := util.CreateIPSetGroup(req.GroupName)
	if err != nil {
		log.Println("ERROR - CreateIPSetGroup -", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": "Internal Server Error"})
		return
	}

	log.Println("INFO - IPSet group " + req.GroupName + " created.")
	err = util.SaveIptablesRules()
	if err != nil {
		log.Println("ERROR - MakePersistent -", err)
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "IPSet group created"})
}

// DeleteIPSetGroup deletes an IPSet group.
// @Summary Delete an IPSet group
// @Description Deletes an IPSet group with the specified group name.
// @Tags IPSet
// @Accept json
// @Produce json
// @Param request body IPSetGroupRequest true "IPSet group details"
// @Success 200 {object} map[string]interface{} "Successful response"
// @Failure 400 {object} map[string]interface{} "Bad request response"
// @Failure 500 {object} map[string]interface{} "Internal server error response"
// @Router /firewall/ipset [delete]
func DeleteIPSetGroup(c *gin.Context) {

	var req IPSetGroupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("ERROR - DeleteIPSetGroup -", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Invalid request"})
		return
	}

	err := util.DeleteIPSetGroup(req.GroupName)
	if err != nil {
		log.Println("ERROR - DeleteIPSetGroup -", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": "Internal Server Error"})
		return
	}

	log.Println("INFO - IPSet group " + req.GroupName + " Deleted")
	err = util.SaveIptablesRules()
	if err != nil {
		log.Println("ERROR - MakePersistent -", err)
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "IPSet group Deleted"})
}

// AddIPToIPSet adds an IP address to an IPSet group.
// @Summary Add IP to IPSet
// @Description Adds the specified IP address to the specified IPSet group.
// @Tags IPSet
// @Accept json
// @Produce json
// @Param request body IPRequest true "IPSet group and IP details"
// @Success 200 {object} map[string]interface{} "Successful response"
// @Failure 400 {object} map[string]interface{} "Bad request response"
// @Failure 500 {object} map[string]interface{} "Internal server error response"
// @Router /firewall/ipset/ip [post]
func AddIPToIPSet(c *gin.Context) {

	var req IPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("ERROR - AddIPToIPSet -", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Invalid request"})
		return
	}

	err := util.AddIPToIPSet(req.GroupName, req.IP)
	if err != nil {
		log.Println("ERROR - AddIPToIPSet -", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": "Internal Server Error"})
		return
	}

	log.Println("INFO - IP added " + req.IP + " to IPSet group " + req.GroupName)
	err = util.SaveIptablesRules()
	if err != nil {
		log.Println("ERROR - MakePersistent -", err)
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "IPSet group edited"})
}

// RemoveIPFromIPSet removes an IP address from an IPSet group.
// @Summary Remove IP from IPSet
// @Description Removes the specified IP address from the specified IPSet group.
// @Tags IPSet
// @Accept json
// @Produce json
// @Param request body IPRequest true "IPSet group and IP details"
// @Success 200 {object} map[string]interface{} "Successful response"
// @Failure 400 {object} map[string]interface{} "Bad request response"
// @Failure 500 {object} map[string]interface{} "Internal server error response"
// @Router /firewall/ipset/ip [delete]
func RemoveIPFromIPSet(c *gin.Context) {

	var req IPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("ERROR - RemoveIPFromIPSet -", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Invalid request"})
		return
	}

	err := util.RemoveIPFromIPSet(req.GroupName, req.IP)
	if err != nil {
		log.Println("ERROR - RemoveIPFromIPSet", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": "Internal Server Error"})
		return
	}

	log.Println("INFO - IP " + req.IP + " removed from IPSet group " + req.GroupName)
	err = util.SaveIptablesRules()
	if err != nil {
		log.Println("ERROR - MakePersistent -", err)
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "IPSet group edited"})
}
