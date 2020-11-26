package amplitude

type Options struct {
	// Minimum permitted length for user_id & device_id fields
	MinIDLength int `json:"min_id_length"`
}
