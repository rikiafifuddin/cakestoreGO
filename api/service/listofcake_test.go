package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

func TestGetListOfCake_Positive(t *testing.T) {
	req, err := http.NewRequest("GET", "/cakes/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListOfCake)
	handler.ServeHTTP(rr, req)

	expected := `{"responsecode":200,"total":1,"data":[{"id":17,"title":"Maple","description":"Cake From heaven","rating":7,"image":"https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg","created_at":"2022-06-10 04:10:03","updated_at":"2022-06-10 04:10:03"}],"message":"Success"}`
	require.Equal(t, expected, strings.TrimSpace(rr.Body.String()))
}
