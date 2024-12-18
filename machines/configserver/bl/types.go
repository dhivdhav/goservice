package bl

import (
	"machines/configserver/dl"
	"machines/configserver/spec/machine"
)

// MachineBL represent the main interface
// which can be implemented by various business logic layers
type MachineBL interface {
	CreateMachineFeed(machineFeed *machine.Machine) (err error)
	ListMachineFeed(machineID uint64) (machineFeed []*machine.Machine, err error)
	// Delete(machineID uint64)
}

// MachineFeedBL represent the struct which implements the MachineBL interface
// This is intentionally done to achieve loose coupling between handler and bl layers
// We can easily substitute new BL in place of this if required
type MachineFeedBL struct {
	inMemDL dl.MachineDL
}
