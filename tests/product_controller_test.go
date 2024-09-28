package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"simplerapi/controllers"
	"simplerapi/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) CreateProduct(c *gin.Context) (models.Product, error) {
	args := m.Called(c)
	return args.Get(0).(models.Product), args.Error(1)
}

func (m *MockProductService) GetProductById(c *gin.Context) (models.Product, error) {
	args := m.Called(c)
	return args.Get(0).(models.Product), args.Error(1)
}

func (m *MockProductService) GetProducts(c *gin.Context) ([]models.Product, error) {
	args := m.Called(c)
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *MockProductService) UpdateProduct(c *gin.Context) (models.Product, error) {
	args := m.Called(c)
	return args.Get(0).(models.Product), args.Error(1)
}

func (m *MockProductService) DeleteProduct(c *gin.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func performRequest(router *gin.Engine, method, path string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return rec
}

func setupController(mockService *MockProductService) (*gin.Engine, *controllers.ProductController) {
	gin.SetMode(gin.TestMode)
	controller := controllers.NewProductController(mockService)
	router := gin.Default()
	return router, controller
}

func TestProductController_CreateProduct(t *testing.T) {
	mockService := new(MockProductService)
	router, controller := setupController(mockService)
	router.POST("/products", controller.CreateProduct)

	product := models.Product{Name: "Test Product"}
	mockService.On("CreateProduct", mock.Anything).Return(product, nil)

	body, _ := json.Marshal(product)
	rec := performRequest(router, http.MethodPost, "/products", body)

	assert.Equal(t, http.StatusCreated, rec.Code)

	var response models.Product
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, product, response)

	mockService.AssertExpectations(t)
}

func TestProductController_CreateProduct_Error(t *testing.T) {
	mockService := new(MockProductService)
	router, controller := setupController(mockService)
	router.POST("/products", controller.CreateProduct)

	mockService.On("CreateProduct", mock.Anything).Return(models.Product{}, errors.New("some error"))

	body := []byte(`{"name": "Test Product"}`)
	rec := performRequest(router, http.MethodPost, "/products", body)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "some error", response["error"])

	mockService.AssertExpectations(t)
}

func TestProductController_GetProductById(t *testing.T) {
	mockService := new(MockProductService)
	router, controller := setupController(mockService)
	router.GET("/products/:id", controller.GetProductById)

	product := models.Product{Name: "Test Product"}
	mockService.On("GetProductById", mock.Anything).Return(product, nil)

	rec := performRequest(router, http.MethodGet, "/products/1", nil)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response models.Product
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, product, response)

	mockService.AssertExpectations(t)
}

func TestProductController_GetProductById_Error(t *testing.T) {
	mockService := new(MockProductService)
	router, controller := setupController(mockService)
	router.GET("/products/:id", controller.GetProductById)

	mockService.On("GetProductById", mock.Anything).Return(models.Product{}, errors.New("product not found"))

	rec := performRequest(router, http.MethodGet, "/products/1", nil)

	assert.Equal(t, http.StatusNotFound, rec.Code)
	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "product not found", response["error"])

	mockService.AssertExpectations(t)
}

func TestProductController_GetProducts(t *testing.T) {
	mockService := new(MockProductService)
	router, controller := setupController(mockService)
	router.GET("/products", controller.GetProducts)

	products := []models.Product{
		{Name: "Product 1"},
		{Name: "Product 2"},
	}
	mockService.On("GetProducts", mock.Anything).Return(products, nil)

	rec := performRequest(router, http.MethodGet, "/products", nil)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response []models.Product
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, products, response)

	mockService.AssertExpectations(t)
}

func TestProductController_GetProducts_Error(t *testing.T) {
	mockService := new(MockProductService)
	router, controller := setupController(mockService)
	router.GET("/products", controller.GetProducts)

	mockService.On("GetProducts", mock.Anything).Return(nil, errors.New("some error"))

	rec := performRequest(router, http.MethodGet, "/products", nil)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)

	var response map[string]string
	assert.Empty(t, response)
	mockService.AssertExpectations(t)
}

func TestProductController_UpdateProduct(t *testing.T) {
	mockService := new(MockProductService)
	router, controller := setupController(mockService)
	router.PUT("/products/:id", controller.UpdateProduct)

	product := models.Product{Name: "Updated Product"}
	mockService.On("UpdateProduct", mock.Anything).Return(product, nil)

	body, _ := json.Marshal(product)
	rec := performRequest(router, http.MethodPut, "/products/1", body)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response models.Product
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, product, response)

	mockService.AssertExpectations(t)
}

func TestProductController_UpdateProduct_Error(t *testing.T) {
	mockService := new(MockProductService)
	router, controller := setupController(mockService)
	router.PUT("/products/:id", controller.UpdateProduct)

	mockService.On("UpdateProduct", mock.Anything).Return(models.Product{}, errors.New("update failed"))

	body := []byte(`{"name": "Updated Product"}`)
	rec := performRequest(router, http.MethodPut, "/products/1", body)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	log.Print(&response)

	assert.Equal(t, "update failed", "update failed")

	mockService.AssertExpectations(t)
}

func TestProductController_DeleteProduct(t *testing.T) {
	mockService := new(MockProductService)
	router, controller := setupController(mockService)
	router.DELETE("/products/:id", controller.DeleteProduct)

	mockService.On("DeleteProduct", mock.Anything).Return(nil)

	rec := performRequest(router, http.MethodDelete, "/products/1", nil)

	assert.Equal(t, http.StatusNoContent, rec.Code)

	mockService.AssertExpectations(t)
}

func TestProductController_DeleteProduct_Error(t *testing.T) {
	mockService := new(MockProductService)
	router, controller := setupController(mockService)
	router.DELETE("/products/:id", controller.DeleteProduct)

	mockService.On("DeleteProduct", mock.Anything).Return(errors.New("product not found"))

	rec := performRequest(router, http.MethodDelete, "/products/1", nil)

	assert.Equal(t, http.StatusNotFound, rec.Code)
	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "product not found", response["error"])

	mockService.AssertExpectations(t)
}
