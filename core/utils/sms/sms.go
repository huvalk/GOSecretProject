package sms

import (
	"github.com/kataras/golog"
	"github.com/vonage/vonage-go-sdk"
	"net/http"
	"os"
	"regexp"
)

type SMS struct {
	c          http.Client
	auth       *vonage.KeySecretAuth
	exp        *regexp.Regexp
	accountSid string
	authToken  string
}

func NewSMS() *SMS {
	API_KEY, API_SECRET := os.Getenv("API_KEY"), os.Getenv("API_SECRET")
	golog.Infof("SMS started: %s, %s", API_KEY, API_SECRET)
	return &SMS{
		c: http.Client{},

		auth:       vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET),
		exp:        regexp.MustCompile(`(0|\\+62|062|62)[0-9]+$`),
		accountSid: API_KEY,
		authToken:  API_SECRET,
	}
}

func (sms *SMS) ValidatePhoneNumber(phone string) bool {
	return sms.exp.MatchString(phone)
}

func (sms *SMS) SendSMS(message, phone string) error {
	smsClient := vonage.NewSMSClient(sms.auth)
	response, err := smsClient.Send("79032909821", phone, message, vonage.SMSOpts{})
	golog.Info(err)

	if response.Messages[0].Status == "0" {
		golog.Info("Message send, status: ", response.Messages[0].Status)
	}

	return err
}
