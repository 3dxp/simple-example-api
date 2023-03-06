package weather

import (
	"math/rand"
	"testing"
	"time"
)

func TestGetPredictions(t *testing.T) {
	t.Run("GetsExpectedNumberOfPredictions", func(st *testing.T) {
		results := GetPredictions()

		if len(results) != 5 {
			st.Errorf("ExpecteTestGetPredictionByLocationID/GetsExpectedPrediction in github.com/3dxp/lame-example-api/weatherd 5 predictions; Got %d", len(results))
		}
	})
	t.Run("BROKEN TEST FOR EXAMPLE", func(t *testing.T) {
		t.Fail()
	})
}

func TestGetPredictionByLocationID(t *testing.T) {
	t.Run("GetsExpectedPrediction", func(t *testing.T) {
		locationID := "ASDF"
		now := time.Now()
		rand.Seed(now.UnixMilli())
		result := GetPredictionByLocationID(locationID)
		t.Run("WithExpectedLocation", func(t *testing.T) {
			if result.LocationID != locationID {
				t.Errorf(
					"Predition did not have expected locationID. Found '%s', not '%s'.",
					result.LocationID,
					locationID,
				)
			}
		})
		t.Run("WithHighTemperatureGreaterThanLowTemperature", func(t *testing.T) {
			if result.TempHigh < result.TempLow {
				t.Error("Expected TempHigh to be greater than TempLow")
			}
		})
		t.Run("WithExpectedDateTime", func(t *testing.T) {
			if result.Datetime.Before(now) {
				t.Errorf("Expected prediction to be generated after '%s'", now)
			}
		})
	})
}
