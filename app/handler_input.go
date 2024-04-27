package pasha_bot

import "strings"

type MatchedCard struct {
	Name string `json:"name" binding:"required"`
}

type LastFaceEvent struct {
	Thumbnail string `json:"thumbnail" binding:"required"`
	FullFrame string `json:"fullframe" binding:"required"`
}

type Input struct {
	Date    string        `json:"created_date" binding:"required"`
	Persona MatchedCard   `json:"matched_card" binding:"required"`
	Image   LastFaceEvent `json:"last_face_event" binding:"required"`
}

func (i *Input) Format() {
	i.Image.Thumbnail = strings.Replace(i.Image.Thumbnail, "http://192.168.1.58/uploads", "/img", 1)
	i.Image.FullFrame = strings.Replace(i.Image.FullFrame, "http://192.168.1.58/uploads", "/img", 1)
}
