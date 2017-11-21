package commands

import "github.com/pivotal-cf/om/models"

//go:generate counterfeiter -o ./fakes/presenter.go --fake-name Presenter . presenter

type presenter interface {
	PresentInstallations([]models.Installation)
	PresentAvailableProducts([]models.Product)
}