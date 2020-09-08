package searcher

import "github.com/roeyaus/zd/common"

type Searcher struct {
	orgs    []common.Organisation
	tickets []common.Ticket
	users   []common.User
}

func NewSearcher(loader common.Loader) *Searcher {
	return &Searcher{orgs: loader.LoadOrgs(), tickets: loader.LoadTickets(), users: loader.LoadUsers()}
}
