package auth

import (
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authapp"
	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authctrls"
	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authinfra"
	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/shared/shareddomain"
	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/shared/sharedinfra"
)

type AuthModule struct {
	Context                context.Context
	Connection             *gorm.DB
	SubscribeEvent         shareddomain.SubscribeDomainEvent
	PublishEvent           shareddomain.PublishDomainEvent
	ConfigureEventHandlers sharedinfra.ConfigureEventHandlers
	Router                 *gin.RouterGroup
}

func (module AuthModule) LoadModule() error {
	// Configure repositories

	patientRepo := authinfra.PatientRepository{DB: module.Connection}
	roleRepo := authinfra.RoleRepository{DB: module.Connection}
	locationRepo := authinfra.LocationRepository{DB: module.Connection}

	// TODO: Configure services

	// Configure event handlers

	sendWelcomeEmailOnPatientCreatedEventHandler := authapp.SendWelcomeEmailOnPatientCreatedEventHandler{
		SubscribeEvent: module.SubscribeEvent,
	}

	module.ConfigureEventHandlers(
		// Context
		module.Context,
		// Handlers
		sendWelcomeEmailOnPatientCreatedEventHandler,
	)

	// Configure controllers

	createPatientController := authctrls.CreatePatientController{
		CreatePatientUseCase: authapp.CreatePatientUseCase{
			// Repositories
			PatientRepository:  patientRepo,
			RoleRepository:     roleRepo,
			LocationRepository: locationRepo,
			// Publisher
			PublishEvent: module.PublishEvent,
		},
	}

	createOwnerController := authctrls.CreateOwnerController{
		CreateOwnerUseCase: authapp.CreateOwnerUseCase{
			// Repositories
			// Publisher
		},
	}

	readCountryListController := authctrls.ReadCountryListController{
		ReadCountryListUseCase: authapp.ReadCountryListUseCase{
			// Repositories
		},
	}

	// Configure routes

	v1 := module.Router.Group("/v1")

	// Sign up
	signUp := v1.Group("/sign-up")
	{
		signUp.POST("/patients", createPatientController.Handle)
		signUp.POST("/owners", createOwnerController.Handle)
		signUp.POST("/collaborators")
	}

	// Sign in
	signIn := v1.Group("/sign-in")
	{
		signIn.POST("/")
	}

	// Resoruces
	resources := v1.Group("/resources")
	{
		resources.GET("/countries", readCountryListController.Handle)
	}

	return nil
}
