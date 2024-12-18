package dl

import (
	"errors"
	"machines/configserver/spec/machine"
	"machines/configserver/svcerror"
	"machines/configserver/svcparams"
	"sync"
	"time"
)

// This pacakge represents the data layer
// it is a separate layer from business logic layer which take cares of actual DB CRUD operations
func NewMachineFeedDL() MachineDL {
	return &MachineFeedDL{
		mcMap: make(map[uint64][]*machine.Machine),
		mu:    &sync.RWMutex{},
	}
}

// Create dl api will store the new machine feed entry in our in-memory db
func (mc *MachineFeedDL) Create(machineFeed *machine.Machine) (err error) {
	mc.mu.Lock()
	defer mc.mu.Unlock()

	// // mocking the data for repairs
	// machineFeed = &machine.Machine{
	// 	ID:        10001,
	// 	CreatedAt: time.Now().Unix(),
	// 	Type:      svcparams.RepairType,
	// 	Data: map[string]interface{}{
	// 		"repair_type":    "engine",
	// 		"cost":           500.0,
	// 		"parts_replaced": []string{"engine", "piston"},
	// 		"additional_info": map[string]interface{}{
	// 			"mechanic": "John Doe",
	// 			"status":   "Completed",
	// 		},
	// 	},
	// }

	// // mocking the data for sessions
	// machineFeed = &machine.Machine{
	// 	ID:        10002,
	// 	CreatedAt: time.Now().Unix(),
	// 	Type:      svcparams.SessionType,
	// 	Data: map[string]interface{}{
	// 		"session_type":      "pump",
	// 		"temp":              "50C",
	// 		"vibration_per_min": 200,
	// 	},
	// }

	mc.mcMap[uint64(machineFeed.ID)] = append(mc.mcMap[uint64(machineFeed.ID)], machineFeed)
	// for simplicity we are returning from here
	return nil
}

// List dl api will return all machine feeds entry in our in-memory db
// need to accept time range between which feeds needs to be return
// We need to return certain items using cookie or token mechanism
// This will help in implementing pagination mechanism
func (mc *MachineFeedDL) List(machineID uint64) (machineFeed []*machine.Machine, err error) {
	mc.mu.Lock()
	defer mc.mu.Unlock()

	// mocking the data for repairs
	// this data will be coming of other micro service
	// i.e. Repair micro service
	machineRepairFeed := &machine.Machine{
		ID:        10001,
		CreatedAt: time.Now().Unix(),
		Type:      svcparams.RepairType,
		Data: map[string]interface{}{
			"repair_type":    "engine",
			"cost":           500.0,
			"parts_replaced": []string{"engine", "piston"},
			"additional_info": map[string]interface{}{
				"mechanic": "John Doe",
				"status":   "Completed",
			},
		},
	}

	// mocking the data for sessions
	// this data will be coming of other micro service
	// i.e. Session micro service
	machineSessionFeed := &machine.Machine{
		ID:        10002,
		CreatedAt: time.Now().Unix(),
		Type:      svcparams.SessionType,
		Data: map[string]interface{}{
			"session_type":      "pump",
			"temp":              "50C",
			"vibration_per_min": 200,
		},
	}

	// we are aggregaing the data as part of response
	machineFeed = append(machineFeed, machineRepairFeed, machineSessionFeed)
	mc.mcMap[uint64(machineID)] = machineFeed
	machineFeed, exist := mc.mcMap[uint64(machineID)]
	if !exist {
		return nil, errors.New(svcerror.ErrNoFeedsExist)
	}

	return machineFeed, nil
}
