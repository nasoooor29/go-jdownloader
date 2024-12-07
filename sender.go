package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (j *Jdownloader) SendReq(meth, endpoint string, data map[string]string) (*http.Response, error) {
	u := fmt.Sprintf("%v%v", BASE_URL, endpoint)
	req, err := http.NewRequest(meth, u, NewBuf(j.CurrentPayload))
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for k, v := range data {
		q.Add(k, v)
	}
	q.Add("rid", j.NextRID)
	beforeEncUrl := req.URL.Path + "?" + q.Encode()
	fmt.Printf("beforeEncUrl: %v\n", beforeEncUrl)
	if !j.Connected {
		sec := secretCreate(j.Email, j.Password, "server")
		j.SessionToken = sec
	}

	sig := generateHMACSignature([]byte(beforeEncUrl), j.SessionToken)

	q.Add("signature", sig)
	req.URL.RawQuery = q.Encode()

	res, err := j.Client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func SendAndDecrypt[T any](j *Jdownloader, meth, endpoint string, data map[string]string) (*T, error) {
	res, err := j.SendReq(meth, endpoint, data)
	if err != nil {
		return nil, err
	}

	d, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	decrypted, err := decrypt(string(j.SessionToken), string(d))
	if err != nil {
		return nil, err
	}
	n := new(T)
	err = json.Unmarshal([]byte(decrypted), n)
	if err != nil {
		return nil, err
	}
	return n, nil

}

func NewBuf(b []byte) io.Reader {
	if b == nil {
		return nil
	}
	return bytes.NewBuffer(b)
}
