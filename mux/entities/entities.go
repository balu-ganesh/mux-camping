package entities

import (
	"time"

	"gorm.io/gorm"
)

type Slot struct {
	gorm.Model
	Active bool `json:"active"`
	SiteID int  `json:"slotID"`
	Site   Site `gorm:"foreignKey:SiteID"`
}

type Camp struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Active      bool   `json:"active"`
}

type Site struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
	CampID      uint   `json:"campID"`
	Camp        Camp   `gorm:"foreignKey:CampID"`
}

type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type Booking struct {
	gorm.Model
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	Description string    `json:"description"`
	SlotID      uint      `json:"slotID"`
	Slot        Slot      `gorm:"foreignKey:ID;references:SlotID"`
	User        User      `gorm:"foreignKey:ID"`
}

type BookSlot struct {
	gorm.Model
	Email        string `json:"email"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	CheckInDate  string `json:"checkInDate"`
	CheckoutDate string `json:"checkoutDate"`
}

type AvailableSlots struct {
	SlotID int `gorm:"column:slot_id"`
	SiteID int `gorm:"column:site_id"`
	// Slots []SlotDetails `json:"availableSlots"`
}

type SlotDetails struct {
	SlotID uint `json:"slotID"`
	SiteID uint `json:"siteID"`
}
