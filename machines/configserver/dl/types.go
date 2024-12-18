package dl

import (
	"machines/configserver/spec/machine"
	"sync"
)

// MachineDL represent the main interface
// which can be implemented by various variation of DL layers
type MachineDL interface {
	Create(machineFeed *machine.Machine) (err error)
	List(machineID uint64) (machineFeed []*machine.Machine, err error)
	// Update(machineID uint64, machineFeed *machine.Machine)
	// Delete(machineID uint64)
}

// MachineFeedBL represent the struct which implements the MachineDL interface
// This is intentionally done to achieve loose coupling between handler and bl layers
// We can easily substitute new BL in place of this if required
type MachineFeedDL struct {
	// mcMap is machine feeds map used for this assignement,
	// we are keeping db in memory in form of map
	// map is not concurrency (/thread) safe
	// so lock will be required to access/update this critical resource
	mcMap map[uint64][]*machine.Machine

	// here we have used lock
	// but in actual production environment
	// concurrent read, writes will be taken care by
	// external DB app like
	// AWS timesteam or InfluxDB which are time series DB
	// Another option we have to use dynamoDB but it is not optimized for
	// time series use cases.
	mu *sync.RWMutex
}
