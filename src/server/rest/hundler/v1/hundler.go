package v1

import (
	"auth-microservice/src/helpers"
	"auth-microservice/src/log"
	"auth-microservice/src/models/apiResponse"
	"auth-microservice/src/models/forms"
	"auth-microservice/src/usecases"
	"github.com/gin-gonic/gin"
)

type HundlerInterface interface {
	Register(c *gin.Context)
	ConfirmRegister(c *gin.Context)
	Login(c *gin.Context)
	Verify(c *gin.Context)
	UpdateTokens(c *gin.Context)
}

type Hundler struct {
	usecase usecases.UseCaseInterface
	logger  *log.Logger
}

func NewHundler(usecase usecases.UseCaseInterface, logger *log.Logger) *Hundler {
	return &Hundler{usecase, logger}
}

func (h *Hundler) Register(c *gin.Context) {
	form := forms.NewRegisterFormIngot()
	if err := c.ShouldBind(form); err != nil {
		helpers.ErrorResponse(c, err)

		return
	}
	if err := form.Validate(); err != nil {
		helpers.ErrorResponse(c, err)

		return
	}

	token, err := h.usecase.RegisterUser(form.Email, form.Password)
	if err != nil {
		helpers.ErrorResponse(c, err)

		return
	}

	helpers.SuccessResponse(c, apiResponse.NewTokenApiResponse(token))
}

func (h *Hundler) ConfirmRegister(c *gin.Context) {
	form := forms.NewTokenFormIngot()
	if err := c.ShouldBind(form); err != nil {
		helpers.ErrorResponse(c, err)

		return
	}
	if err := form.Validate(); err != nil {
		helpers.ErrorResponse(c, err)

		return
	}

	user, err := h.usecase.ConfirmRegister(form.Token)
	if err != nil {
		helpers.ErrorResponse(c, err)

		return
	}

	helpers.SuccessResponse(c, apiResponse.NewUserApiResponse(user.ID, user.Email))
}

func (h *Hundler) Login(c *gin.Context) {
	form := forms.NewLoginFormIngot()
	if err := c.ShouldBind(form); err != nil {
		helpers.ErrorResponse(c, err)

		return
	}
	if err := form.Validate(); err != nil {
		helpers.ErrorResponse(c, err)

		return
	}

	tokens, err := h.usecase.Login(form.Email, form.Password)
	if err != nil {
		helpers.ErrorResponse(c, err)

		return
	}

	helpers.SuccessResponse(c, apiResponse.NewTokensApiResponse(tokens.AccessToken, tokens.RefreshToken))
}

func (h *Hundler) Verify(c *gin.Context) {
	form := forms.NewTokenFormIngot()
	if err := c.ShouldBind(form); err != nil {
		helpers.ErrorResponse(c, err)

		return
	}
	if err := form.Validate(); err != nil {
		helpers.ErrorResponse(c, err)

		return
	}

	isTokenValid, err := h.usecase.Verify(form.Token)
	if err != nil {
		helpers.ErrorResponse(c, err)

		return
	}

	helpers.SuccessResponse(c, apiResponse.NewSuccessApiResponse(isTokenValid, ""))
}

func (h *Hundler) UpdateTokens(c *gin.Context) {
	form := forms.NewTokenFormIngot()
	if err := c.ShouldBind(form); err != nil {
		helpers.ErrorResponse(c, err)

		return
	}
	if err := form.Validate(); err != nil {
		helpers.ErrorResponse(c, err)

		return
	}

	tokens, err := h.usecase.GetTokensByRefresh(form.Token)
	if err != nil {
		helpers.ErrorResponse(c, err)

		return
	}

	helpers.SuccessResponse(c, apiResponse.NewTokensApiResponse(tokens.AccessToken, tokens.RefreshToken))
}
