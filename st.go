package main

import "net/http"

type Jdownloader struct {
	Email        string
	Password     string
	SessionToken []byte
	RegainToken  []byte
	// CurrentRID     string
	NextRID        string
	AppKey         string
	Connected      bool
	CurrentPayload []byte
	Client         *http.Client
}

const BASE_URL = "https://api.jdownloader.org"

func NewJdownloader(e, p string) *Jdownloader {
	return &Jdownloader{
		Email:        e,
		Password:     p,
		SessionToken: []byte{},
		// CurrentRID:     "0",
		NextRID:        "0",
		AppKey:         "Nasoooor_JDOWNLOADER",
		Connected:      false,
		CurrentPayload: nil,
		Client:         http.DefaultClient,
	}
}
