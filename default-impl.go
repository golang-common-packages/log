package log

// DefaultClient manage all slack action
type DefaultClient struct{}

// Default function return empty struct
func Default() ILog {
	return &DefaultClient{}
}

// Write functions return empty write function (Default)
func (d *DefaultClient) Write(p []byte) (n int, err error) {
	return len(p), nil
}

// Close functions return empty close function (Default)
func (d *DefaultClient) Close() {

}
