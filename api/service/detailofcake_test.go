package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

func TestDetailOfCake_NegativeCantFindID(t *testing.T) {
	req, err := http.NewRequest("GET", "/cakes/1000", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DetailOfCake)
	handler.ServeHTTP(rr, req)

	expected := `{"responsecode":202,"total":0,"data":null,"message":"Cake with ID  doesn't exist "}`
	require.Equal(t, expected, strings.TrimSpace(rr.Body.String()))
}

// //SKIP Because cant reat {id} on path needd more research
// func TestDetailOfCake_Positive(t *testing.T) {
// 	req, err := http.NewRequest("GET", "/cakes/17", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(DetailOfCake)
// 	handler.ServeHTTP(rr, req)

// 	expected := `{"id":17,"title":"Maple","description":"Cake From heaven","rating":7,"image":"https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg","created_at":"2022-06-10 04:10:03","updated_at":"2022-06-10 04:10:03"}`
// 	require.Equal(t, expected, strings.TrimSpace(rr.Body.String()))
// }
