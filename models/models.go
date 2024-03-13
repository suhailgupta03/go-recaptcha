package models

import "time"

// RecaptchaResponse https://developers.google.com/recaptcha/docs/verify#api-response
type RecaptchaResponse struct {
	Success     bool      `json:"success"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes,omitempty"`
	Action      string    `json:"action,omitempty"`
}

type Provider interface {
	Verify() (interface{}, error)
}
