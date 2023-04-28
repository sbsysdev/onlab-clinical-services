package authapp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"
	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/shared/shareddomain"
)

type SendWelcomeEmailOnPatientCreatedEventHandler struct {
	// Subscribe Event
	SubscribeEvent shareddomain.SubscribeDomainEvent
	// Services
}

func (handler SendWelcomeEmailOnPatientCreatedEventHandler) Handle(ctx context.Context) {
	messages, err := handler.SubscribeEvent(ctx, authdomain.EVENTS_PATIENT_CREATED_EVENT)

	if err != nil {
		panic(err)
	}

	for msg := range messages {
		var patientCreated authdomain.PatientCreatedEvent

		json.Unmarshal([]byte(msg.Payload), &patientCreated)

		if err := handler.sendWelcomeEmail(patientCreated); err != nil {
			panic(err)
		}

		msg.Ack()
	}
}

func (handler SendWelcomeEmailOnPatientCreatedEventHandler) sendWelcomeEmail(patient authdomain.PatientCreatedEvent) error {
	fmt.Println("SendWelcomeEmailOnPatientCreatedEventHandler", patient)
	fmt.Println("Name", patient.EventName())
	fmt.Println("Meta", patient.EventMetadata())
	fmt.Println("Ocurred", patient.OccurredAt)

	return nil
}
