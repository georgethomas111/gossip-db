package member

import (
	"fmt"

	"github.com/hashicorp/memberlist"
)

func NewSwim(port int, knownSwin int) ([]int, error) {
	c := memberlist.DefaultLocalConfig()
	c.BindPort = port

	list, err := memberlist.Create(c)
	if err != nil {
		return nil, err
	}

	if port != 7000 {
		s, err := list.Join([]string{known})
		if err != nil {
			return nil, err
		}
		fmt.Println("Known = ", known, " s =", s)
	}

	var swimPorts []int
	members := list.Members()
	for _, m := range members {
		swimPorts = append(swimPorts, int(m.Port))

	}
	return swimPorts, nil
}
