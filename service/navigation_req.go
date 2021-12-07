package service

import (
    "encoding/json"
    "errors"
    "fmt"

    "github.com/NamSoGong/DomusPopuli-API/domain/coords"
    "github.com/NamSoGong/DomusPopuli-API/domain/navigation"
)

func NavigationSearch(src, dst coords.Coordinate_t) (*navigation.Response_t, error) {
    // Generate URL
    reqUrl := "https://apis-navi.kakaomobility.com/v1/directions"
    reqUrl += fmt.Sprintf("?origin=%s,%s&destination=%s,%s", src.Long, src.Lat, dst.Long, dst.Lat)
    reqUrl += "&waypoints=&priority=RECOMMEND&car_fuel=&car_hipass=&alternatives=&road_details="

    reader, err := KakaoMAPIGet(reqUrl)
    if err != nil {
        return nil, err
    }

    // Parse JSON
    dec := json.NewDecoder(reader)
    jsonBody := navigation.Response_t{}

    dec.Decode(&jsonBody)

    if len(jsonBody.Routes) < 1 {
        return nil, errors.New("Route not found")
    }

    return &jsonBody, nil
}