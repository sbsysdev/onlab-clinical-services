package authinfra

import (
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/OnLab-Clinical/onlab-clinical-services/utils"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authdomain"
)

type PatientEmailService struct{}

func (service PatientEmailService) SendWelcomeEmailToPatient(patient authdomain.PatientEntity) error {
	return nil
}
func (service PatientEmailService) SendRecoveryEmailToPatient(patient authdomain.PatientEntity, token string) error {
	from := mail.NewEmail("OnLab-Clinical", utils.GetEnv("SENDGRID_SENDER", ""))
	subject := "ObLab-Clinical recovery email"
	to := mail.NewEmail(fmt.Sprintf("%s %s", patient.Person.Name, patient.Person.Surname), string(patient.Contacts.Email))

	plainTextContent := fmt.Sprintf("%s/%s", utils.GetEnv("RECOVERY_CLIENT", ""), token)

	htmlContent := fmt.Sprintf(`<a href="%s/%s" target="_blank">Click Here for Recovery Password</a>`, utils.GetEnv("RECOVERY_CLIENT", ""), token)

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(utils.GetEnv("SENDGRID_API_KEY", ""))

	_, err := client.Send(message)

	if err != nil {
		return err
	}

	return nil
}
