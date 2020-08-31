package test

// test file should ends with *_test.go
import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sanketgupta07/go-rest-example/controllers"
	"github.com/sanketgupta07/go-rest-example/util/constants"
)

// Test function should start with Test*
// This function will test homepage link
func TestHomeLink(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal("Error in running test for HomeLink. ", err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.HomeLink)
	handler.ServeHTTP(recorder, request)

	// http status check
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Request handler return wrong status. required %v got %v.", http.StatusOK, status)
	}

	expected := "Welcome home!"
	//Response test
	if recorder.Body.String() != expected {
		t.Errorf("Handler return wrong body. Required %v Got %v.", expected, recorder.Body.String())

	}
}

// Test Created user
func TestCreateUser(t *testing.T) {
	testUser := []byte(`{"Name": "Sonal","Age": 30,"Address": "Test Address 123"}`)
	request, err := http.NewRequest(http.MethodPost, constants.EndpointCreateUser, bytes.NewBuffer(testUser))

	if err != nil {
		t.Fatal("Error in running test for HomeLink. ", err)
	}

	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.CreateUser)
	handler.ServeHTTP(recorder, request)

	// http status check
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Request handler return wrong status. required %v got %v.", http.StatusOK, status)
	}

	expected := `{"Name":"Sonal","Age":30,"Address":"Test Address 123"}`
	//Response test
	if strings.TrimSpace(recorder.Body.String()) != expected {
		t.Errorf("Handler return wrong body. Required %v Got %v.", expected, recorder.Body.String())

	}
}
