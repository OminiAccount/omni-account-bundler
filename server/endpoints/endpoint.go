package endpoints

import (
	"context"
	"fmt"
)

type MyBusinessService struct{}

func (s *MyBusinessService) Hello(ctx context.Context, name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("name cannot be empty")
	}
	return "Hello, " + name, nil
}
