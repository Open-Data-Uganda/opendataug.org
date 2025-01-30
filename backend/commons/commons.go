package commons

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

func UUIDGenerator() string {
	u := uuid.NewV4().String()
	return strings.ReplaceAll(u, "-", "")
}
