package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/camping/entities"
	"github.com/camping/services"
	"github.com/gorilla/mux"
)

func (h HTTPServer) GetAllAvailableSlotsByDate(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()
	var startDate, endDate time.Time
	layout := "2006-01-02 15:04"
	if s, ok := params["start"]; ok {
		date, err := time.Parse(layout, s[0])
		if err != nil {
			log.Println("error occured while parsing input params ", err)
		}
		startDate = date
	} else {
		fmt.Print("no start date params in request")
	}
	if s, ok := params["end"]; ok {
		date, err := time.Parse(layout, s[0])
		if err != nil {
			log.Println("error occured while parsing input params")
		}
		endDate = date
		fmt.Println(date, err)
	} else {
		fmt.Print("no end date params in request")
	}
	bookingService := services.GetInstance()

	b, err := bookingService.GetAllAvailableSlotsByDate(startDate, endDate)
	if err != nil {
		log.Print("error occured while booking slot ", err)
	}
	err = json.NewEncoder(w).Encode(b)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

}

// TODO : json unmarshaller needs to be added for date fields.
func (h HTTPServer) ReserveSlot(res http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil || len(reqBody) == 0 {
		writeHTTPError(res, &BadRequestError{err: ErrEmptyRequestBody})
		return
	}
	var bookslot entities.BookSlot

	json.Unmarshal(reqBody, &bookslot)
	vars := mux.Vars(r)
	var slotID uint
	if slot, ok := vars["slotid"]; !ok {
		writeHTTPError(res, &BadRequestError{err: ErrNoSlotID})
		return
	} else {
		u64, err := strconv.ParseUint(slot, 10, 32)
		if err != nil {
			fmt.Println(err)
		}
		slotID = uint(u64)
	}

	bookingService := services.GetInstance()

	b, err := bookingService.BookSlot(bookslot, uint(slotID))
	if err != nil {
		log.Print("error occured while booking slot ", err)
		writeHTTPError(res, err)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	writeHTTPResponse(res, b)
}

func (h HTTPServer) GetSlotsByUser(w http.ResponseWriter, r *http.Request) {
	var bookslot entities.BookSlot
	json.NewDecoder(r.Body).Decode(&bookslot)

	vars := mux.Vars(r)
	if userID, ok := vars["id"]; ok {
		bookingService := services.GetInstance()
		b, err := bookingService.GetBookedSlotsByUserID(userID)
		if err != nil {
			log.Print("error occured while booking slot ", err)
		}
		json.NewEncoder(w).Encode(b)
	}
}

func (h HTTPServer) ManageSlot(res http.ResponseWriter, r *http.Request) {

	var bookslot entities.BookSlot
	json.NewDecoder(r.Body).Decode(&bookslot)

	vars := mux.Vars(r)
	var slotID uint
	if slot, ok := vars["slotid"]; !ok {
		writeHTTPError(res, &BadRequestError{err: ErrNoSlotID})
		return
	} else {
		u64, err := strconv.ParseUint(slot, 10, 32)
		if err != nil {
			fmt.Println(err)
		}
		slotID = uint(u64)
	}

	bookingService := services.GetInstance()
	b, err := bookingService.BookSlot(bookslot, slotID)
	if err != nil {
		log.Print("error occured while booking slot ", err)
	}
	res.Header().Set("Content-Type", "application/json")
	writeHTTPResponse(res, b)
}

func writeHTTPResponse(res http.ResponseWriter, val interface{}) {
	jsonData, err := json.Marshal(val)
	if err != nil {
		writeHTTPError(res, err)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonData)
}

func (h HTTPServer) CancelBooking(res http.ResponseWriter, r *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var bookslot entities.BookSlot
	json.NewDecoder(r.Body).Decode(&bookslot)
	vars := mux.Vars(r)
	var slotID uint
	if slot, ok := vars["slotid"]; !ok {
		writeHTTPError(res, &BadRequestError{err: ErrNoSlotID})
		return
	} else {
		u64, err := strconv.ParseUint(slot, 10, 32)
		if err != nil {
			fmt.Println(err)
		}
		slotID = uint(u64)
	}
	bookingService := services.GetInstance()
	b, err := bookingService.BookSlot(bookslot, slotID)
	if err != nil {
		log.Print("error occured while booking slot ", err)
	}

	json.NewEncoder(res).Encode(b)
}
