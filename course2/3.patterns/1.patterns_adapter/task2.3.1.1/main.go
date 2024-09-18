package main

import (
	"context"

	"github.com/google/go-github/github"
)

type RepoLister interface {
	List(ctx context.Context, username string, opt *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error)
}

type GetLister interface {
	List(ctx context.Context, username string, opt *github.GistListOptions) ([]*github.Gist, *github.Response, error)
}

type Githuber interface {
	GetGists(ctx context.Context, username string) ([]Item, error)
	GetRepos(ctx context.Context, username string) ([]Item, error)
}

type GithubAdapter struct {
	RepoList RepoLister
	GetList  GetLister
}

func (g *GithubAdapter) GetGists(ctx context.Context, username string) ([]Item, error) {
	opt := &github.GistListOptions{ListOptions: github.ListOptions{PerPage: 1000}}
	gists, _, err := g.GetList.List(ctx, username, opt)
	if err != nil {
		return nil, err
	}

	res := make([]Item, len(gists))
	for i, gist := range gists {
		res[i].Description = gist.GetDescription()
		res[i].Link = gist.GetHTMLURL()
		res[i].Title = gist.String()
	}
	return res, nil
}

func (g *GithubAdapter) GetRepos(ctx context.Context, username string) ([]Item, error) {
	opt := &github.RepositoryListOptions{ListOptions: github.ListOptions{PerPage: 1000}}
	repos, _, err := g.RepoList.List(ctx, username, opt)
	if err != nil {
		return nil, err
	}

	res := make([]Item, len(repos))
	for i, repo := range repos {
		res[i] = Item{Description: repo.GetDescription(), Link: repo.GetHTMLURL(), Title: repo.GetName()}
	}
	return res, nil
}

func NewGithubAdapter(client *github.Client) *GithubAdapter {
	return &GithubAdapter{
		RepoList: client.Repositories,
		GetList:  client.Gists,
	}
}

type Item struct {
	Title       string
	Description string
	Link        string
}

func main() {}
