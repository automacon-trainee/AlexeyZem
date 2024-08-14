package main

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-github/github"
)

type MockRepoLister struct {
}

func (m *MockRepoLister) List(ctx context.Context, username string, opt *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error) {
	if username == "wrong" {
		return nil, nil, fmt.Errorf("error")
	}

	var repos []*github.Repository
	for i := 0; i < 3; i++ {
		desc := fmt.Sprintf("repo%d", i)
		name := fmt.Sprintf("repoName%d", i)
		url := fmt.Sprintf("https://github.com/%s", name)
		repos = append(repos, &github.Repository{Description: &desc, Name: &name, HTMLURL: &url})
	}
	return repos, nil, nil
}

type MockGistLister struct{}

func (m *MockGistLister) List(ctx context.Context, username string, opt *github.GistListOptions) ([]*github.Gist, *github.Response, error) {
	if username == "wrong" {
		return nil, nil, fmt.Errorf("error")
	}
	var gists []*github.Gist
	for i := 0; i < 3; i++ {
		desc := fmt.Sprintf("repo%d", i)
		url := fmt.Sprintf("https://github.com/%s", desc)
		gists = append(gists, &github.Gist{Description: &desc, HTMLURL: &url})
	}
	return gists, nil, nil
}

func TestGithubGist(t *testing.T) {
	gist := GithubGist{getList: &MockGistLister{}}
	{
		_, err := gist.GetItems(context.Background(), "wrong")
		if err == nil {
			t.Errorf("GetGists should return an error")
		}
	}
	{
		res, err := gist.GetItems(context.Background(), "someName")
		if err != nil {
			t.Errorf("GetGists should not return an error")
		}

		want := []Item{
			{Description: "repo0", Title: "github.Gist{Description:\"repo0\", Files:map[], HTMLURL:\"https://github.com/repo0\"}", Link: "https://github.com/repo0"},
			{Description: "repo1", Title: "github.Gist{Description:\"repo1\", Files:map[], HTMLURL:\"https://github.com/repo1\"}", Link: "https://github.com/repo1"},
			{Description: "repo2", Title: "github.Gist{Description:\"repo2\", Files:map[], HTMLURL:\"https://github.com/repo2\"}", Link: "https://github.com/repo2"},
		}
		if !reflect.DeepEqual(res, want) {
			t.Errorf("GetGists should return %v, got %v", want, res)
		}
	}
}

func TestGithubRepo(t *testing.T) {
	repo := GithubRepo{repoList: &MockRepoLister{}}
	{
		_, err := repo.GetItems(context.Background(), "wrong")
		if err == nil {
			t.Errorf("GetRepos should return an error")
		}
	}
	{
		res, err := repo.GetItems(context.Background(), "someName")
		if err != nil {
			t.Errorf("GetRepos should not return an error")
		}
		want := []Item{{Description: "repo0", Title: "repoName0", Link: "https://github.com/repoName0"},
			{Description: "repo1", Title: "repoName1", Link: "https://github.com/repoName1"},
			{Description: "repo2", Title: "repoName2", Link: "https://github.com/repoName2"},
		}
		if !reflect.DeepEqual(res, want) {
			t.Errorf("GetRepos should return %v, got %v", want, res)
		}
	}
}

func TestGithubQuery(t *testing.T) {
	query := GithubQuery{}
	gist := GithubGist{getList: &MockGistLister{}}
	repo := GithubRepo{repoList: &MockRepoLister{}}
	{
		res, err := query.GetItems(context.Background(), "someName", &repo)
		if err != nil {
			t.Errorf("GetRepos should not return an error")
		}
		want := []Item{{Description: "repo0", Title: "repoName0", Link: "https://github.com/repoName0"},
			{Description: "repo1", Title: "repoName1", Link: "https://github.com/repoName1"},
			{Description: "repo2", Title: "repoName2", Link: "https://github.com/repoName2"},
		}
		if !reflect.DeepEqual(res, want) {
			t.Errorf("GetRepos should return %v, got %v", want, res)
		}
	}
	{
		res, err := query.GetItems(context.Background(), "someName", &gist)
		if err != nil {
			t.Errorf("GetGists should not return an error")
		}

		want := []Item{
			{Description: "repo0", Title: "github.Gist{Description:\"repo0\", Files:map[], HTMLURL:\"https://github.com/repo0\"}", Link: "https://github.com/repo0"},
			{Description: "repo1", Title: "github.Gist{Description:\"repo1\", Files:map[], HTMLURL:\"https://github.com/repo1\"}", Link: "https://github.com/repo1"},
			{Description: "repo2", Title: "github.Gist{Description:\"repo2\", Files:map[], HTMLURL:\"https://github.com/repo2\"}", Link: "https://github.com/repo2"},
		}
		if !reflect.DeepEqual(res, want) {
			t.Errorf("GetGists should return %v, got %v", want, res)
		}
	}
}
