package pasha_bot

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"pasha_bot/pkg/request"
	"strings"
)

type Faces struct {
	results []Face
}

type Face struct {
	thumbnail string
}

type MatchedCard struct {
	Name   string `json:"name" binding:"required"`
	CardId int    `json:"id" binding:"required"`
}

type LastFaceEvent struct {
	Thumbnail string `json:"thumbnail" binding:"required"`
	FullFrame string `json:"fullframe" binding:"required"`
	Original  string
}

type Webhook struct {
	Date    string        `json:"created_date" binding:"required"`
	Persona MatchedCard   `json:"matched_card" binding:"required"`
	Image   LastFaceEvent `json:"last_face_event" binding:"required"`
}

func (w *Webhook) getOriginalFace() {
	q := url.Values{}
	q.Add("card", string(rune(w.Persona.CardId)))
	data, err := request.SendRequest(os.Getenv("FIND_FACE_URL"), os.Getenv("FIND_FACE_API_KEY"), q)
	if err != nil {
		fmt.Println(err)
	}
	var faces Faces
	err = json.Unmarshal(data, &faces)
	if err != nil {
		fmt.Println(err)
	}
	//r := Face{thumbnail: "img/original.jpg"}
	//faces := Faces{
	//	results: []Face{r},
	//}
	w.Image.Original = faces.results[0].thumbnail
}

func (w *Webhook) Format() {
	w.Image.Thumbnail = strings.Replace(w.Image.Thumbnail, os.Getenv("PHOTO_SERVER_URL"), os.Getenv("IMG_PATH"), 1)
	w.Image.FullFrame = strings.Replace(w.Image.FullFrame, os.Getenv("PHOTO_SERVER_URL"), os.Getenv("IMG_PATH"), 1)
	w.getOriginalFace()
}
