package impl

import (
	"bws_test/app/models"
	"bws_test/app/repositories"
)

type Hackers struct {
	repo repositories.Hackers
}

func CreateHackersUseCase(repo repositories.Hackers) *Hackers {
	return &Hackers{repo: repo}
}

func (u *Hackers) Create(hacker models.Hacker) error {
	return u.repo.Create(hacker)
}

func (u *Hackers) Get() ([]models.Hacker, error) {
	return u.repo.GetAll()
}
