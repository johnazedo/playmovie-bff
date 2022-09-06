package hermes

import (
	"fmt"
	"net/http"
	"os"
)

func Get(url string) (*http.Response, error) {
	fullUrl := fmt.Sprintf("%s%s?api_key=%s", os.Getenv("BASE_API_URL"), url, os.Getenv("API_KEY"))
	return http.Get(fullUrl)
}
