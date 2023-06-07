package authctrls

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authapp"
)

type SignInPatientController struct {
	SignInPatientUseCase authapp.SignInPatientUseCase
}

func (ctrl SignInPatientController) Handle(ctx *gin.Context) {
	request := authapp.SignInPatientRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})

		return
	}

	signedIn, signInerr := ctrl.SignInPatientUseCase.Query(request)

	if signInerr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": signInerr.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   FromSignInUseCaseResponseToResponse(signedIn),
	})
}
