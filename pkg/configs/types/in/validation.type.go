package types

// validation error
type ValidationError struct {
	Error       bool        `json:"error"`
	FailedField string      `json:"failedField"`
	Tag         string      `json:"tag"`
	Value       interface{} `json:"value"`
	ErrorMsg    string      `json:"errorMsg"`
}
