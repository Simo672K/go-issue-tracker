package main

import (
	"context"
	"testing"

	"github.com/Simo672K/issue-tracker/internal/auth"
)

func TestHasAccessTo(t *testing.T) {
	ctx := context.Background()
	ownerTest := auth.NewPermission(auth.OWNER)
	profileId := "38e3430d-780f-4c91-8f9a-1434f9506d58"
	projectId := "f597d870-01fd-41b6-b034-7c7c72a727ad"

	if !ownerTest.HasAccessTo(ctx, profileId, projectId) {
		t.Fail()
	}
}
