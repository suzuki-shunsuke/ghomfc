package github

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"
)

type Member struct {
	Login          string `json:"login"`
	Name           string `json:"name"`
	NumOfFollowers int    `json:"number_of_followers"`
}

func (c *Client) ListOrgMembers(ctx context.Context, org string) ([]*Member, error) { //nolint:funlen,cyclop
	/*
		query ($org: String!, $after: String) {
		  organization(login: $org) {
		    membersWithRole(first: 100, after: $after) {
		      edges {
		        node {
		          login
		          name
		          followers {
		            totalCount
		          }
		        }
		      }
		    }
		  }
		}
	*/
	var q struct {
		Organization struct {
			MembersWithRole struct {
				Edges []struct {
					Node struct {
						Login     string
						Name      string
						Followers struct {
							TotalCount int
						}
					}
				}
				PageInfo struct {
					EndCursor   githubv4.String
					HasNextPage bool
				}
			} `graphql:"membersWithRole(first: 100, after: $cursor)"`
		} `graphql:"organization(login: $org)"`
	}
	variables := map[string]interface{}{
		"org":    org,
		"cursor": (*githubv4.String)(nil),
	}
	var members []*Member
	for range 30 {
		if err := c.v4Client.Query(ctx, &q, variables); err != nil {
			return nil, fmt.Errorf("get GitHub Organization members by GitHub GraphQL API: %w", err)
		}
		for _, edge := range q.Organization.MembersWithRole.Edges {
			members = append(members, &Member{
				Login:          edge.Node.Login,
				Name:           edge.Node.Name,
				NumOfFollowers: edge.Node.Followers.TotalCount,
			})
		}
		if !q.Organization.MembersWithRole.PageInfo.HasNextPage {
			return members, nil
		}
		variables["cursor"] = githubv4.NewString(q.Organization.MembersWithRole.PageInfo.EndCursor)
	}
	return members, nil
}
