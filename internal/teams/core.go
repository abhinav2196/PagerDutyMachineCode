package teams

import devPkg "PagerDutyMachineCode/internal/developers"

type ITeams interface {
	Create(name string, developers []devPkg.Developer) (*Team, error)
}

type Core struct{}

var impl ITeams

func SetCore(currImpl ITeams) {
	impl = currImpl
}

func GetCore() ITeams {
	return impl
}

func (c *Core) Create(name string, developers []devPkg.Developer) (*Team, error) {
	var team = Team{}

	team.Name = name

	//call repo to store in db

	for _, dev := range developers {
		_, err := devPkg.GetCore().Create(dev.Name, dev.PhoneNumber, team.Id)
		if err != nil {
			return nil, err
		}
	}

	//Ideally would want it to be a txn in a same repo, so that rollback is easier

	//call repo to store in db

	return &team, nil
}
