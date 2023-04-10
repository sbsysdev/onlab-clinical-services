package authctrls

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authapp"
)

type CreatePatientController struct {
	CreatePatientUseCase authapp.CreatePatientUseCase
}

func (ctrl CreatePatientController) Handle(ctx *gin.Context) {
	request := authapp.CreatePatientRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		// lang := ctx.GetHeader("Accept-Language")

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			//"request": utils.RequestValidator(err, errorMsgs),
		})

		return
	}

	/* file, _ := ctx.FormFile("file")

	f, _ := file.Open()

	println(f, file) */

	if err := ctrl.CreatePatientUseCase.Command(request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": true,
		//"message": shared.MESSAGES_CREATED,
	})
}
