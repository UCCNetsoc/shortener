package database

import (
	"errors"
	"fmt"

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
		return InitDatabase()
	}

	// migrate tablesS
	conn.AutoMigrate(&models.RedirectURL{})

	return &Client{conn}
}

// FindByHash returns a HashURL pointer
func (c *Client) FindByHash(slug string) *models.RedirectURL {
	var hashurl models.RedirectURL
	c.conn.Where("slug = ?", slug).First(&hashurl)
	return &hashurl
}

// CreateNew creates a new shortened url, and reutrns the error
func (c *Client) CreateNew(shortened *models.RedirectURL) (bool, error) {
	result := c.conn.Create(&shortened)
	if result.RowsAffected < 1 {
		return false, result.Error
	}
	return false, nil
}

// Duplicate checks for duplicate hashes
func (c *Client) Duplicate(slug string) bool {
	var preExisting models.RedirectURL
	err := c.conn.Where("slug = ?", slug).First(&preExisting).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
