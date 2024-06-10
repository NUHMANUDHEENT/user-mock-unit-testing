package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"userPage/controller"
	"userPage/database"
	"userPage/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func routerInit() *gin.Engine {
	r := gin.Default()
	return r
}

func TestSignUp(t *testing.T) {
	db := database.InitDatabase()
	router := routerInit()

	router.POST("/signup", controller.SignUp(db))

	input := models.Users{
		Name:     "nuhman",
		Email:    "nuhman@gmail.com",
		Password: "Nuhman@123",
	}
	reqBody, err := json.Marshal(input)
	if err != nil {
		t.Error(err)
	}

	req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	t.Log(res.Body)
	assert.Equal(t, http.StatusOK, res.Result().StatusCode)

}
