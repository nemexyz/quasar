package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"quasar/domain"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SatelliteRepository struct {
	db *gorm.DB
}

func NewSatelliteRepository() *SatelliteRepository {
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_DATABASE")
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db, err2 := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err2 != nil {
		panic(err2)
	}

	return &SatelliteRepository{
		db: db,
	}
}

func (r *SatelliteRepository) SaveMessage(name string, distance float64, message []string) error {
	var s Satellite
	result := r.db.Where("name = ?", name).First(&s)
	if result.Error != nil {
		return result.Error
	}
	m := Message{Distance: distance, Date: time.Now(), Satellite: s}
	result2 := r.db.Create(&m)
	if result2.Error != nil {
		return result2.Error
	}

	for i, word := range message {
		w := Word{Word: word, Position: i + 1, Message: m}
		result3 := r.db.Create(&w)
		if result3.Error != nil {
			return result3.Error
		}
	}

	return nil
}

func (r *SatelliteRepository) GetSatelliteByName(name string) (*domain.Satellite, error) {
	var s Satellite
	result := r.db.Where("name = ?", name).First(&s)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Satellite %s does not exist", name)
		}
		return nil, result.Error
	}

	return domain.NewSatellite(s.Name, [2]float64{float64(s.X), float64(s.Y)}), nil
}

func (r *SatelliteRepository) GetLastMessageFrom(name string) (float64, []string, error) {
	var s Satellite
	result := r.db.Where("name = ?", name).First(&s)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, nil, fmt.Errorf("Satellite %s does not exist", name)
		}
		return 0, nil, result.Error
	}

	var m Message
	result = r.db.Where("satellite_id = ?", s.ID).Order("date desc").First(&m)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, nil, nil
		}
		return 0, nil, result.Error
	}

	var words []Word
	result = r.db.Where("message_id = ?", m.ID).Order("position asc").Find(&words)
	if result.Error != nil {
		return 0, nil, result.Error
	}

	var message []string
	for _, w := range words {
		message = append(message, w.Word)
	}

	return m.Distance, message, nil
}

func (r *SatelliteRepository) GetAllSatellites() ([]*domain.Satellite, error) {
	var ss []Satellite
	result := r.db.Find(&ss)
	if result.Error != nil {
		return nil, result.Error
	}

	var ds []*domain.Satellite
	for _, s := range ss {
		ds = append(ds, domain.NewSatellite(s.Name, [2]float64{float64(s.X), float64(s.Y)}))
	}

	return ds, nil
}

func (r *SatelliteRepository) GetAllSatellitesWithLastMessages() ([]*domain.Satellite, error) {
	ss, err := r.GetAllSatellites()
	if err != nil {
		return nil, err
	}

	for _, s := range ss {
		distance, message, err := r.GetLastMessageFrom(s.Name)
		if err != nil {
			return nil, err
		}
		s.ReceiveMessage(distance, message)
	}

	return ss, nil
}
