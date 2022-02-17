package geocoder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"lesson2/distance/navigator/infra"
	"lesson2/distance/point"
	"net/http"
)

type Geocoder struct {
	client *http.Client
	url    string
	token  string
}

func (g Geocoder) Geocode(address string) (point point.Point2d, err error) {
	//TODO implement me
	panic("implement me")
}

func (g Geocoder) ReverseGeocode(point point.Point2d) (data infra.GeocodeData, err error) {
	jsonRequest, _ := json.Marshal(map[string]string{"lat": fmt.Sprintf("%f", point.X()), "lon": fmt.Sprintf("%f", point.Y()), "count": "1"})
	req, _ := http.NewRequest("POST", g.url, bytes.NewBuffer(jsonRequest))
	req.Header.Add("Authorization", "Token "+g.token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	response, err := g.client.Do(req)
	if err != nil {
		return
	}

	//buf, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(buf))

	var res map[string]interface{}
	json.NewDecoder(response.Body).Decode(&res)
	suggestions := res["suggestions"].([]interface{})
	sugElem := suggestions[0].(map[string]interface{})
	data1 := sugElem["data"].(map[string]interface{})
	postalCode := data1["postal_code"].(string)
	fmt.Println("postalcode ", postalCode)

	return
}

/*
curl -X POST \
-H "Content-Type: application/json" \
-H "Accept: application/json" \
-H "Authorization: Token 11704dfd376beff80fa6b77d454026c9ddd025f0" \
-d '{ "lat": 55.878, "lon": 37.653 }' \
https://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address
*/

func NewGeocoder() *Geocoder {
	return &Geocoder{client: &http.Client{}, url: "https://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address", token: "11704dfd376beff80fa6b77d454026c9ddd025f0"}
}
