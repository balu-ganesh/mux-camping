package server

import (
	"encoding/json"
	"errors"
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

func (h HTTPServer) GetAllAvailableSlotsByDate(res http.ResponseWriter, req *http.Request) {

	params := req.URL.Query()
	var startDate, endDate time.Time
	layout := "2006-01-02 15:04"

	if s, ok := params["start"]; ok {
		date, err := time.Parse(layout, s[0])
		if err != nil {
			log.Println("error occured while parsing input params ", err)
			writeHTTPError(res, &BadRequestError{err: ErrNoSlotID})
			return
		}
		startDate = date
	} else {
		writeHTTPError(res, &BadRequestError{err: errors.New("no start date params in request")})
		return
	}
	if s, ok := params["end"]; ok {
		date, err := time.Parse(layout, s[0])
		if err != nil {
			log.Println("error occured while parsing input params")
			writeHTTPError(res, &BadRequestError{err: errors.New("error occured while parsing end date")})
			return
		}
		endDate = date
	} else {
		writeHTTPError(res, &BadRequestError{err: errors.New("no end date params in request")})
		return
	}

	bookingService := services.GetInstance()
	b, err := bookingService.GetAllAvailableSlotsByDate(startDate, endDate)
	if err != nil {
		log.Print("error occured while booking slot ", err)
		writeHTTPError(res, errors.New("error occured while processing"+err.Error()))
	}
	writeHTTPResponse(res, b)
}

// TODO : json unmarshaller needs to be added for date fields.
func (h HTTPServer) ReserveSlot(res http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil || len(reqBody) == 0 {
		writeHTTPError(res, &BadRequestError{err: ErrEmptyRequestBody})
		return
	}
	var bookslot entities.BookSlot
	layout := "2006-01-02 15:04"
	json.Unmarshal(reqBody, &bookslot)
	vars := mux.Vars(r)
	var slotID, userID uint
	if slot, ok := vars["slotid"]; !ok {
		writeHTTPError(res, &BadRequestError{err: ErrNoSlotID})
		return
	} else {
		u64, err := strconv.ParseUint(slot, 10, 32)
		if err != nil {
			writeHTTPError(res, errors.New("slot id should be a number"))
		}
		slotID = uint(u64)
	}
	if userid, ok := vars["userid"]; !ok {
		writeHTTPError(res, &BadRequestError{err: ErrNoUserID})
		return
	} else {
		u64, err := strconv.ParseUint(userid, 10, 32)
		if err != nil {
			writeHTTPError(res, errors.New("user id should be a number"))
		}
		userID = uint(u64)
	}

	if bookslot.CheckInDate == "" && bookslot.CheckoutDate == "" {
		bookslot.CheckInDate = time.Now().Format(layout)
		bookslot.CheckoutDate = time.Now().Add(time.Hour * 24 * 30).Format(layout)
	}

	bookingService := services.GetInstance()

	b, err := bookingService.BookSlot(bookslot, uint(slotID), uint(userID))
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
	b, err := bookingService.BookSlot(bookslot, slotID, 123)
	if err != nil {
		log.Print("error occured while booking slot ", err)
	}
	res.Header().Set("Content-Type", "application/json")
	writeHTTPResponse(res, b)
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
	b, err := bookingService.BookSlot(bookslot, slotID, 123)
	if err != nil {
		log.Print("error occured while booking slot ", err)
	}

	json.NewEncoder(res).Encode(b)
}
