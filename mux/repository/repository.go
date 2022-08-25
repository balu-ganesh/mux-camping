package repository

import (
	"time"

	"github.com/camping/entities"
)

type Repository interface {
	Close()
	// FindByID(id string) (*entities.Booking, error)
	// Find() ([]*entities.Booking, error)
	GetAllAvailableSlotsByDate(time.Time, time.Time) (*[]entities.AvailableSlots, error)
	GetSlotByID(uint) (*entities.AvailableSlots, error)
	BookSlot(booking *entities.Booking) (*entities.Booking, error)
	GetBookedSlotsByUserID(userID string) ([]entities.Booking, error)

	// Create(user **entities.Booking) error
	// Update(user **entities.Booking) error
	// Delete(id string) error
}
