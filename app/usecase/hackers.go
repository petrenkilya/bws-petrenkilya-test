package usecase

import "bws_test/app/models"

type Hackers interface {
	Create(hacker models.Hacker) error
	Get() ([]models.Hacker, error)
}
