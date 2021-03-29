package dbcontrollers

import (
	"time"

	"github.com/artofimagination/timescaledb-statistics-go-service/models"
	timescaledb "github.com/artofimagination/timescaledb-statistics-go-service/timescaledb"
)

func (c *TimescaleController) AddDailyDataFP(data []models.StatsFP) error {
	if err := c.DBFunctions.AddRowFP(timescaledb.FPTableDaily, data); err != nil {
		return ErrFailedToAddData
	}
	return nil
}

func (c *TimescaleController) DeleteDailyFPByCategory(category int) error {
	if err := c.DBFunctions.DeleteFPByCategory(timescaledb.FPTableDaily, category); err != nil {
		return ErrFailedToAddData
	}
	return nil
}

func (c *TimescaleController) DeleteDailyFPByTime(intervalString string) error {
	if err := c.DBFunctions.DeleteFPByTime(timescaledb.FPTableDaily, intervalString); err != nil {
		return ErrFailedToAddData
	}
	return nil
}

func (c *TimescaleController) GetDailyFPByCategoryAndTime(category int, time time.Time) ([]models.StatsFP, error) {
	data, err := c.DBFunctions.GetFPByCategoryAndTime(timescaledb.FPTableDaily, category, time)
	if err != nil {
		return nil, ErrFailedToAddData
	}
	return data, nil
}

func (c *TimescaleController) AddOverallDataFP(data []models.StatsFP) error {
	if err := c.DBFunctions.AddRowFP(timescaledb.FPTableOverall, data); err != nil {
		return ErrFailedToAddData
	}
	return nil
}

func (c *TimescaleController) DeleteOverallFPByCategory(category int) error {
	if err := c.DBFunctions.DeleteFPByCategory(timescaledb.FPTableOverall, category); err != nil {
		return ErrFailedToAddData
	}
	return nil
}

func (c *TimescaleController) DeleteOverallFPByTime(intervalString string) error {
	if err := c.DBFunctions.DeleteFPByTime(timescaledb.FPTableOverall, intervalString); err != nil {
		return ErrFailedToAddData
	}
	return nil
}

func (c *TimescaleController) GetOverallFPByCategoryAndTime(category int, time time.Time) ([]models.StatsFP, error) {
	data, err := c.DBFunctions.GetFPByCategoryAndTime(timescaledb.FPTableOverall, category, time)
	if err != nil {
		return nil, ErrFailedToAddData
	}
	return data, nil
}
