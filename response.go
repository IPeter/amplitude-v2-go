package amplitude

type Response struct {
	// success code
	Code int `json:"code"`
	// Error description. Possible values are Invalid request path, Missing request body, Invalid JSON request body, Request missing required field, Invalid event JSON, Invalid API key, Invalid field values on some events
	Error *string `json:"error"`
	// Indicates which request-level required field is missing.
	MissingField *string `json:"missing_field"`
}
