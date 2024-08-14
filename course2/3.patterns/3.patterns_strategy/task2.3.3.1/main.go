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

type Item struct {
	Title       string
	Description string
	Link        string
}

type GithubLister interface {
	GetItems(ctx context.Context, username string) ([]Item, error)
}

type GithubGist struct {
	getList GetLister
}

func (g *GithubGist) GetItems(ctx context.Context, username string) ([]Item, error) {
	opt := &github.GistListOptions{ListOptions: github.ListOptions{PerPage: 1000}}
	gists, _, err := g.getList.List(ctx, username, opt)
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

type GithubRepo struct {
	repoList RepoLister
}

func (g *GithubRepo) GetItems(ctx context.Context, username string) ([]Item, error) {
	opt := &github.RepositoryListOptions{ListOptions: github.ListOptions{PerPage: 1000}}
	repos, _, err := g.repoList.List(ctx, username, opt)
	if err != nil {
		return nil, err
	}

	res := make([]Item, len(repos))
	for i, repo := range repos {
		res[i] = Item{Description: repo.GetDescription(), Link: repo.GetHTMLURL(), Title: repo.GetName()}
	}
	return res, nil
}

type GeneralGithubLister interface {
	GetItems(ctx context.Context, username string, strategy GithubLister) ([]Item, error)
}

type GithubQuery struct {
}

func (g *GithubQuery) GetItems(ctx context.Context, username string, strategy GithubLister) ([]Item, error) {
	return strategy.GetItems(ctx, username)
}

func main() {}
