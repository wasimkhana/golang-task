package service

import (
	"encoding/json"
	"webserver/sqlite"
)

func GetHaisCounterService() ([]byte, error) {
	//var haisData []model.Hais
	//var err error
	//haisData, err = sqlite.GetHaiCounter()
	//if err != nil {
	//	return nil, wraperrors.New("Error fetching profile details")
	//}
	haicounter, _ := sqlite.GetHaiCounter()
	//fmt.Println(err)
	j, _ := json.Marshal(haicounter)
	return j, nil
}
