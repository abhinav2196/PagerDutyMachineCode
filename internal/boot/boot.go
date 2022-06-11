package boot

import (
	"PagerDutyMachineCode/internal/developers"
	"PagerDutyMachineCode/internal/teams"
)

func Init() {
	developers.SetCore(&developers.Core{})
	teams.SetCore(&teams.Core{})
}
