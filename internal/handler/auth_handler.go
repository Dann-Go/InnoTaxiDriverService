package handler

import (
	"net/http"

	"github.com/Dann-Go/InnoTaxiDriverService/internal/domain"
	"github.com/Dann-Go/InnoTaxiDriverService/internal/domain/responses"
	"github.com/gin-gonic/gin"
)

// SignUp godoc
// @Summary      SignUp
// @Description  create driver
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input body domain.Driver true "Driver info"
// @Success      200  {object}  responses.ServerResponse
// @Failure      400  {object}  responses.ServerResponse
// @Failure      404  {object}  responses.ServerResponse
// @Failure      500  {object}  responses.ServerResponse
// @Router       /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	json := domain.Driver{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewServerResponse(false, err.Error()))
		return
	}
	simplePassword := json.PasswordHash
	driver, err := h.driverService.CreateDriver(&json)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewServerResponse(false, err.Error()))
		return
	}

	token, err := h.authorizationService.GenerateToken(json.Phone, simplePassword)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewServerResponse(false, err.Error()))
		return
	}

	c.JSON(http.StatusOK, responses.NewServerResponse(true, map[string]interface{}{
		"driver": driver,
		"token":  token,
	}))
}

type signInInput struct {
	Phone        string `db:"phone" json:"phone" binding:"required"`
	PasswordHash string `db:"password_hash" json:"passwordHash" binding:"required"`
}

// SignIn godoc
// @Summary      SignIn
// @Description  Check if driver exists and generate token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input body signInInput true "Driver signIN info"
// @Success      200  {object}  responses.ServerResponse
// @Failure      400  {object}  responses.ServerResponse
// @Failure      404  {object}  responses.ServerResponse
// @Failure      500  {object}  responses.ServerResponse
// @Router       /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	json := signInInput{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewServerResponse(false, err.Error()))
		return
	}
	token, err := h.authorizationService.GenerateToken(json.Phone, json.PasswordHash)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewServerResponse(false, err.Error()))
		return
	}

	driverFull, err := h.driverService.GetDriverByPhone(json.Phone)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewServerResponse(false, err.Error()))
		return
	}
	driverResponse := domain.DriverResponse{
		ID:       driverFull.ID,
		Name:     driverFull.Name,
		Phone:    driverFull.Phone,
		Email:    driverFull.Email,
		Rating:   driverFull.Rating,
		TaxiType: driverFull.TaxiType,
	}

	c.JSON(http.StatusOK, responses.NewServerResponse(true, map[string]interface{}{
		"driver": driverResponse,
		"token":  token,
	}))
}
