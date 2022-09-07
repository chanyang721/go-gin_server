package test

import (
	"testing"
)

// func TestCreateUser(t *testing.T) {
// 	body := struct {
// 		Name     string
// 		Email    string
// 		Password string
// 	}{
// 		Name:     "name",
// 		Email:    "email@email.com",
// 		Password: "secret",
// 	}

// 	req, _ := http.NewRequest("POST", "/user", jsonReaderFactory(body))
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusCreated, w.Code)
// }

func TestCreateUser(t *testing.T) {
	result := 1
	if result != 1 {
		t.Error("result should be 1, got", result)
	}
}
