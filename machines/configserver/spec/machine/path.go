package machine

// MachineConfigsvc paths
const (
	// swagger:route GET /machine-feeds/{machineid} machines ListMachineFeeds
	// ---
	// summary: List machine feeds
	// description: Retrieve machine feeds for a specific machine identified by its machine ID.
	// parameters:
	// - name: machineid
	//   in: path
	//   required: true
	//   type: string
	// responses:
	//   200:
	//     description: Successful operation
	ListPath = "/machine-feeds/:machineid"

	// swagger:route POST /machine-feeds machines createMachineFeeds
	// ---
	// summary: Create machine feeds
	// description: Create new machine feeds.
	// responses:
	//   201:
	//     description: Created
	CreatePath = "/machine-feeds"
)
