package services

type Services struct {
	Status Status
	Config Config
}

func NewLogic() *Services {
	return &Services{
		Status: NewStatusLogic(),
		Config: NewConfigLogic(),
	}
}
