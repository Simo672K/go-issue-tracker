package service

import (
	"fmt"

	"github.com/Simo672K/issue-tracker/utils"
)

type Verification struct {
	VerifID    string
	IsVerified bool
	UserId     string
}

var VerifiedUsers = make(map[string]*Verification)

func newVerification(userId, verifID string) *Verification {
	return &Verification{
		IsVerified: false,
		UserId:     userId,
		VerifID:    verifID,
	}
}

func CreateVerification(userId string) string {
	uniqueVerfId := utils.StrUniqueId()
	VerifiedUsers[uniqueVerfId] = newVerification(userId, uniqueVerfId)

	return uniqueVerfId
}

func ValidateVerification(verifId string) error {
	if _, ok := VerifiedUsers[verifId]; !ok {
		return fmt.Errorf("invalid verification id")
	}

	VerifiedUsers[verifId].IsVerified = true
	fmt.Println(VerifiedUsers)
	return nil
}

func UserEmailVerificationMockService() string {
	userID := utils.StrUniqueId()
	verifId := CreateVerification(userID)

	return verifId
}
