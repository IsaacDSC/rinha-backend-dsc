package contract_repository

import (
	"context"

	"github.com/IsaacDSC/rinha-backend-dsc/internal/models"
)

type PersonContractRepository interface {
	Create(ctx context.Context, person models.Person) (err error)
	FindById(ctx context.Context, personID string) (output models.Person, err error)
	Search(ctx context.Context, person models.Person) (output []models.Person, err error)
	CounterPersons(ctx context.Context) (total_person int32, err error)
}
