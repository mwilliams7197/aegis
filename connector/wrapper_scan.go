package connector

import (
	"github.com/nortonlifelock/domain"
	"github.com/nortonlifelock/qualys"
	"strings"
	"time"
)

type scan struct {
	session *QsSession

	Name string ` json:"name,omitempty"`

	ScanID string `json:"scanId,omitempty"`

	// The identifier of the scan template
	TemplateID string `json:"templateId,omitempty"`

	// AssetGroupID holds the ID of the asset group that the scan is being executed against
	AssetGroupID string `json:"groupId,omitempty"`

	// The identifier of the scan engine.
	EngineIDs []string `json:"engineId,omitempty"`

	Created time.Time `json:"created"`

	Scheduled bool `json:"scheduled"`
}

func (s *scan) ID() string {
	return s.ScanID
}

func (s *scan) Title() string {
	return s.Name
}

func (s *scan) GroupID() string {
	return s.AssetGroupID
}

func (s *scan) Status() (status string, err error) {
	var scan qualys.ScanQualys
	scan, err = s.session.apiSession.GetScanByReference(s.ScanID)
	if err == nil {
		status = scan.Status.State

		if strings.ToLower(status) == strings.ToLower(domain.ScanFINISHED) {
			if scan.Processed == 0 {
				status = domain.ScanPROCESSING
			}
		}

		status = strings.ToLower(status)
	}

	return status, err
}
