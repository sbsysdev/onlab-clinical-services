package authctrls

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authapp"
)

type RecoveryPatientController struct {
	RecoveryPatientUseCase authapp.RecoveryPatientUseCase
}

func (ctrl RecoveryPatientController) Handle(ctx *gin.Context) {
	request := authapp.RecoveryPatientRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		// lang := ctx.GetHeader("Accept-Language")

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})

		return
	}

	if err := ctrl.RecoveryPatientUseCase.Command(request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}
