package services

import (
	"fmt"
	"log"
	"time"

	"github.com/camping/entities"
	"github.com/camping/repository"
)

var bookingService *Booking

func init() {
	bookingService = &Booking{}
}

func GetInstance() *Booking {
	return bookingService
}

type Booking struct {
	BookingDAO repository.BookingRepostory
}

func (b *Booking) GetAllAvailableSlotsByDate(startDate time.Time, endDate time.Time) (*SlotResponse, error) {
	slots, err := b.BookingDAO.GetAllAvailableSlotsByDate(startDate, endDate)
	if err != nil {
		return nil, err
	}
	availableSlots := &SlotResponse{
		Slots: *slots,
	}
	return availableSlots, nil
}

func (b *Booking) CancelBooking() {
	//TODO
}

func (b *Booking) ModifyBooking() {
	//Merge the updates
}

func (b *Booking) BookSlot(bookSlot entities.BookSlot, slotID uint) (*entities.Booking, error) {

	layout := "2006-01-02 15:04"

	startDate, err := time.Parse(layout, bookSlot.CheckInDate)
	if err != nil {
		log.Println("error occured while parsing input params ", err)
		return nil, fmt.Errorf("error occured while parsing startDate params %+v", err)
	}

	endDate, err := time.Parse(layout, bookSlot.CheckoutDate)
	if err != nil {
		log.Println("error occured while parsing input params")
		return nil, fmt.Errorf("error occured while parsing startDate params %+v", err)
	}

	return b.BookingDAO.BookSlot(&entities.Booking{
		SlotID:      slotID,
		StartDate:   startDate,
		EndDate:     endDate,
		Description: "booking",
		User: entities.User{
			FirstName: bookSlot.FirstName,
			LastName:  bookSlot.LastName,
			Email:     bookSlot.Email,
		},
	})
}

func (b *Booking) GetBookedSlotsByUserID(userID string) ([]entities.Booking, error) {
	return b.BookingDAO.GetBookedSlotsByUserID(userID)
}
