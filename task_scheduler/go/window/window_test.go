package window

import (
	"context"
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.skia.org/infra/go/git/repograph"
	git_testutils "go.skia.org/infra/go/git/testutils"
	"go.skia.org/infra/go/testutils"
	"go.skia.org/infra/go/testutils/unittest"
)

// A Window with no repos should just be a time range check.
func TestWindowNoRepos(t *testing.T) {
	unittest.SmallTest(t)
	period := time.Hour
	w, err := New(context.Background(), period, 0, nil)
	require.NoError(t, err)
	now, err := time.Parse(time.RFC3339Nano, "2016-11-29T16:44:27.192070480Z")
	require.NoError(t, err)
	start := now.Add(-period)
	startTs := start.UnixNano()
	require.NoError(t, w.UpdateWithTime(now))
	repo := "..."
	require.Equal(t, startTs, w.Start(repo).UnixNano())

	require.False(t, w.TestTime(repo, time.Unix(0, 0)))
	require.False(t, w.TestTime(repo, time.Time{}))
	require.True(t, w.TestTime(repo, time.Now()))
	require.True(t, w.TestTime(repo, time.Unix(0, startTs))) // Inclusive.
	require.True(t, w.TestTime(repo, time.Unix(0, startTs+1)))
	require.False(t, w.TestTime(repo, time.Unix(0, startTs-1)))
}

// setupRepo initializes a temporary Git repo with the given number of commits.
// Returns the repo URL, a repograph.Graph instance, a slice of commit hashes,
// and a cleanup function.
func setupRepo(t *testing.T, numCommits int) (string, *repograph.Graph, []string, func()) {
	ctx := context.Background()
	gb := git_testutils.GitInit(t, ctx)
	f := "somefile"
	commits := make([]string, 0, numCommits)
	t0, err := time.Parse(time.RFC3339Nano, "2016-11-29T16:44:27.192070480Z")
	require.NoError(t, err)
	for i := 0; i < numCommits; i++ {
		gb.AddGen(ctx, f)
		ts := t0.Add(time.Duration(int64(5) * int64(i) * int64(time.Second)))
		h := gb.CommitMsgAt(ctx, fmt.Sprintf("C%d", i), ts)
		commits = append(commits, h)
	}

	tmp, err := ioutil.TempDir("", "")
	require.NoError(t, err)
	repo, err := repograph.NewLocalGraph(ctx, gb.Dir(), tmp)
	require.NoError(t, err)
	require.NoError(t, repo.Update(ctx))
	return gb.Dir(), repo, commits, func() {
		gb.Cleanup()
		testutils.RemoveAll(t, tmp)
	}
}

// setup initializes all of the test inputs, including a temporary git repo and
// a Window instance. Returns the Window instance, a convenience function for
// asserting that the Window returns a particular value for a given commit
// index, and a cleanup function.
func setup(t *testing.T, period time.Duration, numCommits, threshold int) (*Window, func(int, bool), func()) {
	unittest.LargeTest(t)

	repoUrl, repo, commits, cleanup := setupRepo(t, numCommits)
	rm := repograph.Map{
		repoUrl: repo,
	}
	w, err := New(context.Background(), period, threshold, rm)
	require.NoError(t, err)
	now := repo.Get(commits[len(commits)-1]).Timestamp.Add(5 * time.Second)
	require.NoError(t, w.UpdateWithTime(now))

	test := func(idx int, expect bool) {
		actual, err := w.TestCommitHash(repoUrl, commits[idx])
		require.NoError(t, err)
		require.Equal(t, expect, actual)
	}
	return w, test, cleanup
}

// Only test the repo, duration is zero.
func TestWindowRepoOnly(t *testing.T) {
	_, test, cleanup := setup(t, 0, 100, 50)
	defer cleanup()

	test(0, false)
	test(20, false)
	test(49, false)
	test(50, true)
	test(51, true)
	test(55, true)
	test(99, true)
}

// Fewer than N commits in the repo.
func TestWindowFewCommits(t *testing.T) {
	_, test, cleanup := setup(t, 0, 5, 10)
	defer cleanup()

	test(0, true)
	test(1, true)
	test(4, true)
}

// Test both repo and duration.
func TestWindowRepoAndDuration1(t *testing.T) {
	_, test, cleanup := setup(t, 30*time.Second, 20, 10)
	defer cleanup()

	// Commits are 5 seconds apart, so the last 6 commits are within 30
	// seconds. In this case the repo will win out and the last 10 commits
	// (index 10-19) will be in range.
	test(0, false)
	test(9, false)
	test(10, true)
	test(11, true)
	test(19, true)
}

func TestWindowRepoAndDuration2(t *testing.T) {
	_, test, cleanup := setup(t, 60*time.Second, 20, 10)
	defer cleanup()

	// Commits are 5 seconds apart, so the last 12 commits are within 60
	// seconds. In this case the time period will win out and the last 11
	// commits (index 8-19) will be in range.
	test(0, false)
	test(7, false)
	test(8, true)
	test(19, true)
}

// Test multiple repos.
func TestWindowMultiRepo(t *testing.T) {
	unittest.LargeTest(t)
	url1, repo1, commits1, cleanup1 := setupRepo(t, 20)
	defer cleanup1()
	url2, repo2, commits2, cleanup2 := setupRepo(t, 10)
	defer cleanup2()

	rm := repograph.Map{
		url1: repo1,
		url2: repo2,
	}
	w, err := New(context.Background(), 0, 6, rm)
	require.NoError(t, err)
	now := repo1.Get(commits1[len(commits1)-1]).Timestamp.Add(5 * time.Second)
	require.NoError(t, w.UpdateWithTime(now))

	test := func(repoUrl, commit string, expect bool) {
		actual, err := w.TestCommitHash(repoUrl, commit)
		require.NoError(t, err)
		require.Equal(t, expect, actual)
	}

	// The last 6 commits of each repo should be in the Window.
	test(url1, commits1[0], false)
	test(url1, commits1[13], false)
	test(url1, commits1[14], true)
	test(url1, commits1[19], true)

	test(url2, commits2[0], false)
	test(url2, commits2[3], false)
	test(url2, commits2[4], true)
	test(url2, commits2[9], true)
}
