package loader

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"github.com/roeyaus/zd/v2/common"
)

type FileLoader struct {
	orgFile    string
	ticketFile string
	userFile   string
}

func (c *FileLoader) Init(OrgFile, TicketFile, UserFile string) {
	c.orgFile = OrgFile
	c.ticketFile = TicketFile
	c.userFile = UserFile
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
	if err := json.Unmarshal([]byte(byteValue), orgs); err != nil {
		return nil, errors.Errorf("Couldn't unmarshall %v because %v", c.orgFile, err)
	}
	return orgs, nil
}
