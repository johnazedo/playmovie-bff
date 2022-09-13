package hermes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Client interface {
	Get(model interface{}, path string) error
}

type ClientConcrete struct{}

func NewClient() Client {
	return &ClientConcrete{}
}

func (c *ClientConcrete) Get(model any, path string) error {
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

//func (d *decodeState) unmarshal(v any) error {
//	rv := reflect.ValueOf(v)
//	if rv.Kind() != reflect.Pointer || rv.IsNil() {
//		return &InvalidUnmarshalError{reflect.TypeOf(v)}
//	}
//
//	d.scan.reset()
//	d.scanWhile(scanSkipSpace)
//	// We decode rv not rv.Elem because the Unmarshaler interface
//	// test must be applied at the top level of the value.
//	err := d.value(rv)
//	if err != nil {
//		return d.addErrorContext(err)
//	}
//	return d.savedError
//}

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
