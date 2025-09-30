package models

import (
	ierr "github.com/flexprice/flexprice/internal/errors"
)

// PriceSyncWorkflowInput represents input for the price sync workflow
type PriceSyncWorkflowInput struct {
	PlanID        string `json:"plan_id"`
	TenantID      string `json:"tenant_id"`
	EnvironmentID string `json:"environment_id"`
	UserID        string `json:"user_id"`
}

func (p *PriceSyncWorkflowInput) Validate() error {
	if p.PlanID == "" {
		return ierr.NewError("plan ID is required").
			WithHint("Plan ID is required").
			Mark(ierr.ErrValidation)
	}

	if p.TenantID == "" || p.EnvironmentID == "" || p.UserID == "" {
		return ierr.NewError("tenant ID, environment ID and user ID are required").
			WithHint("Tenant ID and environment ID are required").
			Mark(ierr.ErrValidation)
	}

	return nil
}
