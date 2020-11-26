package amplitude

type request struct {
	// Amplitude project API key
	ApiKey string `json:"api_key"`
	// Array of Events to upload
	Events []Event `json:"events"`
	// Object
	Options *Options `json:"options"`
}
