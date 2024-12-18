package handlers

import (
	"net/http"

	"machines/configserver/bl"
	"machines/configserver/spec/machine"
	"machines/configserver/svcparams"

	"github.com/gin-gonic/gin"
)

// This can be initialized and handled in some better way
var mcfBL = bl.NewMachineFeedBL()

func ListMachineFeeds(c *gin.Context) {
	machineid := c.Param(svcparams.MachinIDStr)

	result, err := mcfBL.ListMachineFeed(c.GetUint64(machineid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	contentType := c.GetHeader("Content-Type")
	if contentType == "" || contentType == "application/json" {
		c.JSON(http.StatusOK, result)
	} else if contentType == "application/xml" {
		c.XML(http.StatusOK, result)
	} else {
		c.String(http.StatusOK, "received %s", result)
	}
}

func CreateMachineFeed(c *gin.Context) {
	var machineFeed machine.Machine

	err := c.ShouldBind(&machineFeed)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	err = mcfBL.CreateMachineFeed(&machineFeed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
