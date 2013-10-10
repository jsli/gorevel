package models

import (
	"fmt"
	"github.com/coocood/qbs"
)

type Permissions struct {
	Id     int64
	UserId int64
	Perm   int
}

func (p *Permissions) Save() bool {
	q, err := qbs.GetQbs()
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer q.Close()

	_, err = q.Save(p)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
