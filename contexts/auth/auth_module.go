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
	"github.com/OnLab-Clinical/onlab-clinical-services/middlewares"
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

	roleRepo := authinfra.RoleRepository{DB: module.Connection}
	locationRepo := authinfra.LocationRepository{DB: module.Connection}
	patientRepo := authinfra.PatientRepository{
		DB:                 module.Connection,
		LocationRepository: locationRepo,
		RoleRepository:     roleRepo,
	}

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
			LocationRepository: locationRepo,
		},
	}

	signInPatientController := authctrls.SignInPatientController{
		SignInPatientUseCase: authapp.SignInPatientUseCase{
			// Repositories
			PatientRepository: patientRepo,
		},
	}

	refreshPatientTokenController := authctrls.RefreshPatientTokenController{
		RefreshPatientTokenUseCase: authapp.RefreshPatientTokenUseCase{
			// Repositories
			PatientRepository: patientRepo,
		},
	}

	recoveryPatientController := authctrls.RecoveryPatientController{
		RecoveryPatientUseCase: authapp.RecoveryPatientUseCase{
			// Repositories
			PatientRepository: patientRepo,
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
		signIn.POST("/patients", signInPatientController.Handle)
	}

	// Recovery
	recovery := v1.Group("/recovery")
	{
		recovery.POST("/patients", recoveryPatientController.Handle)
	}

	// Resoruces
	resources := v1.Group("/resources")
	{
		resources.GET("/countries", readCountryListController.Handle)
		resources.POST("/refresh", refreshPatientTokenController.Handle)
		resources.GET("/protected", middlewares.CheckTokenMiddleware(), func(ctx *gin.Context) {
			patientId := ctx.MustGet("patientId").(string)

			ctx.JSON(200, gin.H{
				"status":  true,
				"message": "Hello from protected " + patientId,
			})
		})
	}

	return nil
}
