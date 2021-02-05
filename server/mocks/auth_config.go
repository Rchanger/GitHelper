package mocks

import (
	"context"

	"golang.org/x/oauth2"
)

type MockExchangeDS struct {
	Code string
	Err  error
}

func (m *MockExchangeDS) Exchange(ctx context.Context, code string, _ ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	m.Code = code
	return &oauth2.Token{}, m.Err
}
