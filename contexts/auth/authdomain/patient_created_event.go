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
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Surname    string    `json:"surname"`
	OccurredAt time.Time `json:"occurredAt"`
}

func (PatientCreatedEvent) EventName() shareddomain.DomainEvent {
	return EVENTS_PATIENT_CREATED_EVENT
}
func (event PatientCreatedEvent) EventOccurredAt() time.Time {
	return event.OccurredAt
}

func CreatePatientCreatedEvent(patient PatientEntity) PatientCreatedEvent {
	return PatientCreatedEvent{
		ID:         patient.ID,
		Name:       string(patient.Person.Name),
		Surname:    string(patient.Person.Surname),
		OccurredAt: time.Now(),
	}
}
