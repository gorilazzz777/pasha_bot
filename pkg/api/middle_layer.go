package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"pasha_bot/pkg/requests"
	"strings"
)

type MlResponse struct {
	Parcels []Order `json:"parcels"`
}

type Order struct {
	TrackNumber         string `json:"trackNumber"`
	ShopCode            string `json:"shopCode"`
	ShopName            string
	DeliveryType        int64 `json:"deliveryType"`
	DeliveryTypeName    string
	DeliveryAddress     string `json:"deliveryAddress"`
	DeliveryDate        string `json:"deliveryDate"`
	TrackDate           string `json:"trackDate"`
	DeliveryCityCode    string `json:"deliveryCityCode"`
	ReceiverPhone       string `json:"recipientPhone"`
	ReceiverEmail       string `json:"recipientEmail"`
	ReceiverName        string `json:"recipientName"`
	PointStopCode       string `json:"pointStopCode"`
	DeliveryCityName    string
	DeliveryCountryName string
}

type MiddleLayer struct {
	token string
	url   string
	order Order
}

type Point struct {
	UpdateDate  string `json:"updateDate"`
	PointCode   string `json:"pointCode"`
	Name        string `json:"name"`
	CountryName string `json:"countryName"`
	CityName    string `json:"cityName"`
}

type Status struct {
	Code           int    `json:"code"`
	StatusDateTime string `json:"statusDateTime"`
}

type Parcel struct {
	TrackNumber    string   `json:"trackNumber"`
	ParcelStatuses []Status `json:"parcelStatuses"`
}

type Shop struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type middleLayerPointResponse struct {
	Points []Point `json:"points"`
}

type middleLayerShopResponse struct {
	Shops []Point `json:"shops"`
}

type middleLayerStatusResponse struct {
	Parcels []Parcel `json:"parcels"`
}

func (ml *MiddleLayer) init() {
	ml.url = os.Getenv("ML_URL")
	ml.token = fmt.Sprintf("Bearer %s", os.Getenv("ML_TOKEN"))
	return
}

func (ml *MiddleLayer) fillGeoData() error {
	var result middleLayerPointResponse
	q := url.Values{}
	q.Add("code", ml.order.PointStopCode)
	q.Add("returnedFields[0]", "Name")
	q.Add("returnedFields[1]", "CountryName")
	q.Add("returnedFields[2]", "CityName")
	response, err := requests.SendRequest(
		fmt.Sprintf("%spoints", ml.url),
		ml.token,
		q)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(response, &result); err != nil {
		return err
	}
	if len(result.Points) != 0 {
		ml.order.DeliveryAddress = result.Points[0].Name
		ml.order.DeliveryCityName = result.Points[0].CityName
		ml.order.DeliveryCountryName = result.Points[0].CountryName
	}
	return nil
}

func (ml *MiddleLayer) fillDeliveryType() {
	if ml.order.DeliveryType == 1 || ml.order.DeliveryType == 2 {
		ml.order.DeliveryTypeName = "ПВЗ"
	} else if ml.order.DeliveryType == 3 || ml.order.DeliveryType == 4 {
		ml.order.DeliveryTypeName = "Дверь"
	}
}

func (ml *MiddleLayer) fillShopName() error {
	var result middleLayerShopResponse
	q := url.Values{}
	q.Add("searchParameter.code", ml.order.ShopCode)
	q.Add("limit", "1")
	q.Add("offset", "0")
	response, err := requests.SendRequest(
		fmt.Sprintf("%s/partner/shop/find", ml.url),
		ml.token,
		q)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(response, &result); err != nil {
		return err
	}
	if len(result.Shops) != 0 {
		ml.order.ShopName = result.Shops[0].Name
	}
	return nil
}

func (ml *MiddleLayer) fillTrackDate() error {
	var result middleLayerStatusResponse
	q := url.Values{}
	q.Add("trackNumbers", ml.order.TrackNumber)
	q.Add("sort", "1")
	response, err := requests.SendRequest(
		fmt.Sprintf("%s/parcel/status-parcels", ml.url),
		ml.token,
		q)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(response, &result); err != nil {
		return err
	}
	if len(result.Parcels) != 0 && len(result.Parcels[0].ParcelStatuses) != 0 {
		ml.order.TrackDate = strings.Split(result.Parcels[0].ParcelStatuses[0].StatusDateTime, "T")[0]
	}
	return nil
}

func (ml *MiddleLayer) getParcelData() error {
	var result MlResponse
	q := url.Values{}
	q.Add("trackNumbers", ml.order.TrackNumber)
	q.Add("returnedFields[0]", "shopCode")
	q.Add("returnedFields[1]", "deliveryType")
	q.Add("returnedFields[2]", "pointStopCode")
	q.Add("returnedFields[3]", "recipientPhone")
	q.Add("returnedFields[4]", "recipientName")
	q.Add("returnedFields[5]", "trackNumber")
	q.Add("returnedFields[6]", "recipientEmail")
	response, err := requests.SendRequest(
		fmt.Sprintf("%sparcel", ml.url),
		ml.token,
		q)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(response, &result); err != nil {
		return err
	}
	if len(result.Parcels) != 0 {
		ml.order = result.Parcels[0]
	}
	return nil
}

func (ml *MiddleLayer) ParcelData(trackNum string) (Order, error) {
	ml.init()
	ml.order.TrackNumber = trackNum
	err := ml.getParcelData()
	if err != nil {
		return Order{}, err
	}
	err = ml.fillGeoData()
	if err != nil {
		return Order{}, err
	}
	ml.fillDeliveryType()
	err = ml.fillShopName()
	if err != nil {
		return Order{}, err
	}
	err = ml.fillTrackDate()
	if err != nil {
		return Order{}, err
	}
	return ml.order, nil
}
