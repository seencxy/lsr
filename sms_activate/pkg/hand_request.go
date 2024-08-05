package pkg

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const (
	EmptyString = ""
)

// HandRequestUnmarshalData unmarshal the data from the request body.
func HandRequestUnmarshalData(client http.Client, ctx context.Context, req *http.Request, data interface{}) error {
	// check context
	if ctx.Err() != nil {
		return ctx.Err()
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return errors.New("status code is not 200")
	}

	all, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var errorRes ErrorRes
	_ = json.NewDecoder(res.Body).Decode(&errorRes)

	if errorRes.Status != "" {
		return errors.New(errorRes.Error)
	}

	return json.Unmarshal(all, data)
}

// HandRequestData handles the request data.
func HandRequestData(client http.Client, ctx context.Context, req *http.Request) (string, error) {
	// check context
	if ctx.Err() != nil {
		return EmptyString, ctx.Err()
	}

	res, err := client.Do(req)
	if err != nil {
		return EmptyString, err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	resDataBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return EmptyString, err
	}

	return string(resDataBytes), nil
}

type ErrorRes struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}
