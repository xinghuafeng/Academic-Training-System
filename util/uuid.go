package util

import (
	"github.com/satori/go.uuid"
	"encoding/hex"
)

/*
	生成 uuid
*/
func GetUuid() string {
	bytes := uuid.NewV4()
	uuid := hex.EncodeToString(bytes[:])
	return uuid
}
