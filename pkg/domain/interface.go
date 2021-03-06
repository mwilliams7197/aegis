package domain

import "time"

// DatabaseConnection holds a combination of the generated stored procedure calls and hand-written database calls
type DatabaseConnection interface {
	GeneratedDatabaseConnection

	GetDeviceByAssetOrgID(_AssetID string, OrgID string) (Device, error)
	GetDeviceByIP(_IP string, _OrgID string) (Device, error)
	//GetDeviceByIPMACAndRegion(_IP string, _MAC string, _Region string, _OrgID string) (Device, error)
	GetDeviceByCloudSourceIDAndIP(_IP string, _CloudSourceID string, _OrgID string) ([]Device, error)
	GetDeviceByScannerSourceID(_IP string, _GroupID string, _OrgID string) (Device, error)
	GetDeviceByInstanceID(_InstanceID string, _OrgID string) ([]Device, error)
	GetDevicesBySourceID(_SourceID string, _OrgID string) ([]Device, error)
	GetDevicesByCloudSourceID(_CloudSourceID string, _OrgID string) ([]Device, error)

	GetDetection(_SourceDeviceID string, _VulnerabilityID string, _Port int, _Protocol string) (Detection, error)
	GetDetectionBySourceVulnID(_SourceDeviceID string, _SourceVulnerabilityID string, _Port int, _Protocol string) (Detection, error)
	GetDetectionsForDevice(_DeviceID string) ([]Detection, error)
	GetDetectionsAfter(after time.Time, orgID string) (detections []Detection, err error)
	GetDetectionForGroupAfter(_LastUpdatedAfter time.Time, _LastFoundAfter time.Time, _OrgID string, inGroupID string, ticketInactiveKernels bool) ([]Detection, error)
	GetDetectionForDeviceID(inDeviceID string, _OrgID string, ticketInactiveKernels bool) ([]Detection, error)

	GetVulnReferences(vulnInfoID string, sourceID string) (references []VulnerabilityReference, err error)
	GetVulnRef(vulnInfoID string, sourceID string, reference string) (existing VulnerabilityReference, err error)

	GetVulnBySourceVulnID(_SourceVulnID string) (vulnerability Vulnerability, err error)
}
