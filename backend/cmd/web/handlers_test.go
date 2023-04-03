package main

import (
	"github.com/wederer/go-chi-demo/internal"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func executeRequest(req *http.Request, s *Server) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func (s *Server) setupMockDatabase() {
	s.Books2 = &internal.MockBooks2{}
}

func TestServer_HelloWorld(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()

	req, _ := http.NewRequest("GET", "/", nil)

	response := executeRequest(req, s)

	checkResponseCode(t, http.StatusOK, response.Code)

	require.Equal(t, "Hello World!", response.Body.String())
}

func TestServer_GetBook(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()
	s.setupMockDatabase()

	req, _ := http.NewRequest("GET", "/books/123", nil)
	response := executeRequest(req, s)
	checkResponseCode(t, http.StatusNotFound, response.Code)

	req, _ = http.NewRequest("GET", "/books/correct_key", nil)
	response = executeRequest(req, s)
	checkResponseCode(t, http.StatusOK, response.Code)
	want := "{correct_key some_title 42}"
	got := response.Body.String()
	if want != got {
		t.Errorf("want is not equal to got. want: %v, got: %v", want, got)
	}
}

func TestServer_GetBooks(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()
	s.setupMockDatabase()

	req, _ := http.NewRequest("GET", "/books", nil)
	response := executeRequest(req, s)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestServer_GetProtectedData(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()

	req, _ := http.NewRequest("GET", "/protected", nil)

	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9."+
		"eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ."+
		"XbPfbIHMI6arZ3Y922BhjWgQzWXcXNrz0ogtVhfEd2o")
	response := executeRequest(req, s)
	checkResponseCode(t, http.StatusOK, response.Code)
	require.Equal(t, "Secret Protected Info", response.Body.String())

	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9."+
		"eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ."+
		"wrong_signature")
	response = executeRequest(req, s)
	checkResponseCode(t, http.StatusUnauthorized, response.Code)
}

func TestServer_GetAdminInfo(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()

	req, _ := http.NewRequest("GET", "/admin/info", nil)

	// claim admin:true in JWT
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9."+
		"eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJhZG1pbiI6dHJ1ZX0."+
		"yfHYiOV8YznrNOVNYGrldeMhUZqu0ZDqi4t8oI1dhoU")
	response := executeRequest(req, s)
	checkResponseCode(t, http.StatusOK, response.Code)
	require.Equal(t, "Even more secret info that only admins can read.", response.Body.String())

	// no admin claim in JWT
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9."+
		"eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ."+
		"XbPfbIHMI6arZ3Y922BhjWgQzWXcXNrz0ogtVhfEd2o")
	response = executeRequest(req, s)
	checkResponseCode(t, http.StatusForbidden, response.Code)

	// admin: false in JWT
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9."+
		"eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJhZG1pbiI6ZmFsc2V9."+
		"LBqPKgILN3MwO-IjVfqtzLjv1kLpq4ReXgu653s2eLc")
	response = executeRequest(req, s)
	checkResponseCode(t, http.StatusForbidden, response.Code)
}
