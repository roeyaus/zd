package loader

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"github.com/roeyaus/zd/common"
)

type FileLoader struct {
	orgFile    string
	ticketFile string
	userFile   string
}

func (c *FileLoader) NewFileLoader(OrgFile, TicketFile, UserFile string) *FileLoader {
	return &FileLoader{orgFile: OrgFile,
		ticketFile: TicketFile,
		userFile:   UserFile}
}

func (c FileLoader) LoadOrgs() ([]common.Organisation, error) {
	jsonFile, err := os.Open(c.orgFile)
	if err != nil {
		return nil, errors.Errorf("Couldn't open %v because %v", c.orgFile, err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, errors.Errorf("Couldn't read %v because %v", c.orgFile, err)
	}
	//we load our
	var orgs []common.Organisation
	if err := json.Unmarshal([]byte(byteValue), &orgs); err != nil {
		return nil, errors.Errorf("Couldn't unmarshall %v because %v", c.orgFile, err)
	}
	return orgs, nil
}
func (c FileLoader) LoadTickets() ([]common.Ticket, error) {
	jsonFile, err := os.Open(c.ticketFile)
	if err != nil {
		return nil, errors.Errorf("Couldn't open %v because %v", c.ticketFile, err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, errors.Errorf("Couldn't read %v because %v", c.ticketFile, err)
	}
	//we load our
	var tickets []common.Ticket
	if err := json.Unmarshal([]byte(byteValue), &tickets); err != nil {
		return nil, errors.Errorf("Couldn't unmarshall %v because %v", c.ticketFile, err)
	}
	return tickets, nil
}

func (c FileLoader) LoadUsers() ([]common.User, error) {
	jsonFile, err := os.Open(c.userFile)
	if err != nil {
		return nil, errors.Errorf("Couldn't open %v because %v", c.userFile, err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, errors.Errorf("Couldn't read %v because %v", c.userFile, err)
	}
	//we load our
	var users []common.User
	if err := json.Unmarshal([]byte(byteValue), &users); err != nil {
		return nil, errors.Errorf("Couldn't unmarshall %v because %v", c.userFile, err)
	}
	return users, nil
}
