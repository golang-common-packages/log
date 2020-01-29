package log

// Fluent model provide for Fluent config
type Fluent struct {
	Tag    string `json:tag`
	Host   string `json:"host"`
	Port   int    `json:"port"`
	Prefix string `json:"prefix"`
}
