package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"test/controller"
	"test/database"
	"test/routers"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	routers.SetupRouter(router)
	return router
}

func TestSignUp(t *testing.T) {
	// Initialize the mock database
	mockDB := new(database.MockDatabase)
	controller.DB = mockDB

	// Define expected behavior
	mockDB.On("CreateUser", "testuser", "password123").Return(nil)

	router := setupRouter()

	signUpData := controller.SignUpRequest{
		Username: "testuser",
		Password: "password123",
	}
	jsonData, _ := json.Marshal(signUpData)

	req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	fmt.Println("----", recorder)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "User signed up successfully")

	// Assert that the expectations were met
	mockDB.AssertExpectations(t)
}

// func TestSignUpPostError(t *testing.T) {
//     // Initialize the mock database
//     mockDB := new(database.MockDatabase)
//     controller.DB = mockDB

//     // Define expected behavior
//     mockDB.On("CreateUser", "testuser", "password123").Return(mock.Anything)

//     router := setupRouter()

//     signUpData := controller.SignUpRequest{
//         Username: "testuser",
//         Password: "password123",
//     }
//     jsonData, _ := json.Marshal(signUpData)

//     req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(jsonData))
//     req.Header.Set("Content-Type", "application/json")

//     recorder := httptest.NewRecorder()
//     router.ServeHTTP(recorder, req)

//     assert.Equal(t, http.StatusInternalServerError, recorder.Code)
//     assert.Contains(t, recorder.Body.String(), "error")

//     // Assert that the expectations were met
//     mockDB.AssertExpectations(t)
// }
