package authctrls

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authapp"
)

type CreateOwnerController struct {
	CreateOwnerUseCase authapp.CreateOwnerUseCase
}

func (ctrl CreateOwnerController) Handle(ctx *gin.Context) {
	request := authapp.CreateOwnerRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		// lang := ctx.GetHeader("Accept-Language")

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})

		return
	}

	/* file, _ := ctx.FormFile("file")

	f, _ := fileauthapp

	println(f, file) */

	if err := ctrl.CreateOwnerUseCase.Command(request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": true,
	})
}
