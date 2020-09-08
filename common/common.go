package common

import (
	"time"

	"github.com/google/uuid"
)

//Define our data structures
type Organisation struct {
	Id            string
	Url           string
	ExternalId    uuid.UUID
	Name          string
	DomainNames   []string
	CreatedAt     time.Time
	Details       string
	SharedTickets bool
	Tags          []string
}

type Ticket struct {
	Id           string
	Url          string
	ExternalId   uuid.UUID
	CreatedAt    time.Time
	Type         string
	Subject      string
	Description  string
	Priority     string
	Status       string
	SubmitterId  int
	AssigneeId   int
	OrgId        int
	Tags         []string
	HasIncidents bool
	DueAt        time.Time
	Via          string
}

type User struct {
	Id          string
	Url         string
	ExternalId  uuid.UUID
	CreatedAt   time.Time
	Name        string
	Alias       string
	Active      bool
	Verified    bool
	Shared      bool
	Locale      string
	Timezone    string
	LastLoginAt time.Time
	Email       string
	Phone       string
	Signature   string
	OrgId       int
	Tags        []string
	Suspended   bool
	Role        string
}

type Loader interface {
	LoadOrgs() []Organisation
	LoadTickets() []Ticket
	LoadUsers() []User
}
