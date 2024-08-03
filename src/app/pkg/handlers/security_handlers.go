package handlers

import (
	"time"

	"github.com/AlexisRC4512/Api_Go_Fiber/src/app/pkg/models"
	"github.com/AlexisRC4512/Api_Go_Fiber/src/app/pkg/repository"
	"github.com/AlexisRC4512/Api_Go_Fiber/src/config"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
)

// Login inicia sesión con datos predeterminados
//
// @Summary      Inicio de Sesión
// @Description  Recibe un json con email y password para iniciar sesión y devuelve un token JWT.
// @Tags         Login
// @Accept       json
// @Produce      json
// @Param        loginRequest  body      models.LoginRequest     true  "Datos de inicio de sesión"
// @Success      200           {object}  models.LoginResponse    "Token JWT"
// @Failure      400           {object}  models.ErrorResponse    "Invalid input"
// @Failure      401           {object}  models.ErrorResponse    "Unauthorized"
// @Failure      500           {object}  models.ErrorResponse    "Internal server error"
// @Router       /login [post]
func Login(c *fiber.Ctx) error {
	loginRequest := new(models.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: err.Error(),
		})
	}
	user, err := repository.FindByCredentials(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: err.Error(),
		})
	}

	day := time.Hour * 24

	claims := jtoken.MapClaims{
		"ID":    user.ID,
		"email": user.Email,
		"fav":   user.FavoritePhrase,
		"exp":   time.Now().Add(day * 1).Unix(),
	}

	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.GetSecret()))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(models.LoginResponse{
		Token: t,
	})
}
