package machine

// Machine model contains information about various configurations
type Machine struct {
	ID        int32                  `json:"id"`
	Type      int                    `json:"feedType"`
	CreatedAt int64                  `json:"created_at"`
	Data      map[string]interface{} `json:"data"`
}
