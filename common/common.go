package common

import (
	"fmt"

	"github.com/gofrs/uuid"
)

func UUID() string {
	s, err := uuid.NewV4()
	if err != nil {
		return UUID()
	}
	return s.String()
}

func Message() string {
	return fmt.Sprintf("common package message via %s", UUID())
}
