package service

import (
	"encoding/json"
	wraperrors "github.com/pkg/errors"
	"net/url"
	"strconv"
	"webserver/model"
	"webserver/sqlite"
)

func GetHaisData(values url.Values) ([]byte, error) {
	var haisData []model.Hais
	var err error

	id := values.Get("id")
	hai := values.Get("hai")
	city := values.Get("city")

	if id != "" {
		id, _ := strconv.Atoi(id)
		haisData, err = sqlite.GetHaisByID(id)

		if err != nil {
			return nil, wraperrors.New("Error fetching hais details")
		}
	} else if city != "" && hai != "" {

		haisData, err = sqlite.GetHaiByCityAndHai(city, hai)
		if err != nil {
			return nil, wraperrors.New("Error fetching hais details")
		}
	} else if hai != "" {

		haisData, err = sqlite.GetHaiByHai(hai)
		if err != nil {
			return nil, wraperrors.New("Error fetching hais details")
		}

	} else if city != "" {
		haisData, err = sqlite.GetHaiByCity(city)
		if err != nil {
			return nil, wraperrors.New("Error fetching hais details")
		}
	} else {
		return nil, wraperrors.Errorf("client side issue")
	}

	j, _ := json.Marshal(haisData)
	return j, nil
}
