package inbound

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mikemonzo/goshare/internal/tenant/application"
	"github.com/mikemonzo/goshare/internal/tenant/domain"
)

type CreateTenantUseCase interface {
	Execute(input application.CreateTenantInput) (*domain.Tenant, error)
}

type TenantHandler struct {
	CreateTenantUseCase CreateTenantUseCase
}

func NewTenantHandler(createTenantUseCase CreateTenantUseCase) *TenantHandler {
	return &TenantHandler{
		CreateTenantUseCase: createTenantUseCase,
	}
}

type createTenantRequest struct {
	Name     string `json:"name" binding:"required"`
	Branding string `json:"branding"`
}

type createTenantResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Branding  string    `json:"branding"`
	CreatedAt int64     `json:"created_at"`
}

func (h *TenantHandler) CreateTenant(c *gin.Context) {
	var req createTenantRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := application.CreateTenantInput{
		Name:     req.Name,
		Branding: req.Branding,
	}

	tenant, err := h.CreateTenantUseCase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := createTenantResponse{
		ID:        tenant.ID,
		Name:      tenant.Name,
		Branding:  tenant.Branding,
		CreatedAt: tenant.CreatedAt,
	}

	c.JSON(http.StatusCreated, resp)
}
