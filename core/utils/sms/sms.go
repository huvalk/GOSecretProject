package sms

import (
	"encoding/json"
	"github.com/kataras/golog"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type SMS struct {
	c http.Client
	exp *regexp.Regexp
	accountSid string
	authToken string
}

func NewSMS() *SMS {
	AccountSid, AuthToken := "AC74e2da512adfcb3015b636fee644dd6a", "65f0ac55235f41291bf37fc47d95ec58"
	return &SMS{
		c: http.Client{},
		exp: regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`),
		accountSid: AccountSid,
		authToken: AuthToken,
	}
}

func (sms *SMS) ValidatePhoneNumber(phone string) bool {
	return sms.exp.MatchString(phone)
}

func (sms *SMS) SendSMS(message, phone string) error {
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + sms.accountSid + "/Messages.json"
	msgData := url.Values{}
	msgData.Set("To","89032909821")
	msgData.Set("From","89032909821")
	msgData.Set("Body", "Message me")
	msgDataReader := *strings.NewReader(msgData.Encode())

	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(sms.accountSid, sms.authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make HTTP POST request and return message SID
	resp, _ := sms.c.Do(req)
	golog.Info(resp.StatusCode)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		golog.Info(data, err)
	}

	return nil
}
