package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/camping/entities"
	"github.com/camping/repository"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewRepository(dialect, dsn string, idleConn, maxConn int) (repository.Repository, error) {
	db, err := sql.Open(dialect, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(idleConn)
	db.SetMaxOpenConns(maxConn)

	return &MySQLRepository{db}, nil
}

func (m *MySQLRepository) Close() {
	m.db.Close()
}

func formatDate(inputDate time.Time) string {
	return inputDate.Format("2006-01-02 15:04")
}

func (m *MySQLRepository) GetAllAvailableSlotsByDate(startDate time.Time, endDate time.Time) (*[]entities.AvailableSlots, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var availableSlots []entities.AvailableSlots
	formattedStartDate := formatDate(startDate)
	formattedEndDate := formatDate(endDate)
	sqlStr := `select s.id slot_id, s.site_id site_id from slots s 
					left join bookings b on s.id=b.slot_id and (
					( b.start_date <= ? and b.end_date between ? and ? ) and
					( b.start_date between ? and ? and b.end_date >= ? ) and
					( b.start_date <= ? and b.end_date >= ? )  )
				where b.id isnull `

	err := m.db.QueryRowContext(ctx, sqlStr, formattedEndDate, formattedStartDate, formattedEndDate, formattedStartDate, formattedEndDate, formattedEndDate, formattedEndDate, formattedEndDate).Scan(&availableSlots)
	if err != nil {
		return nil, err
	}
	return &availableSlots, nil
}

func (b *MySQLRepository) GetSlotByID(id uint) (*entities.AvailableSlots, error) {
}

func (b *MySQLRepository) BookSlot(booking *entities.Booking) (*entities.Booking, error) {
}

func (b *MySQLRepository) GetBookedSlotsByUserID(userID string) ([]entities.Booking, error) {
}
