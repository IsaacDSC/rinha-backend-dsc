package di

import (
	services "github.com/IsaacDSC/rinha-backend-dsc/internal/app"
	"github.com/IsaacDSC/rinha-backend-dsc/internal/infra/repositories"
	"github.com/IsaacDSC/rinha-backend-dsc/internal/models"
)

func GetInstanceService(person models.Person) services.PersonServiceContract {
	return services.NewServicePerson(
		models.Person(person),
		new(repositories.PersonRepository),
	)
}
