package domain

// Device defines the interface
type Device interface {
	// ID is the ID of the device as reported by the backend database of Aegis
	ID() string

	// SourceID is the ID of the device as reported by the scanner
	SourceID() *string

	OS() string
	HostName() string
	MAC() string
	IP() string

	// Region is the area that the device is stored in (if the device is a cloud device)
	Region() *string

	// InstanceID identifies which instance a device is (the the device is a cloud device)
	InstanceID() *string

	// GroupID identifies which asset group the device lies in (an identifier that originates from the scanner)
	GroupID() *string

	TrackingMethod() *string
}
