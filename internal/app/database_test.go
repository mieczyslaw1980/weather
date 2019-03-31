package app

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type fakeDatabase struct {
	err        error
	errSave    error
	locations  []Location
	weather    Weather
	statistics Statistics
}

func (f fakeDatabase) getDBLocation(id int) (Location, error) {
	if len(f.locations) > 0 {
		return f.locations[0], f.err
	}
	return Location{}, f.err
}

func (f fakeDatabase) getDBLocations() ([]Location, error) {
	return f.locations, f.err
}

func (f fakeDatabase) saveDBLocation(location Location) error {
	return f.errSave
}

func (f fakeDatabase) deleteDBLocation(id int) error {
	return f.err
}

func (f fakeDatabase) saveDBWeather(s Weather) error {
	return f.errSave
}

func (f fakeDatabase) getStatistics(id int) (Statistics, error) {
	return f.statistics, f.err
}

func TestNewDB(t *testing.T) {

	t.Run("Invalid database configuration", func(t *testing.T) {
		user := os.Getenv("DB_USER")
		database := os.Getenv("DB_DATABASE")
		password := os.Getenv("DB_PASSWORD")
		address := os.Getenv("DB_ADDRESS")
		defer func() {
			os.Setenv("DB_USER", user)
			os.Setenv("DB_DATABASE", database)
			os.Setenv("DB_PASSWORD", password)
			os.Setenv("DB_ADDRESS", address)
		}()

		os.Setenv("DB_ADDRESS", "")
		db, err := NewDB()
		assert.NotNil(t, err)
		assert.Nil(t, db)
	})

	t.Run("Valid database configuration", func(t *testing.T) {
		user := os.Getenv("DB_USER")
		database := os.Getenv("DB_DATABASE")
		password := os.Getenv("DB_PASSWORD")
		address := os.Getenv("DB_ADDRESS")
		defer func() {
			os.Setenv("DB_USER", user)
			os.Setenv("DB_DATABASE", database)
			os.Setenv("DB_PASSWORD", password)
			os.Setenv("DB_ADDRESS", address)
		}()

		os.Setenv("DB_USER", "user")
		os.Setenv("DB_DATABASE", "db")
		os.Setenv("DB_PASSWORD", "")
		os.Setenv("DB_ADDRESS", "localhost:5432")
		db, err := NewDB()
		assert.Nil(t, err)
		assert.NotNil(t, db)
	})
}
