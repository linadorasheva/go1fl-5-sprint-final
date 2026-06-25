package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(val string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, v := range dataset {

		err := dp.Parse(v)

		if err != nil {
			log.Println(err)
			continue
		}

		str, err := dp.ActionInfo()

		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println(str)
	}
}
