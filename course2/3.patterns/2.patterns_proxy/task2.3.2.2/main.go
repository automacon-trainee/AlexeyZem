package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v53/github"
)

type RepoLister interface {
	List(ctx context.Context, username string, opt *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error)
}

type GistLister interface {
	List(ctx context.Context, username string, opt *github.GistListOptions) ([]*github.Gist, *github.Response, error)
}

type Githuber interface {
	GetGists(ctx context.Context, username string) ([]Item, error)
	GetRepos(ctx context.Context, username string) ([]Item, error)
}

type GithubAdapter struct {
	repoLister RepoLister
	gistLister GistLister
}

func (g *GithubAdapter) GetGists(ctx context.Context, username string) ([]Item, error) {
	opt := &github.GistListOptions{ListOptions: github.ListOptions{PerPage: 1000}}
	gists, _, err := g.gistLister.List(ctx, username, opt)
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
	repos, _, err := g.repoLister.List(ctx, username, opt)
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
		repoLister: client.Repositories,
		gistLister: client.Gists,
	}
}

type GithubProxy struct {
	github Githuber
	cache  map[string][]Item
}

func NewGithubProxy(client Githuber) *GithubProxy {
	return &GithubProxy{
		github: client,
		cache:  make(map[string][]Item),
	}
}

func (gh *GithubProxy) GetGists(ctx context.Context, username string) ([]Item, error) {
	key := fmt.Sprintf("gists_%s", username)
	if items, ok := gh.cache[key]; ok {
		return items, nil
	}
	items, err := gh.github.GetGists(ctx, username)
	if err != nil {
		return nil, err
	}
	gh.cache[key] = items
	return items, nil
}

func (gh *GithubProxy) GetRepos(ctx context.Context, username string) ([]Item, error) {
	key := fmt.Sprintf("repos_%s", username)
	if items, ok := gh.cache[key]; ok {
		return items, nil
	}
	items, err := gh.github.GetRepos(ctx, username)
	if err != nil {
		return nil, err
	}
	gh.cache[key] = items
	return items, nil
}

type Item struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
}
