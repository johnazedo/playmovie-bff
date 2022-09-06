package tools

import (
	"github.com/johnazedo/playmovie-bff/src/tools"
	"io"
	"log"
	"os"
	"testing"
)

func TestApiHandler(t *testing.T) {
	err := os.Setenv("BASE_API_URL", "")
	if err != nil {
		log.Fatal(err)
	}

	err = os.Setenv("API_KEY", "")
	if err != nil {
		log.Fatal(err)
	}

	response, err := tools.ApiHandler{}.Get("/movie/550")
	if err != nil {
		log.Fatal(err)
	}

	buffer, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	print(string(buffer))
}
