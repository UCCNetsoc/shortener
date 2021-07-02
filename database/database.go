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
			viper.GetString("db.user"), viper.GetString("db.pass"), viper.GetString("db.host"),
			viper.GetString("db.name"), viper.GetString("db.port"))), &gorm.Config{},
	)
	if err != nil {
		log.Println("couldn't connect to database\n", err)
		log.Println("Password: ", viper.GetString("db.pass"))
	}

	// migrate tables
	conn.AutoMigrate(&models.Link{})
	return &Client{conn}
}

// FetchLinks gets a full dump of links from db
func (c *Client) FetchLinks() *[]models.Link {
	var links []models.Link
	c.conn.Find(&links)
	return &links
}

// FindRedirect returns a request object, taking the domain and slug as params
func (c *Client) FindRedirect(slug string) *models.Link {
	var req models.Link
	c.conn.Where("slug = ?", slug).First(&req)
	return &req
}

// CreateNew creates a new shortened url, and reutrns the error
func (c *Client) CreateNew(req *models.Link) (bool, error) {
	result := c.conn.Create(&req)
	if result.RowsAffected < 1 {
		return false, result.Error
	}
	return true, nil
}

// DeleteSlug deletes a slug
func (c *Client) DeleteSlug(slug string) (bool, error) {
	req := &models.Link{Slug: slug}
	result := c.conn.Delete(req)
	if result.RowsAffected < 1 {
		return false, result.Error
	}
	return true, nil
}

// Duplicate checks for duplicate hashes
func (c *Client) Duplicate(slug string) bool {
	var preExisting models.Link
	err := c.conn.Where("slug = ?", slug).First(&preExisting).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
