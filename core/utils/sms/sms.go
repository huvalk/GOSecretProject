package sms

import (
	"github.com/kataras/golog"
	"github.com/vonage/vonage-go-sdk"
	"net/http"
	"regexp"
)

type SMS struct {
	c http.Client
	auth *vonage.KeySecretAuth
	exp *regexp.Regexp
	accountSid string
	authToken string
}

func NewSMS() *SMS {
	//AccountSid, AuthToken := "AC74e2da512adfcb3015b636fee644dd6a", "8de195d83afb98a4287b95297db15ffc"
	API_KEY, API_SECRET := "873fed4d", "TcI0spIw2496O0Ql"
	return &SMS{
		c: http.Client{},

		auth: vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET),
		exp: regexp.MustCompile(`(0|\\+62|062|62)[0-9]+$`),
		accountSid: API_KEY,
		authToken: API_SECRET,
	}
}

func (sms *SMS) ValidatePhoneNumber(phone string) bool {
	return sms.exp.MatchString(phone)
}

func (sms *SMS) SendSMS(message, phone string) error {
	//urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + sms.accountSid + "/Messages.json"
	//msgData := url.Values{}
	//msgData.Set("To","79032909821")
	//msgData.Set("From","79032909821")
	//msgData.Set("Body", "Message me")
	//msgDataReader := *strings.NewReader(msgData.Encode())
	//
	//req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	//req.SetBasicAuth(sms.accountSid, sms.authToken)
	//req.Header.Add("Accept", "application/json")
	//req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//
	//// Make HTTP POST request and return message SID
	//resp, _ := sms.c.Do(req)
	//golog.Info(resp.StatusCode)
	//if resp.StatusCode >= 200 && resp.StatusCode < 300 {
	//	var data map[string]interface{}
	//	decoder := json.NewDecoder(resp.Body)
	//	err := decoder.Decode(&data)
	//	golog.Info(data, err)
	//}
	//
	//return nil

	smsClient := vonage.NewSMSClient(sms.auth)
	response, err := smsClient.Send("79032909821", phone, message, vonage.SMSOpts{})
	golog.Info(err)

	if response.Messages[0].Status == "0" {
		golog.Info("Message send, status: ", response.Messages[0].Status)
	}

	return err
}
