package repositories

import "bws_test/app/models"

type Hackers interface {
	GetAll() ([]models.Hacker, error)
	Create(hacker models.Hacker) error
}
