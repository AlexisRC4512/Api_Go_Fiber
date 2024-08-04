package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AlexisRC4512/Api_Go_Fiber/src/app/pkg/models"
	"github.com/AlexisRC4512/Api_Go_Fiber/src/config"
	"github.com/AlexisRC4512/Api_Go_Fiber/src/services"
	"github.com/gofiber/fiber/v2"
)

// FactorizeMatrix Factorización QR de una matriz
//
//	@Summary		Factorización QR de una matriz
//	@Description	Recibe una matriz rectangular rota la matriz y devuelve la factorización QR de dicha matriz. Luego envía los datos resultantes a una API de Node.js.
//	@Tags			Matrices
//	@Accept			json
//	@Produce		json
//	@Param			matrix	body		models.Matrix			true	"Matriz rectangular a factorizar"
//	@Success		200		{object}	map[string]interface{}	"Factores Q y R de la matriz"//
//
// @Failure      400           {object}  models.ErrorResponse    "Invalid input"
// @Failure      401           {object}  models.ErrorResponse    "Unauthorized"
// @Failure      500           {object}  models.ErrorResponse    "Internal server error"
//
//	@Router			/factorize [post]
func FactorizeMatrix(c *fiber.Ctx) error {
	matrix := new(models.Matrix)
	if err := c.BodyParser(matrix); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: "Invalid input"})
	}
	matrixRotate := services.RotateMatrix(matrix.Data)

	Q, R, err := services.QRFactorization(matrixRotate)

	if len(Q) == 0 && len(R) == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "La matriz que ingreso no se puede factorizar en QR"})
	}
	fmt.Println("Matrix Rotate:")
	for _, row := range matrixRotate {
		fmt.Println(row)
	}
	fmt.Println("Matrix Q:")
	for _, row := range Q {
		fmt.Println(row)
	}

	fmt.Println("Matrix R:")
	for _, row := range R {
		fmt.Println(row)
	}
	data := map[string]interface{}{
		"Q": Q,
		"R": R,
	}
	fmt.Println(data)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to marshal data"})
	}

	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{Error: "Authorization token missing"})
	}

	req, err := http.NewRequest("POST", config.GetEndPoint(), bytes.NewBuffer(jsonData))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Error: "Failed to create request"})
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Error: "Failed to send data to Node.js API"})
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Error: "Node.js API responded with an error"})
	}

	var nodeResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&nodeResponse); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Error: "Failed to decode response from Node.js API"})
	}

	return c.JSON(nodeResponse)
}

// getRotateMatrix Rotación de una matriz
//
//	@Summary		Rotación de una matriz
//	@Description	Recibe una matriz rectangular, rota la matriz y devuelve la matriz rotada.
//	@Tags			Matrices
//	@Accept			json
//	@Produce		json
//	@Param			matrix	body		models.Matrix			true	"Matriz rectangular a rotar"
//	@Success		200		{object}	map[string]interface{}	"Matriz rotada"
//
// @Failure      400           {object}  models.ErrorResponse    "Invalid input"
// @Failure      401           {object}  models.ErrorResponse    "Unauthorized"
// @Failure      500           {object}  models.ErrorResponse    "Internal server error"
//
//	@Router			/getRotateMatrix [get]
func GetRotateMatrix(c *fiber.Ctx) error {
	matrix := new(models.Matrix)
	if err := c.BodyParser(matrix); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: "Invalid input"})
	}

	matrixRotate := services.RotateMatrix(matrix.Data)

	return c.JSON(fiber.Map{
		"rotated_matrix": matrixRotate,
	})
}
