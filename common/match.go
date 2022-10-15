package common

import (
	"fmt"
	"strings"
)

func Match(list []string) {
	for _, i := range list {
		switch {
		case strings.Contains(i, "nginx"):
			fmt.Println("is nginx")
		case strings.Contains(i, "master"):
			fmt.Println("is master")
		}
		//return true
	}
}
