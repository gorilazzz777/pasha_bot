package pasha_bot

import (
	"encoding/json"
	"fmt"
	"net/url"
	"pasha_bot/pkg/request"
	"strings"
)

const ffUrl = "https://127.0.0.1:5467/objects/faces/"
const ffApiKey = "1234"

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

type Input struct {
	Date    string        `json:"created_date" binding:"required"`
	Persona MatchedCard   `json:"matched_card" binding:"required"`
	Image   LastFaceEvent `json:"last_face_event" binding:"required"`
}

func (i *Input) getOriginalFace() {
	q := url.Values{}
	q.Add("card", string(rune(i.Persona.CardId)))
	data, err := request.SendRequest(ffUrl, ffApiKey, q)
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
	i.Image.Original = faces.results[0].thumbnail
}

func (i *Input) Format() {
	i.Image.Thumbnail = strings.Replace(i.Image.Thumbnail, "http://192.168.1.58/uploads", "/img", 1)
	i.Image.FullFrame = strings.Replace(i.Image.FullFrame, "http://192.168.1.58/uploads", "/img", 1)
	i.getOriginalFace()
}
