package dbcontrollers

import (
	"time"

	"github.com/artofimagination/timescaledb-statistics-go-service/models"
	timescaledb "github.com/artofimagination/timescaledb-statistics-go-service/timescaledb"
)

func (c *TimescaleController) AddDailyDataJSON(data []models.StatsJSON) error {
	if err := c.DBFunctions.AddRowJSON(timescaledb.JSONTableDaily, data); err != nil {
		return ErrFailedToAddData
	}
	return nil
}

func (c *TimescaleController) DeleteDailyJSONByCategory(category int) error {
	if err := c.DBFunctions.DeleteJSONByCategory(timescaledb.JSONTableDaily, category); err != nil {
		return ErrFailedToAddData
	}
	return nil
}

func (c *TimescaleController) DeleteDailyJSONByTime(intervalString string) error {
	if err := c.DBFunctions.DeleteJSONByTime(timescaledb.JSONTableDaily, intervalString); err != nil {
		return ErrFailedToAddData
	}
	return nil
}

func (c *TimescaleController) GetDailyJSONByCategoryAndTime(category int, time time.Time) ([]models.StatsJSON, error) {
	data, err := c.DBFunctions.GetJSONByCategoryAndTime(timescaledb.JSONTableDaily, category, time)
	if err != nil {
		return nil, ErrFailedToAddData
	}
	return data, nil
}

func (c *TimescaleController) AddOverallDataJSON(data []models.StatsJSON) error {
	if err := c.DBFunctions.AddRowJSON(timescaledb.JSONTableOverall, data); err != nil {
		return ErrFailedToAddData
	}
	return nil
}

func (c *TimescaleController) DeleteOverallJSONByCategory(category int) error {
	if err := c.DBFunctions.DeleteFPByCategory(timescaledb.JSONTableOverall, category); err != nil {
		return ErrFailedToAddData
	}
	return nil
}

func (c *TimescaleController) DeleteOverallJSONByTime(intervalString string) error {
	if err := c.DBFunctions.DeleteJSONByTime(timescaledb.JSONTableOverall, intervalString); err != nil {
		return ErrFailedToAddData
	}
	return nil
}

func (c *TimescaleController) GetOverallJSONByCategoryAndTime(category int, time time.Time) ([]models.StatsJSON, error) {
	data, err := c.DBFunctions.GetJSONByCategoryAndTime(timescaledb.JSONTableOverall, category, time)
	if err != nil {
		return nil, ErrFailedToAddData
	}
	return data, nil
}
