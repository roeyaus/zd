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
func TestLoadTickets(t *testing.T) {
	fileLoader := &FileLoader{}
	fileLoader.Init("", "../common/data/tickets.json", "")
	tickets, err := fileLoader.LoadTickets()
	assert.NoError(t, err)
	assert.NotNil(t, tickets)

}
func TestLoadUsers(t *testing.T) {
	fileLoader := &FileLoader{}
	fileLoader.Init("", "", "../common/data/users.json")
	users, err := fileLoader.LoadUsers()
	assert.NoError(t, err)
	assert.NotNil(t, users)

}
