package model

import "encoding/json"

type Request struct {
	Type            string
	RequireResponse bool
}

var (
	R_NAME = Request{
		Type:            "NAME",
		RequireResponse: true,
	}
	R_TALK = Request{
		Type:            "TALK",
		RequireResponse: true}
	R_WHISPER = Request{
		Type:            "WHISPER",
		RequireResponse: true}
	R_VOTE = Request{
		Type:            "VOTE",
		RequireResponse: true}
	R_DIVINE = Request{
		Type:            "DIVINE",
		RequireResponse: true}
	R_GUARD = Request{
		Type:            "GUARD",
		RequireResponse: true}
	R_ATTACK = Request{
		Type:            "ATTACK",
		RequireResponse: true}
	R_INITIALIZE = Request{
		Type:            "INITIALIZE",
		RequireResponse: false}
	R_DAILY_INITIALIZE = Request{
		Type:            "DAILY_INITIALIZE",
		RequireResponse: false}
	R_DAILY_FINISH = Request{
		Type:            "DAILY_FINISH",
		RequireResponse: false}
	R_FINISH = Request{
		Type:            "FINISH",
		RequireResponse: false}
)

func (r Request) String() string {
	return r.Type
}

func (r Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func RequestFromString(s string) Request {
	switch s {
	case "NAME":
		return R_NAME
	case "TALK":
		return R_TALK
	case "WHISPER":
		return R_WHISPER
	case "VOTE":
		return R_VOTE
	case "DIVINE":
		return R_DIVINE
	case "GUARD":
		return R_GUARD
	case "ATTACK":
		return R_ATTACK
	case "INITIALIZE":
		return R_INITIALIZE
	case "DAILY_INITIALIZE":
		return R_DAILY_INITIALIZE
	case "DAILY_FINISH":
		return R_DAILY_FINISH
	case "FINISH":
		return R_FINISH
	}
	return Request{}
}
