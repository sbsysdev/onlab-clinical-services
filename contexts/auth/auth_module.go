package auth

import (
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authapp"
	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authctrls"
	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authinfra"
)

type AuthModule struct {
	Context    context.Context
	Connection *gorm.DB
	Router     *gin.RouterGroup
}

func (module AuthModule) LoadModule() error {
	// Configure repositories
	patientRepo := authinfra.PatientRepository{DB: module.Connection}
	locationRepo := authinfra.LocationRepository{DB: module.Connection}

	// TODO: Configure services

	// TODO: Configure event handlers

	// TODO: Configure controllers
	createPatientController := authctrls.CreatePatientController{
		CreatePatientUseCase: authapp.CreatePatientUseCase{
			PatientRepository:  patientRepo,
			LocationRepository: locationRepo,
		},
	}

	// Configure routes
	v1 := module.Router.Group("/v1")
	{
		v1.POST("/patients", createPatientController.Handle)
	}

	return nil
}
