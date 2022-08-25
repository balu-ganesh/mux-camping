package services

import (
	"time"

	"github.com/camping/entities"
)

type SlotService interface {
	GetSlotsByDate(startDate time.Time, endDate time.Time) (*SlotResponse, error)
	GetSlotsByUser() SlotResponse
}

type BookingService interface {
	CancelBooking()
	ModifyBooking()
	BookSlot(bookSlot entities.BookSlot) (*entities.Booking, error)
}

type SlotResponse struct {
	Slots []entities.AvailableSlots `json:"slots,omitempty"`
}

type ModifiedSlotResponse struct {
	Slot entities.Slot `json:"slots,omitempty"`
}
