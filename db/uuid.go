// Package server is using google's UUID generator to handle uuid generation with the least collisions
// 16 Byte UUID is most suitable / industry standard , but going for 5 byte to meet project requirements
package db

import (
	"github.com/google/uuid"
)

func GenerateUUID() string {
	UUID_16 := uuid.New().String()
	if len(UUID_16) != 36 {
		panic("UUID Generation Failed")
	}
	// trimming UUID to char(5) according to our project req
	UUID := UUID_16[:5]
	return UUID
}
