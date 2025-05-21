package inbound_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mikemonzo/goshare/internal/tenant/application"
	"github.com/mikemonzo/goshare/internal/tenant/domain"
	"github.com/mikemonzo/goshare/internal/tenant/infrastructure/adapter/inbound"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCreateTenantUseCase struct {
	mock.Mock
}

func (m *MockCreateTenantUseCase) Execute(input application.CreateTenantInput) (*domain.Tenant, error) {
	args := m.Called(input)
	return args.Get(0).(*domain.Tenant), args.Error(1)
}

func TestTenantHandler_CreateTenant(t *testing.T) {
	mockUseCase := new(MockCreateTenantUseCase)
	handler := inbound.NewTenantHandler(mockUseCase)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/tenants", handler.CreateTenant)

	tenant := &domain.Tenant{
		ID:        uuid.New(),
		Name:      "Test Tenant",
		Branding:  "Test Branding",
		CreatedAt: 1234567890,
	}

	mockUseCase.On("Execute", application.CreateTenantInput{
		Name:     "Test Tenant",
		Branding: "Test Branding",
	}).Return(tenant, nil)

	body, _ := json.Marshal(map[string]string{
		"name":     "Test Tenant",
		"branding": "Test Branding",
	})

	req, _ := http.NewRequest(http.MethodPost, "/tenants", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var response map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &response)

	assert.Equal(t, tenant.ID.String(), response["id"])
	assert.Equal(t, tenant.Name, response["name"])
	assert.Equal(t, tenant.Branding, response["branding"])

	mockUseCase.AssertExpectations(t)
	mockUseCase.AssertCalled(t, "Execute", application.CreateTenantInput{ // Validar llamada al mock
		Name:     "Test Tenant",
		Branding: "Test Branding",
	})
}
