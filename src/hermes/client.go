package hermes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Client struct{}

func (c *Client) Get(model interface{}, path string) (interface{}, error) {
	var client http.Client
	url := getFullUrl(path)
	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	buffer, err := readResponse(response)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buffer, model)
	if err != nil {
		return nil, err
	}
	return model, err
}

func readResponse(response *http.Response) ([]byte, error) {
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		buffer, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return buffer, nil
	}
	return nil, NewHttpError(response.StatusCode)
}

func getFullUrl(path string) string {
	return fmt.Sprintf("%s%s?api_key=%s", os.Getenv("BASE_API_URL"), path, os.Getenv("API_KEY"))
}
