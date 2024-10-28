package utils

import (
	"strconv"
	"time"

	"github.com/google/uuid"
)

func StrUniqueId() string {
	currentTimestamp := time.Now().Nanosecond()
	newId := uuid.New().ID()
	ID := currentTimestamp + int(newId)
	strID := strconv.FormatInt(int64(ID), 10)

	return strID
}
