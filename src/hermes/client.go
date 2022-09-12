package hermes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Model interface{}
type Client struct{}

func (c *Client) Get(model Model, path string) error {
	var client http.Client
	url := getFullUrl(path)
	response, err := client.Get(url)
	if err != nil {
		return err
	}
	buffer, err := readResponse(response)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buffer, model)
	if err != nil {
		return err
	}
	return nil
}

func readResponse(response *http.Response) ([]byte, error) {
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		buffer, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, NewHttpError(err.Error(), response.StatusCode)
		}
		return buffer, nil
	}
	return nil, NewHttpError("error to read response", response.StatusCode)
}

func getFullUrl(path string) string {
	return fmt.Sprintf("%s%s?api_key=%s", os.Getenv("BASE_API_URL"), path, os.Getenv("API_KEY"))
}
