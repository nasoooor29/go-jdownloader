package main

type MyConnectResponse struct {
	Sessiontoken string `json:"sessiontoken"`
	Regaintoken  string `json:"regaintoken"`
	Rid          int    `json:"rid"`
}

func (j *Jdownloader) Connect() (*MyConnectResponse, error) {
	r, err := SendAndDecrypt[MyConnectResponse](j, "POST", "/my/connect", map[string]string{
		"email":  j.Email,
		"appkey": j.AppKey,
	})
	if err != nil {
		return nil, err
	}
	j.SessionToken = []byte(r.Sessiontoken)
	j.RegainToken = []byte(r.Regaintoken)
	return r, nil
}
