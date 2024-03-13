package recaptcha

import (
	"encoding/json"
	"errors"
	"github.com/suhailgupta03/go-recaptcha/models"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type RequestParams struct {
	// The shared key between your site and reCAPTCHA.
	secret string
	// The user response token provided by the reCAPTCHA client-side integration on your site.
	response string
	// Optional. The user's IP address.
	remoteIp string
}

const (
	recaptchaURL               = "https://www.google.com/recaptcha/api/siteverify"
	requestToVerifyFailed      = "did not return 200"
	secretCannotBeEmpty        = "secret cannot be empty"
	responseTokenCannotBeEmpty = "user response token cannot be empty"
)

func New(secret, response, remoteIP string) *RequestParams {
	return &RequestParams{
		secret:   secret,
		response: response,
		remoteIp: remoteIP,
	}
}

func (params *RequestParams) Verify() (interface{}, error) {
	if params.secret == "" {
		return nil, errors.New(secretCannotBeEmpty)
	}

	if params.response == "" {
		return nil, errors.New(responseTokenCannotBeEmpty)
	}

	requestData := url.Values{}
	requestData.Set("secret", params.secret)
	requestData.Set("response", params.response)

	if params.remoteIp != "" {
		requestData.Set("remoteitp", params.remoteIp)
	}

	req, err := http.NewRequest("POST", recaptchaURL, strings.NewReader(requestData.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(requestToVerifyFailed)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var recaptchaResponse models.RecaptchaResponse
	if err := json.Unmarshal(b, &recaptchaResponse); err != nil {
		return err, nil
	}
	
	return &recaptchaResponse, nil

}
