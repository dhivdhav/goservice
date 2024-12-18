package bl

import (
	"errors"
	"log"
	"machines/configserver/dl"
	"machines/configserver/spec/machine"
)

// This package represents the business logic layer, separate from the data layer which handles actual DB CRUD operations
func NewMachineFeedBL() MachineBL {
	return &MachineFeedBL{
		inMemDL: dl.NewMachineFeedDL(),
	}
}

// ListMachineFeed retrieves a list of machine feeds for a given machine ID.
// It first attempts to fetch the machine feeds from the in-memory data layer.
// If an error occurs, it performs a retry based on the error and applies a fallback mechanism.
func (mc *MachineFeedBL) ListMachineFeed(machineID uint64) (machineFeed []*machine.Machine, err error) {
	machineFeed, err = mc.inMemDL.List(machineID)
	if err != nil {
		// do some retry based on error
		// fall back mechanism
		return nil, err
	}
	return machineFeed, nil
}

// CreateMachineFeed creates a new machine feed in the in-memory data layer.
// Returns: err: An error object if an error occurs during the creation process or if the input is invalid.
func (mc *MachineFeedBL) CreateMachineFeed(machineFeed *machine.Machine) (err error) {
	if machineFeed == nil {
		return errors.New("machineFeed cannot be nil")
	}
	if mc.inMemDL == nil {
		return errors.New("inMemDL is not initialized")
	}

	err = mc.inMemDL.Create(machineFeed)
	if err != nil {
		// log the error
		log.Printf("Error creating machine feed: %v", err)
		return err
	}
	return nil
}
