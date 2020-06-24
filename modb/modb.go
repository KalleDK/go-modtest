package modb

import (
	"crypto"
	"fmt"

	"github.com/google/uuid"
)

func ModB() crypto.Hash {
	u := uuid.New()
	fmt.Println(u)
	return crypto.SHA256
}
