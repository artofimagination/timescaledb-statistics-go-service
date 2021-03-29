package restcontrollers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/artofimagination/timescaledb-statistics-go-service/dbcontrollers"
	"github.com/artofimagination/timescaledb-statistics-go-service/models"
	"github.com/pkg/errors"
)

var timeLayout = "Mon Jan 02 2006 15:04:05.0000 GMT-0700"

func buildJSONStat(input map[string]interface{}) ([]models.StatsJSON, error) {
	inputDataList, ok := input["data_to_store"].([]interface{})
	if !ok {
		return nil, errors.New("Missing 'data_to_store'")
	}

	dataList := make([]models.StatsJSON, len(inputDataList))
	for i, data := range inputDataList {
		createAtString, ok := data.(map[string]interface{})["created_at"].(string)
		if !ok {
			return nil, errors.New("Missing 'created_at'")
		}
		createdAt, err := time.Parse(timeLayout, createAtString)
		if err != nil {
			return nil, err
		}

		category, ok := data.(map[string]interface{})["category"].(int)
		if !ok {
			return nil, errors.New("Missing 'category'")
		}

		content, ok := data.(map[string]interface{})["data"]
		if !ok {
			return nil, errors.New("Missing 'data'")
		}
		d := models.StatsJSON{
			CreatedAt: createdAt,
			Category:  category,
			Data:      content.(map[string]interface{}),
		}
		dataList[i] = d
	}
	return dataList, nil
}

func (c *RESTController) addDailyJSON(w ResponseWriter, r *Request) {
	log.Println("Adding JSON daily stats data")
	input, err := decodePostData(w, r)
	if err != nil {
		return
	}

	dataList, err := buildJSONStat(input)
	if err != nil {
		w.writeError(err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.DBController.AddDailyDataJSON(dataList); err != nil {
		if err.Error() == dbcontrollers.ErrFailedToAddData.Error() {
			w.writeError(err.Error(), http.StatusAccepted)
			return
		}
		w.writeError(err.Error(), http.StatusInternalServerError)
		return
	}

	w.writeData("OK", http.StatusCreated)
}

func (c *RESTController) addOverallJSON(w ResponseWriter, r *Request) {
	log.Println("Adding JSON overall stats data")
	input, err := decodePostData(w, r)
	if err != nil {
		return
	}

	dataList, err := buildJSONStat(input)
	if err != nil {
		w.writeError(err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.DBController.AddOverallDataJSON(dataList); err != nil {
		if err.Error() == dbcontrollers.ErrFailedToAddData.Error() {
			w.writeError(err.Error(), http.StatusAccepted)
			return
		}
		w.writeError(err.Error(), http.StatusInternalServerError)
		return
	}

	w.writeData("OK", http.StatusCreated)
}

func (c *RESTController) getDailyJSONByCategory(w ResponseWriter, r *Request) {
	log.Println("Getting daily JSON stats by category")
	if err := checkRequestType(GET, w, r); err != nil {
		return
	}

	categories, ok := r.URL.Query()["category"]
	if !ok || len(categories[0]) < 1 {
		w.writeError("Url Param 'category' is missing", http.StatusBadRequest)
		return
	}

	times, ok := r.URL.Query()["interval"]
	if !ok || len(times[0]) < 1 {
		w.writeError("Url Param 'interval' is missing", http.StatusBadRequest)
		return
	}
	interval, err := time.Parse(timeLayout, times[0])
	if err != nil {
		w.writeError(err.Error(), http.StatusBadRequest)
		return
	}

	category, err := strconv.Atoi(categories[0])
	if err != nil {
		w.writeError(err.Error(), http.StatusBadRequest)
		return
	}

	stats, err := c.DBController.GetDailyJSONByCategoryAndTime(category, interval)
	if err != nil {
		if err.Error() == dbcontrollers.ErrFailedToAddData.Error() {
			w.writeError(err.Error(), http.StatusAccepted)
			return
		}
		w.writeError(err.Error(), http.StatusInternalServerError)
		return
	}

	w.writeData(stats, http.StatusOK)
}

func (c *RESTController) getOverallJSONByCategory(w ResponseWriter, r *Request) {
	log.Println("Getting overall JSON stats by category")
	if err := checkRequestType(GET, w, r); err != nil {
		return
	}

	categories, ok := r.URL.Query()["category"]
	if !ok || len(categories[0]) < 1 {
		w.writeError("Url Param 'category' is missing", http.StatusBadRequest)
		return
	}

	times, ok := r.URL.Query()["interval"]
	if !ok || len(times[0]) < 1 {
		w.writeError("Url Param 'interval' is missing", http.StatusBadRequest)
		return
	}
	interval, err := time.Parse(timeLayout, times[0])
	if err != nil {
		w.writeError(err.Error(), http.StatusBadRequest)
		return
	}

	category, err := strconv.Atoi(categories[0])
	if err != nil {
		w.writeError(err.Error(), http.StatusBadRequest)
		return
	}

	stats, err := c.DBController.GetOverallJSONByCategoryAndTime(category, interval)
	if err != nil {
		if err.Error() == dbcontrollers.ErrFailedToAddData.Error() {
			w.writeError(err.Error(), http.StatusAccepted)
			return
		}
		w.writeError(err.Error(), http.StatusInternalServerError)
		return
	}

	w.writeData(stats, http.StatusOK)
}
