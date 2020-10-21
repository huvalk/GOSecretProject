package baseModels

type CodeConfirmation struct {
	Code string `json:"code"`
}

type SessionConfirmation struct {
	Session string `json:"session"`
}
