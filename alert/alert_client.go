package alert

import (
	"github.com/libopenstorage/openstorage/api"
)

// AlertInstance is a singleton which should be used to raise/clear alert
type alertInstance struct {
	NodeID    string
	ClusterID string
	Version   string
	kva       AlertClient
}

// Singleton AlertInstance
var inst *alertInstance

// Clear clears an alert
func (ai *alertInstance) Clear(resourceType api.ResourceType, alertID int64) error {
	return ai.kva.Clear(resourceType, alertID)
}

// Alarm raises an alert with severity : ALARM
func (ai *alertInstance) Alarm(alertType int64, msg string, resourceType api.ResourceType, resourceID string) (int64, error) {
	alert, err := ai.raiseAlert(alertType, msg, resourceType, resourceID, api.SeverityType_SEVERITY_TYPE_ALARM)
	return alert.Id, err
}

// Notify raises an alert with severity : NOTIFY
func (ai *alertInstance) Notify(alertType int64, msg string, resourceType api.ResourceType, resourceID string) (int64, error) {
	alert, err := ai.raiseAlert(alertType, msg, resourceType, resourceID, api.SeverityType_SEVERITY_TYPE_NOTIFY)
	return alert.Id, err
}

// Warn raises an alert with severity : WARNING
func (ai *alertInstance) Warn(alertType int64, msg string, resourceType api.ResourceType, resourceID string) (int64, error) {
	alert, err := ai.raiseAlert(alertType, msg, resourceType, resourceID, api.SeverityType_SEVERITY_TYPE_WARNING)
	return alert.Id, err
}

// Alert :  Keeping this function for backward compatibility until we remove all calls to this function
func (ai *alertInstance) Alert(name string, msg string) error {
	// Do nothing
	return nil
}

// Sets the new singleton alert instance
func newAlertInstance(nodeID, clusterID, version string, kva AlertClient) {
	inst = &alertInstance{
		NodeID:    nodeID,
		ClusterID: clusterID,
		Version:   version,
		kva:       kva,
	}
}

// Instance returns the singleton AlertInstance
func instance() *alertInstance {
	return inst
}

func (ai *alertInstance) raiseAlert(alertType int64, msg string, resourceType api.ResourceType, resourceID string, severity api.SeverityType) (*api.Alert, error) {
	alert := api.Alert{
		AlertType:  alertType,
		Resource:   resourceType,
		ResourceId: resourceID,
		Severity:   severity,
		Message:    msg,
	}
	err := ai.kva.Raise(&alert)
	return &alert, err
}