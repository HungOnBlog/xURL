package logger

type LogHeader struct {
	RequestId string `json:"requestId"`
	IP        string `json:"ip"`
	Path      string `json:"path"`
	Method    string `json:"method"`
	ApiKey    string `json:"apiKey"`
}

type LogInfo struct {
	Header   LogHeader   `json:"header"`
	Body     interface{} `json:"body"`
	Addition interface{} `json:"addition"`
	Status   int         `json:"status"`
	Message  interface{} `json:"message"`
}
