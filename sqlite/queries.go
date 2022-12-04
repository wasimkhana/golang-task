package sqlite

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	wraperrors "github.com/pkg/errors"
	"github.com/spf13/viper"
	"webserver/config"
	"webserver/model"
)

// DBconnection - to connect with sqlite3 database
func DBconnection() (*sqlx.DB, error) {
	return sqlx.Connect(viper.GetString(config.DbName), viper.GetString(config.DbPath))
}

// GetHaisByID - to retrieve record for id param
func GetHaisByID(id int) ([]model.Hais, error) {
	var hais []model.Hais

	cli, err := DBconnection()
	if err != nil {
		return nil, wraperrors.Wrap(err, "database connection issue")
	}
	defer cli.Close()

	if err := cli.Select(&hais, fmt.Sprintf(`SELECT * FROM Hais WHERE h_id=%d`, id)); err != nil {
		fmt.Println(err.Error())
		return nil, wraperrors.Wrap(err, "record not found")
	}

	return hais, nil
}

// GetHaiByCityAndHai - to retrieve records based on city and hai params
func GetHaiByCityAndHai(city, hai string) ([]model.Hais, error) {
	var hais []model.Hais

	cli, err := DBconnection()
	if err != nil {
		return nil, wraperrors.Wrap(err, "database connection issue")
	}
	defer cli.Close()

	// Wildcards
	hai = "%" + hai + "%"
	city = "%" + city + "%"
	if err := cli.Select(&hais, fmt.Sprintf(`SELECT * FROM Hais where city like '%s' and hai like '%s'`, city, hai)); err != nil {
		return nil, wraperrors.Wrap(err, "records not found")
	}

	return hais, nil
}

// GetHaiByHai - to retrieve records based on hai param
func GetHaiByHai(hai string) ([]model.Hais, error) {
	var hais []model.Hais

	cli, err := DBconnection()
	if err != nil {
		return nil, wraperrors.Wrap(err, "database connection issue")
	}
	defer cli.Close()

	hai = "%" + hai + "%"
	if err := cli.Select(&hais, fmt.Sprintf(`SELECT * FROM Hais where hai like '%s'`, hai)); err != nil {
		return nil, wraperrors.Wrap(err, "records not found")
	}

	return hais, nil
}

// GetHaiByCity - to retrieve records based on city param
func GetHaiByCity(city string) ([]model.Hais, error) {
	var hais []model.Hais

	cli, err := DBconnection()
	if err != nil {
		return nil, wraperrors.Wrap(err, "database connection issue")
	}
	defer cli.Close()

	city = "%" + city + "%"
	if err := cli.Select(&hais, fmt.Sprintf(`SELECT * FROM Hais where city like '%s'`, city)); err != nil {
		return nil, wraperrors.Wrap(err, "records not found")
	}

	return hais, nil
}

// GetHaiCounter - To retrieve list of cities and total hais in that city
func GetHaiCounter() ([]model.HaiCounter, error) {
	var haicounter []model.HaiCounter

	cli, err := DBconnection()
	if err != nil {
		return nil, wraperrors.Wrap(err, "database connection issue")
	}
	defer cli.Close()

	if err := cli.Select(&haicounter, "Select distinct city, count(hai) haicounter from hais group by city;"); err != nil {
		return nil, wraperrors.Wrap(err, "records not found")
	}

	return haicounter, nil
}
