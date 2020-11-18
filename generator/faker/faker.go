package faker

import (
	"github.com/FoksVHox/algoritime-for-htx/generator/faker/factories"
)

func GetNames() string {
	return names.Generate()
}
