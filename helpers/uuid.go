package helpers

import "github.com/google/uuid"

func RandomUuidAsString() string {
	return uuid.New().String()
}

func SafeUuidFromString(s string) uuid.UUID {
	val, _ := uuid.Parse(s)
	return val
}
