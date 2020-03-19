package utils
import (
	"fmt"
	"net/http"
)

func CheckError(w *http.ResponseWriter, errors *[]error) bool {
	var x http.ResponseWriter = *w
	if len(*errors) > 0 {
		fmt.Println("Errors Detected from GORM")
		x.WriteHeader(http.StatusInternalServerError)
		for _, err := range *errors {
			x.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		}
		return true
	}
	return false
}