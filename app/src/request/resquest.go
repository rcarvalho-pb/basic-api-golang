package request

import (
	"io"
	"net/http"
	"webapp/src/cookies"
)

func Request(r *http.Request, method, url string, data io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	cookie, _ := cookies.Read(r)
	req.Header.Add("Authorization", "Bearer "+cookie["token"])

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
