package utils

import "strings"

func CompleteAddress(address string) string {
	const localhost = "localhost"

	if strings.HasPrefix(address, ":") {
		return localhost + address
	}

	return address
}
