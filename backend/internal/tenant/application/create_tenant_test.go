package application_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/mikemonzo/goshare/internal/tenant/application"
	"github.com/mikemonzo/goshare/internal/tenant/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTenantRepository es un mock de la interfaz TenantRepository.
type MockTenantRepository struct {
	mock.Mock
}

func (m *MockTenantRepository) Create(tenant *domain.Tenant) error {
	args := m.Called(tenant)
	return args.Error(0)
}

func (m *MockTenantRepository) GetByID(id uuid.UUID) (*domain.Tenant, error) {
	return nil, nil
}

func (m *MockTenantRepository) GetByName(name string) (*domain.Tenant, error) {
	return nil, nil
}

func (m *MockTenantRepository) Update(tenant *domain.Tenant) error {
	return nil
}

func (m *MockTenantRepository) Delete(id uuid.UUID) error {
	return nil
}

func (m *MockTenantRepository) ListActive() ([]*domain.Tenant, error) {
	return nil, nil
}

func (m *MockTenantRepository) GetAll() ([]*domain.Tenant, error) {
	return nil, nil
}

func TestCreateTenantUseCase_Execute(t *testing.T) {
	mockRepo := new(MockTenantRepository)
	useCase := application.NewCreateTenantUseCase(mockRepo)

	input := application.CreateTenantInput{
		Name:     "Test Tenant",
		Branding: "Test Branding",
	}

	mockRepo.On("Create", mock.Anything).Return(nil)

	tenant, err := useCase.Execute(input)

	assert.NoError(t, err)
	assert.Equal(t, "Test Tenant", tenant.Name)
	assert.Equal(t, "Test Branding", tenant.Branding)
	assert.True(t, tenant.IsActive)
	assert.WithinDuration(t, time.Now(), time.Unix(tenant.CreatedAt, 0), time.Second)
	mockRepo.AssertExpectations(t)
}
