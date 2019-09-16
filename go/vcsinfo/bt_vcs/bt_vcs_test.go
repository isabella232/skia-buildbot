package bt_vcs

import (
	"context"
	"io/ioutil"
	"math"
	"testing"
	"time"

	assert "github.com/stretchr/testify/require"
	"go.skia.org/infra/go/gitiles"
	"go.skia.org/infra/go/gitstore"
	"go.skia.org/infra/go/gitstore/mocks"
	gs_testutils "go.skia.org/infra/go/gitstore/testutils"
	"go.skia.org/infra/go/testutils"
	"go.skia.org/infra/go/testutils/unittest"
	"go.skia.org/infra/go/util"
	"go.skia.org/infra/go/vcsinfo"
	vcs_testutils "go.skia.org/infra/go/vcsinfo/testutils"
	"golang.org/x/sync/errgroup"
)

const (
	skiaRepoURL  = "https://skia.googlesource.com/skia.git"
	localRepoURL = "https://example.com/local.git"
)

func TestVCSSuite(t *testing.T) {
	unittest.LargeTest(t)
	vcs, _, cleanup := setupVCSLocalRepo(t, "master")
	defer cleanup()

	// Run the VCS test suite.
	vcs_testutils.TestByIndex(t, vcs)
	vcs_testutils.TestDisplay(t, vcs)
	vcs_testutils.TestFrom(t, vcs)
	vcs_testutils.TestIndexOf(t, vcs)
	vcs_testutils.TestLastNIndex(t, vcs)
	vcs_testutils.TestRange(t, vcs)
}

func TestBranchInfo(t *testing.T) {
	unittest.LargeTest(t)
	vcs, gitStore, cleanup := setupVCSLocalRepo(t, gitstore.ALL_BRANCHES)
	defer cleanup()

	branchPointers, err := gitStore.GetBranches(context.Background())
	assert.NoError(t, err)
	branches := []string{}
	for branchName := range branchPointers {
		if branchName != gitstore.ALL_BRANCHES {
			branches = append(branches, branchName)
		}
	}

	vcs_testutils.TestBranchInfo(t, vcs, branches)
}

// TestConcurrentUpdate verifies that BigTableVCS.Update() behaves correctly
// when called concurrently.
func TestConcurrentUpdate(t *testing.T) {
	unittest.LargeTest(t)

	numGoroutines := 10
	mg := &mocks.GitStore{}
	defer mg.AssertExpectations(t)

	// Pretend the repo has one commit at first.
	ics := makeTestIndexCommits()
	lcs := makeTestLongCommits()
	hashes := make([]string, 0, len(ics))
	for _, ic := range ics {
		hashes = append(hashes, ic.Hash)
	}

	mg.On("RangeN", testutils.AnyContext, 0, math.MaxInt32, "master").Return(ics[:1], nil).Once()
	mg.On("Get", testutils.AnyContext, hashes[:1]).Return(lcs[:1], nil).Once()

	ctx := context.Background()
	vcs, err := New(ctx, mg, "master", nil)
	assert.NoError(t, err)

	// Now, pretend that the other two commits have landed, and run Update
	// in several goroutines. We expect the first call to Update() to run
	// RangeN starting at zero (despite our already having the first commit
	// from above, because we request a range which overlaps by one commit.
	// Subsequent calls should already be loaded with IndexCommits and thus
	// should start at 2.
	mg.On("RangeN", testutils.AnyContext, 0, math.MaxInt32, "master").Return(ics, nil).Once()
	mg.On("Get", testutils.AnyContext, hashes[1:]).Return(lcs[1:], nil).Once()
	mg.On("RangeN", testutils.AnyContext, 2, math.MaxInt32, "master").Return(ics[2:], nil)

	var egroup errgroup.Group
	for i := 0; i < numGoroutines; i++ {
		egroup.Go(func() error {
			return vcs.Update(ctx, true, false)
		})
	}
	assert.NoError(t, egroup.Wait())
}

// TestGetFile makes sure that we can use gittiles to fetch an
// arbitrary file (DEPS) from the Skia repo at a chosen commit.
func TestGetFile(t *testing.T) {
	unittest.LargeTest(t)
	gtRepo := gitiles.NewRepo(skiaRepoURL, nil)
	hash := "9be246ed747fd1b900013dd0596aed0b1a63a1fa"
	vcs := &BigTableVCS{
		gitiles: gtRepo,
	}
	_, err := vcs.GetFile(context.Background(), "DEPS", hash)
	assert.NoError(t, err)
}

// TestDetailsCaching makes sure that multiple calls to Details do
// not result in multiple calls to the underlying gitstore, that is,
// the details per commit hash are cached.
func TestDetailsCaching(t *testing.T) {
	unittest.SmallTest(t)

	mg := &mocks.GitStore{}
	defer mg.AssertExpectations(t)

	commits := makeTestLongCommits()

	mg.On("RangeN", testutils.AnyContext, 0, math.MaxInt32, "master").Return(makeTestIndexCommits(), nil)
	mg.On("Get", testutils.AnyContext, []string{firstHash, secondHash, thirdHash}).Return(commits, nil).Once()

	vcs, err := New(context.Background(), mg, "master", nil)
	assert.NoError(t, err)

	// query details 3 times, and make sure it uses the cache after the
	// first time. Since we said Once() on the mocked Get function, we are
	// assured that gitstore.Get() is only called once.
	ctx := context.Background()
	c, err := vcs.Details(ctx, firstHash, false)
	assert.NoError(t, err)
	assert.Equal(t, commits[0], c)
	assert.Nil(t, c.Branches)
	c, err = vcs.Details(ctx, firstHash, false)
	assert.NoError(t, err)
	assert.Equal(t, commits[0], c)
	c, err = vcs.Details(ctx, firstHash, false)
	assert.NoError(t, err)
	assert.Equal(t, commits[0], c)
}

// TestDetailsMultiCaching makes sure that multiple calls to DetailsMulti do
// not result in multiple calls to the underlying gitstore, that is,
// the details per commit hash are cached.
func TestDetailsMultiCaching(t *testing.T) {
	unittest.SmallTest(t)

	mg := &mocks.GitStore{}
	defer mg.AssertExpectations(t)

	commits := makeTestLongCommits()

	mg.On("RangeN", testutils.AnyContext, 0, math.MaxInt32, "master").Return(makeTestIndexCommits(), nil)
	mg.On("Get", testutils.AnyContext, []string{firstHash, secondHash, thirdHash}).Return(commits, nil).Once()

	vcs, err := New(context.Background(), mg, "master", nil)
	assert.NoError(t, err)

	// query details 3 times, and make sure it uses the cache after the
	// first time. Since we said Once() on the mocked Get function, we are
	// assured that gitstore.Get() is only called once.
	ctx := context.Background()
	c, err := vcs.DetailsMulti(ctx, []string{firstHash, secondHash}, false)
	assert.NoError(t, err)
	assert.NotNil(t, c)
	assert.Len(t, c, 2)
	assert.Equal(t, commits[0], c[0])
	assert.Equal(t, commits[1], c[1])
	c, err = vcs.DetailsMulti(ctx, []string{firstHash, secondHash}, false)
	assert.NoError(t, err)
	assert.NotNil(t, c)
	assert.Len(t, c, 2)
	assert.Equal(t, commits[0], c[0])
	assert.Equal(t, commits[1], c[1])
	c, err = vcs.DetailsMulti(ctx, []string{firstHash, secondHash}, false)
	assert.NoError(t, err)
	assert.NotNil(t, c)
	assert.Len(t, c, 2)
	assert.Equal(t, commits[0], c[0])
	assert.Equal(t, commits[1], c[1])
}

// setupVCSLocalRepo loads the test repo into a new GitStore and returns an instance of vcsinfo.VCS.
func setupVCSLocalRepo(t *testing.T, branch string) (vcsinfo.VCS, gitstore.GitStore, func()) {
	repoDir, cleanup := vcs_testutils.InitTempRepo()
	wd, err := ioutil.TempDir("", "")
	assert.NoError(t, err)
	ctx := context.Background()
	_, _, btgs := gs_testutils.SetupAndLoadBTGitStore(t, ctx, wd, "file://"+repoDir, true)
	vcs, err := New(ctx, btgs, branch, nil)
	assert.NoError(t, err)
	return vcs, btgs, func() {
		util.RemoveAll(wd)
		cleanup()
	}
}

const (
	// arbitrary sha1 hashes
	firstHash  = "ae76331b95dfc399cd776d2fc68021e0db03cc4f"
	secondHash = "b295e0bdde1938d1fbfd343e5a3e569e868e1465"
	thirdHash  = "cf70f4c33de2200b76651bbe1e54aa55fcd77447"
)

var (
	// arbitrary times
	firstTime  = time.Date(2019, time.May, 2, 12, 0, 3, 0, time.UTC)
	secondTime = time.Date(2019, time.May, 2, 14, 1, 3, 0, time.UTC)
	thirdTime  = time.Date(2019, time.May, 2, 17, 5, 3, 0, time.UTC)
)

// This test data (for a repo of 3 commits) is returned via functions
// to make it convenient to have a copy of the data for each test,
// so the tests can write all over the returned values w/o impacting
// tests that follow.
func makeTestLongCommits() []*vcsinfo.LongCommit {
	return []*vcsinfo.LongCommit{
		{
			ShortCommit: &vcsinfo.ShortCommit{
				Author:  "alpha@example.com",
				Hash:    firstHash,
				Subject: "initial commit",
			},
			Body:      "awesome message",
			Parents:   []string{},
			Timestamp: firstTime,
		},
		{
			ShortCommit: &vcsinfo.ShortCommit{
				Author:  "beta@example.com",
				Hash:    secondHash,
				Subject: "followup commit",
			},
			Body:      "bug fixes",
			Parents:   []string{firstHash},
			Timestamp: secondTime,
		},
		{
			ShortCommit: &vcsinfo.ShortCommit{
				Author:  "gamma@example.com",
				Hash:    thirdHash,
				Subject: "last commit",
			},
			Body:      "now deprecated",
			Parents:   []string{secondHash},
			Timestamp: thirdTime,
		},
	}
}

func makeTestIndexCommits() []*vcsinfo.IndexCommit {
	return []*vcsinfo.IndexCommit{
		{
			Hash:      firstHash,
			Index:     0,
			Timestamp: firstTime,
		},
		{
			Hash:      secondHash,
			Index:     1,
			Timestamp: secondTime,
		},
		{
			Hash:      thirdHash,
			Index:     2,
			Timestamp: thirdTime,
		},
	}
}

func makeTestBranchPointerMap() map[string]*gitstore.BranchPointer {
	return map[string]*gitstore.BranchPointer{
		"master": {
			Head:  "master",
			Index: 3,
		},
	}
}
