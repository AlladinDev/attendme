package interfaces

type IRepo struct {
	CreateContractor func(details *models.Contractor) error
}
