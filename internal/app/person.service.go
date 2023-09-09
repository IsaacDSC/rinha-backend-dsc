package services

import (
	"context"
	"log"

	"github.com/IsaacDSC/rinha-backend-dsc/internal/models"
	contract_repository "github.com/IsaacDSC/rinha-backend-dsc/shared/contracts/repositories"

	"github.com/docker/distribution/uuid"
)

type PersonServiceContract interface {
	CreatePerson(ctx context.Context) (models.Person, error)
	RetrievePerson(ctx context.Context) (models.Person, error)
	SearchPerson(ctx context.Context) ([]models.Person, error)
	TotalPerson(ctx context.Context) int32
}

type PersonService struct {
	person           models.Person
	repositoryPerson contract_repository.PersonContractRepository
}

func NewServicePerson(
	model models.Person,
	repository contract_repository.PersonContractRepository,
) *PersonService {
	p := new(PersonService)
	p.person = model
	p.repositoryPerson = repository
	return p
}

func (p *PersonService) CreatePerson(ctx context.Context) (output models.Person, err error) {
	p.person.ID = uuid.Generate().String()
	err = p.repositoryPerson.Create(ctx, p.person)
	if err != nil {
		return
	}
	output = p.person
	return
}

func (p *PersonService) RetrievePerson(ctx context.Context) (models.Person, error) {
	person, err := p.repositoryPerson.FindById(ctx, p.person.ID)
	if err != nil {
		log.Fatal(err)
	}
	return person, err
}

func (p *PersonService) SearchPerson(ctx context.Context) ([]models.Person, error) {
	return p.repositoryPerson.Search(ctx, p.person)
}

func (p *PersonService) TotalPerson(ctx context.Context) int32 {
	totalPerson, err := p.repositoryPerson.CounterPersons(ctx)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return totalPerson
}
