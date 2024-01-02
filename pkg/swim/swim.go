package swim

import (
	"fmt"

	"github.com/hashicorp/memberlist"
)

var FirstNode = false

func New(port int, known string) ([]string, error) {
	c := memberlist.DefaultLocalConfig()
	c.BindPort = port

	list, err := memberlist.Create(c)
	if err != nil {
		return nil, err
	}

	if FirstNode == false {
		s, err := list.Join([]string{known})
		if err != nil {
			return nil, err
		}
		fmt.Println("Known = ", known, " s =", s)
	}

	var others []string
	members := list.Members()
	fmt.Println("members =", members)
	for _, m := range members {
		others = append(others, m.Addr.String())

	}
	return others, nil
}
