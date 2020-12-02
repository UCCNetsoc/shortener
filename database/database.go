package database

import (
	"errors"
	"fmt"
	"log"

	// postgres driver
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/UCCNetsoc/shortener/models"
	"github.com/spf13/viper"
)

// Client interface to database connection
type Client struct {
	conn *gorm.DB
}

// InitDatabase initiates a db connection
func InitDatabase() *Client {
	conn, err := gorm.Open(postgres.Open(
		fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%s sslmode=disable",
			viper.GetString("db.user"), viper.GetString("db.password"), viper.GetString("db.host"),
			viper.GetString("db.database"), viper.GetString("db.port"))), &gorm.Config{},
	)
	if err != nil {
		log.Fatal("couldn't connect to database\n", err)
	}

	// migrate tables
	conn.AutoMigrate(&models.Request{})
	return &Client{conn}
}

// FindRedirect returns a request object, taking the domain and slug as params
func (c *Client) FindRedirect(domain, slug string) *models.Request {
	var req models.Request
	c.conn.Where("domain = ? AND slug = ?", domain, slug).First(&req)
	return &req
}

// CreateNew creates a new shortened url, and reutrns the error
func (c *Client) CreateNew(req *models.Request) (bool, error) {
	result := c.conn.Create(&req)
	if result.RowsAffected < 1 {
		return false, result.Error
	}
	return false, nil
}

// Duplicate checks for duplicate hashes
func (c *Client) Duplicate(domain, slug string) bool {
	var preExisting models.Request
	err := c.conn.Where("domain = ? AND slug = ?", domain, slug).First(&preExisting).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
