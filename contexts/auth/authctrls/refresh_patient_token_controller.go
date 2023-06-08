package authctrls

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authapp"
)

type RefreshPatientTokenController struct {
	RefreshPatientTokenUseCase authapp.RefreshPatientTokenUseCase
}

func (ctrl RefreshPatientTokenController) Handle(ctx *gin.Context) {
	request := authapp.RefreshPatientTokenRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})

		return
	}

	refreshed, refreshedErr := ctrl.RefreshPatientTokenUseCase.Query(request)

	if refreshedErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": refreshedErr.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   refreshed,
	})
}
