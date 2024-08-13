package main

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-github/v53/github"
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

func TestGithubAdapter(t *testing.T) {
	adapter := GithubAdapter{repoLister: &MockRepoLister{}, gistLister: &MockGistLister{}}

	{
		_, err := adapter.GetGists(context.Background(), "wrong")
		if err == nil {
			t.Errorf("GetGists should return an error")
		}
		_, err = adapter.GetRepos(context.Background(), "wrong")
		if err == nil {
			t.Errorf("GetRepos should return an error")
		}
	}

	{
		res, err := adapter.GetRepos(context.Background(), "someName")
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
		res, err := adapter.GetGists(context.Background(), "someName")
		if err != nil {
			t.Errorf("GetRepos should not return an error")
		}
		want := []Item{
			{Description: "repo0", Title: "github.Gist{Description:\"repo0\", HTMLURL:\"https://github.com/repo0\"}", Link: "https://github.com/repo0"},
			{Description: "repo1", Title: "github.Gist{Description:\"repo1\", HTMLURL:\"https://github.com/repo1\"}", Link: "https://github.com/repo1"},
			{Description: "repo2", Title: "github.Gist{Description:\"repo2\", HTMLURL:\"https://github.com/repo2\"}", Link: "https://github.com/repo2"},
		}
		if !reflect.DeepEqual(res, want) {
			t.Errorf("GetGists should return %v, got %v", want, res)
		}
	}
}

func TestGithubProxy(t *testing.T) {
	adapter := GithubAdapter{repoLister: &MockRepoLister{}, gistLister: &MockGistLister{}}
	proxy := NewGithubProxy(&adapter)
	{
		_, err := proxy.GetGists(context.Background(), "wrong")
		if err == nil {
			t.Errorf("GetGists should return an error")
		}

		_, err = proxy.GetRepos(context.Background(), "wrong")
		if err == nil {
			t.Errorf("GetRepos should return an error")
		}
	}

	{
		start := time.Now()
		res, err := proxy.GetRepos(context.Background(), "someName")
		end := time.Since(start)
		if err != nil {
			t.Errorf("GetRepos should not return an error")
		}
		startC := time.Now()
		want, _ := proxy.GetRepos(context.Background(), "someName")
		endC := time.Since(startC)
		if !reflect.DeepEqual(res, want) {
			t.Errorf("GetRepos should return %v, got %v", want, res)
		}
		if 2*endC > end*3 {
			t.Errorf("GetRepos with cache too long")
		}
	}
	{
		start := time.Now()
		res, err := proxy.GetGists(context.Background(), "someName")
		end := time.Since(start)
		if err != nil {
			t.Errorf("GetRepos should not return an error")
		}
		startC := time.Now()
		want, _ := proxy.GetGists(context.Background(), "someName")
		endC := time.Since(startC)
		if !reflect.DeepEqual(res, want) {
			t.Errorf("GetRepos should return %v, got %v", want, res)
		}
		if 2*endC > end*3 {
			t.Errorf("GetRepos with cache too long")
		}
	}
}
