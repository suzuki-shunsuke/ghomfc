package github

import (
	"context"
	"errors"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type Client struct {
	v4Client *githubv4.Client
}

func New(ctx context.Context, token string) (*Client, error) {
	if token == "" {
		return nil, errors.New("GitHub Access Token is required")
	}
	httpClient := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	))

	return &Client{
		v4Client: githubv4.NewClient(httpClient),
	}, nil
}
