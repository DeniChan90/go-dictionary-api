package models

import (
	"golang.org/x/text/language"
	"time"
)

type Language struct {
	Code   language.Tag `json:"code"`
	Symbol string       `json:"symbol"`
	Name   string       `json:"name"`
}

type TranslateRequest struct {
	Text string `bson:"text,omitempty"`
	From string `bson:from,omitempty`
	To   string `bson:to,omitempty`
}

type TranslateResponse struct {
	Lang          string  `bson:"lang,omitempty"`
	Text          string  `bson:"text,omitempty"`
	Pronunciation *string `bson:"pronunciation,omitempty"`
}

type Tranlsation struct {
	Word_id       string `json:"word_id"`
	User_id       string `json: "user_id"`
	Lang          string `json: "lang"`
	Text          string `json: "text"`
	Pronunciation string `json:"pronunciation"`
}

type Word struct {
	Word_id    string    `bson:"_id"`
	Created_at time.Time `json: "created_at"`
	User_id    string    `json: "user_id"`
}
