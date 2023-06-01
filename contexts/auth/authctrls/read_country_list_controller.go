package authctrls

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/auth/authapp"
)

type ReadCountryListController struct {
	ReadCountryListUseCase authapp.ReadCountryListUseCase
}

func (ctrl ReadCountryListController) Handle(ctx *gin.Context) {
	countryList, countryListErr := ctrl.ReadCountryListUseCase.Query()

	if countryListErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": countryListErr.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   countryList,
	})
}
