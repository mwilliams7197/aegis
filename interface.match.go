package domain

// Match is an interface that holds a device/vulnerability combination, where a vulnerability scanner found
// the vulnerability on the device
type Match interface {
	IP() string
	Device() string
	Vulnerability() string
	GroupID() string
}
