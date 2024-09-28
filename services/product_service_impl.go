package services

import (
	"errors"
	"simplerapi/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productServiceImpl struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) ProductService {
	return &productServiceImpl{db: db}
}

func (s *productServiceImpl) CreateProduct(c *gin.Context) (models.Product, error) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		return product, err
	}

	// Check if the product already exists
	var existingProduct models.Product
	if err := s.db.Where("name = ?", product.Name).First(&existingProduct).Error; err == nil {
		return product, gorm.ErrDuplicatedKey // Return error if product already exists
	}

	if err := s.db.Create(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (s *productServiceImpl) GetProductById(c *gin.Context) (models.Product, error) {
	id := c.Param("id")
	var product models.Product

	if err := s.db.First(&product, id).Error; err != nil {
		return product, gorm.ErrRecordNotFound
	}

	return product, nil
}

func (s *productServiceImpl) GetProducts(c *gin.Context) ([]models.Product, error) {
	var products []models.Product
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")

	var offset int

	if page == "1" && pageSize == "10" {
		if err := s.db.Find(&products).Error; err != nil {
			return nil, err
		}
		return products, nil
	}

	if p, err := strconv.Atoi(page); err == nil && p > 0 {
		offset = (p - 1) * 10
	} else {
		return nil, errors.New("invalid page number")
	}

	ps, err := strconv.Atoi(pageSize)
	if err != nil || ps <= 0 {
		ps = 10
	}

	if err := s.db.Limit(ps).Offset(offset).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (s *productServiceImpl) UpdateProduct(c *gin.Context) (models.Product, error) {
	id := c.Param("id")
	var product models.Product

	// Check if the product exists
	if err := s.db.First(&product, id).Error; err != nil {
		return product, gorm.ErrRecordNotFound
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		return product, err
	}

	if err := s.db.Save(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (s *productServiceImpl) DeleteProduct(c *gin.Context) error {
	id := c.Param("id")
	var product models.Product

	if err := s.db.Delete(&product, id).Error; err != nil {
		return gorm.ErrRecordNotFound
	}

	c.JSON(204, gin.H{})
	return nil
}
