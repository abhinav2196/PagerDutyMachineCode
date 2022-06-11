package developers

type IDevelopers interface {
	Create(name string, phoneNumber string, teamId string) (*Developer, error)
}

type Core struct{}

var impl IDevelopers

func SetCore(currImpl IDevelopers) {
	impl = currImpl
}

func GetCore() IDevelopers {
	return impl
}

func (c *Core) Create(name string, phoneNumber string, teamId string) (*Developer, error) {
	var dev = Developer{}

	dev.Name = name
	dev.PhoneNumber = phoneNumber
	dev.TeamId = teamId

	//call repo to store in db
	return &dev, nil
}
