package repository

import (
	"database/sql"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/camping/database"
	"github.com/camping/entities"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func Test_GetAllAvailableSlotsByDate(t *testing.T) {
	db, _ := database.GetDB()
	repo := BookingRepostory{
		db: db,
	}

	testCases := map[string]struct {
		params map[string]string
	}{
		"goodParams": {
			map[string]string{
				"startDate": "2022-08-25 12:00",
				"endDate":   "2022-08-25 12:00",
			},
		},
	}

	for _, tp := range testCases {
		layout := "2006-01-02 15:04"

		startDate, err := time.Parse(layout, tp.params["startDate"])
		if err != nil {
			t.Fatalf("error occured while parsing input params %+v ", err)
		}

		endDate, err := time.Parse(layout, tp.params["endDate"])
		if err != nil {
			t.Fatalf("error occured while parsing input params %+v", err)
		}

		slots, err := repo.GetAllAvailableSlotsByDate(startDate, endDate)
		if err != nil {
			t.Fatalf("error occured while accessing available slots %+v", err)
		}
		if slots == nil {
			t.Fatalf("no available slots %+v", err)
		}
	}

}

func Test_BookSlot(t *testing.T) {
	db, _ := database.GetDB()
	repo := BookingRepostory{
		db: db,
	}

	testCases := []struct {
		params         entities.Booking
		slotID         uint
		expectedStatus int
		startDate      string
		endDate        string
	}{
		{
			params: entities.Booking{
				SlotID: 1002,
				User: entities.User{
					Email:     "test@test1.com",
					FirstName: "Dave",
					LastName:  "Sheing",
				},
			},
			expectedStatus: http.StatusOK,
			startDate:      "2022-08-25 12:00",
			endDate:        "2022-08-26 12:00",
		},
	}

	for _, tp := range testCases {
		layout := "2006-01-02 15:04"

		startDate, err := time.Parse(layout, tp.startDate)
		if err != nil {
			t.Fatalf("error occured while parsing input params %+v", err)
		}

		endDate, err := time.Parse(layout, tp.endDate)
		if err != nil {
			t.Fatalf("error occured while parsing input params %+v", err)
		}
		tp.params.StartDate = startDate
		tp.params.EndDate = endDate

		slots, err := repo.BookSlot(&tp.params)
		if err != nil {
			t.Fatalf("error occured while accessing available slots %+v", err)
		}
		if slots == nil {
			t.Fatalf("no available slots %+v", err)
		}
	}

}

func Test_BookSlotMultiUser(t *testing.T) {
	db, _ := database.GetDB()
	repo := BookingRepostory{
		db: db,
	}

	testCases := []struct {
		params         entities.Booking
		slotID         uint
		expectedStatus int
		startDate      string
		endDate        string
	}{
		{
			params: entities.Booking{
				SlotID: 1002,
				User: entities.User{
					Email:     "test@test1.com",
					FirstName: "Dave",
					LastName:  "Sheing",
				},
			},
			expectedStatus: http.StatusOK,
			startDate:      "2022-08-25 12:00",
			endDate:        "2022-08-26 12:00",
		},
	}

	for _, tp := range testCases {
		layout := "2006-01-02 15:04"

		startDate, err := time.Parse(layout, tp.startDate)
		if err != nil {
			t.Fatalf("error occured while parsing input params %+v", err)
		}

		endDate, err := time.Parse(layout, tp.endDate)
		if err != nil {
			t.Fatalf("error occured while parsing input params %+v", err)
		}
		tp.params.StartDate = startDate
		tp.params.EndDate = endDate

		done := make(chan bool)

		for i := 0; i < 2; i++ {
			go func() {
				slots, err := repo.BookSlot(&tp.params)
				if err != nil {
					log.Fatalf("error occured while accessing available slots %+v", err)
				}
				if slots == nil {
					log.Fatalf("no available slots %+v", err)
				}
				done <- true
			}()
			log.Println("waiting to finish two simultaneous booking")
		}
		for i := 0; i < 2; i++ {
			log.Println(<-done)
		}
	}

}
