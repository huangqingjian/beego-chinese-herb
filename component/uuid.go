package component

import (
	"github.com/satori/go.uuid"
	"strings"
)

// 获取uuid
func GetUuid() string {
	uuid := uuid.NewV4()
	id := strings.ReplaceAll(uuid.String(), "-", "")
	return id
}