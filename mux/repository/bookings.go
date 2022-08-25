package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/camping/database"
	"github.com/camping/entities"
	"gorm.io/gorm"
)

type BookingRepostory struct {
	db *gorm.DB
}

func (b *BookingRepostory) GetAllAvailableSlotsByDate(startDate time.Time, endDate time.Time) (*[]entities.AvailableSlots, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	sqlStr := `select s.id slot_id, s.site_id site_id from slots s 
					left join bookings b on s.id=b.slot_id and (
					( b.start_date <= strftime('%Y-%m-%d %H:%M',?) and b.end_date between strftime('%Y-%m-%d %H:%M',?) and strftime('%Y-%m-%d %H:%M',?) ) and
					( b.start_date between strftime('%Y-%m-%d %H:%M',?) and strftime('%Y-%m-%d %H:%M',?) and b.end_date >= strftime('%Y-%m-%d %H:%M',?) ) and
					( b.start_date <= strftime('%Y-%m-%d %H:%M',?) and b.end_date >= strftime('%Y-%m-%d %H:%M',?) )  )
				where b.id isnull `

	var availableSlots []entities.AvailableSlots
	formattedStartDate := formatDate(startDate)
	formattedEndDate := formatDate(endDate)

	db.Raw(sqlStr, formattedEndDate, formattedStartDate, formattedEndDate, formattedStartDate, formattedEndDate, formattedEndDate, formattedEndDate, formattedEndDate).Scan(&availableSlots)
	return &availableSlots, err
}

func formatDate(inputDate time.Time) string {
	return inputDate.Format("2006-01-02 15:04")
}

func (b *BookingRepostory) GetSlotByID(id uint) (*entities.AvailableSlots, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}
	var currentSlot entities.AvailableSlots
	db.First(&currentSlot)
	return &currentSlot, err
}

func (b *BookingRepostory) BookSlot(booking *entities.Booking) (*entities.Booking, error) {
	fmt.Println("booking started ")
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}
	db.Transaction(func(tx *gorm.DB) error {
		tx.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()
		if err := tx.Error; err != nil {
			return err
		}

		if availableSlots, errorAlreadyBooked := b.GetAllAvailableSlotsByDate(booking.StartDate, booking.EndDate); errorAlreadyBooked == nil {
			for _, s := range *availableSlots {
				if s.SlotID == int(booking.SlotID) {
					log.Printf("selected slot %v available. booking..", booking.SlotID)
					if err := tx.Create(&booking).Error; err != nil {
						return err
					}
					fmt.Println("created the record")
					continue
				}
			}
		}

		return tx.Commit().Error
	})

	return booking, nil
}

func (b *BookingRepostory) GetBookedSlotsByUserID(userID string) ([]entities.Booking, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}
	bookings := make([]entities.Booking, 0)
	db.Where("user_id = ?", userID).Preload("Slot").Find(&bookings)
	return bookings, nil
}
