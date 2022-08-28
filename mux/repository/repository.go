package repository

import (
	"time"

	"github.com/camping/entities"
)

type Repository interface {
	Close()
	GetAllAvailableSlotsByDate(time.Time, time.Time) (*[]entities.AvailableSlots, error)
	GetSlotByID(uint) (*entities.AvailableSlots, error)
	BookSlot(booking *entities.Booking) (*entities.Booking, error)
	GetBookedSlotsByUserID(userID string) ([]entities.Booking, error)
}
