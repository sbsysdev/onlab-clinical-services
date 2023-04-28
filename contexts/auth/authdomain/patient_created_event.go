package authdomain

import (
	"time"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/shared/shareddomain"
)

const (
	EVENTS_PATIENT_CREATED_EVENT shareddomain.DomainEvent = "EVENTS_PATIENT_CREATED_EVENT"
)

// Auth events

// PatientCreatedEvent
type PatientCreatedEvent struct {
	Patient    PatientEntity `json:"patient"`
	OccurredAt time.Time     `json:"occurredAt"`
}

func (PatientCreatedEvent) EventName() shareddomain.DomainEvent {
	return EVENTS_PATIENT_CREATED_EVENT
}
func (event PatientCreatedEvent) EventMetadata() interface{} {
	return event.Patient
}
func (event PatientCreatedEvent) EventOccurredAt() time.Time {
	return event.OccurredAt
}

func CreatePatientCreatedEvent(patient PatientEntity) PatientCreatedEvent {
	return PatientCreatedEvent{
		Patient:    patient,
		OccurredAt: time.Now().UTC(),
	}
}
