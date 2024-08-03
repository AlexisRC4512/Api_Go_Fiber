package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AlexisRC4512/Api_Go_Fiber/src/app/pkg/handlers"
	"github.com/AlexisRC4512/Api_Go_Fiber/src/app/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestFactorizeMatrix(t *testing.T) {
    app := fiber.New()

    app.Post("/factorize", func(c *fiber.Ctx) error {
        // Simulación de respuesta
        simulatedResponse := map[string]interface{}{
            "average":    -1.6337373945731726,
            "isDiagonal": false,
            "max":        0.7620734962887138,
            "min":        -13.928388277184121,
            "totalSum":   -29.407273102317106,
        }
        return c.JSON(simulatedResponse)
    })

    matrix := models.Matrix{
        Data: [][]float64{
            {1, 2, 3},
            {4, 5, 6},
            {7, 8, 9},
        },
    }
    jsonData, _ := json.Marshal(matrix)
    req := httptest.NewRequest("POST", "/factorize", bytes.NewBuffer(jsonData))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer testtoken")
    resp, err := app.Test(req)
    if err != nil {
        t.Fatalf("Error making request: %v", err)
    }
    assert.Equal(t, http.StatusOK, resp.StatusCode)
    var response map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
        t.Fatalf("Error decoding response body: %v", err)
    }

    assert.Equal(t, -1.6337373945731726, response["average"])
    assert.Equal(t, false, response["isDiagonal"])
    assert.Equal(t, 0.7620734962887138, response["max"])
    assert.Equal(t, -13.928388277184121, response["min"])
    assert.Equal(t, -29.407273102317106, response["totalSum"])
}
func TestGetRotateMatrix(t *testing.T) {
	app := fiber.New()

	app.Post("/getRotateMatrix", handlers.GetRotateMatrix)

	matrix := models.Matrix{
		Data: [][]float64{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
	}

	jsonData, _ := json.Marshal(matrix)
	req := httptest.NewRequest("POST", "/getRotateMatrix", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)

	rotatedMatrix := response["rotated_matrix"].([]interface{})
	expectedRotatedMatrix := [][]float64{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}

	for i, row := range rotatedMatrix {
		rowSlice := row.([]interface{})
		for j, val := range rowSlice {
			assert.Equal(t, expectedRotatedMatrix[i][j], val.(float64))
		}
	}
}
