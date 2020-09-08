package loader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadOrgs(t *testing.T) {
	fileLoader := &FileLoader{}
	fileLoader.Init("../common/data/organizations.json", "", "")
	orgs, err := fileLoader.LoadOrgs()
	assert.NoError(t, err)
	assert.NotNil(t, orgs)

}
