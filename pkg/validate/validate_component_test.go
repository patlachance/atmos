package validate

import (
	"github.com/stretchr/testify/assert"
	"testing"

	e "github.com/cloudposse/atmos/internal/exec"
	cfg "github.com/cloudposse/atmos/pkg/config"
	u "github.com/cloudposse/atmos/pkg/utils"
)

func TestValidateComponent(t *testing.T) {
	info := cfg.ConfigAndStacksInfo{}

	cliConfig, err := cfg.InitCliConfig(info, true)
	assert.Nil(t, err)

	_, err = e.ExecuteValidateComponent(cliConfig, info, "infra/vpc", "tenant1-ue2-dev", "validate-infra-vpc-component.rego", "opa")
	u.PrintError(err)
	assert.Error(t, err)
}

func TestValidateComponent2(t *testing.T) {
	info := cfg.ConfigAndStacksInfo{}

	cliConfig, err := cfg.InitCliConfig(info, true)
	assert.Nil(t, err)

	_, err = e.ExecuteValidateComponent(cliConfig, info, "infra/vpc", "tenant1-ue2-prod", "", "")
	u.PrintError(err)
	assert.Error(t, err)
}

func TestValidateComponent3(t *testing.T) {
	info := cfg.ConfigAndStacksInfo{}

	cliConfig, err := cfg.InitCliConfig(info, true)
	assert.Nil(t, err)

	_, err = e.ExecuteValidateComponent(cliConfig, info, "infra/vpc", "tenant1-ue2-staging", "", "")
	u.PrintError(err)
	assert.Error(t, err)
}
