package endpoints

import (
	"context"
	"fmt"
	"github.com/OAAC/utils/userOperation"
)

type AggregateService struct{}

func (s *AggregateService) SendUserOperation(ctx context.Context, signedUserOp userOperation.SignedUserOperation) (string, error) {
	if len(signedUserOp.Signature) == 0 {
		return "", fmt.Errorf("invalid signature")
	}
	return "Hello", nil
}
