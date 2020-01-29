package log

// ILog store function in logging package
type ILog interface {
	Write(p []byte) (n int, err error)
	Close()
}

const (
	FLUENT = iota
	DEFAULT
)

// New function for Factory Pattern
func New(logEnable bool, logServiceType int, config *Fluent) ILog {
	if !logEnable {
		logServiceType = DEFAULT
	}

	switch logServiceType {
	case FLUENT:
		return NewFluent(config)
	default:
		return Default()
	}

	return nil
}
