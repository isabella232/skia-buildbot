// Package alertstest contains test utils for the alerts package.
package alertstest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.skia.org/infra/perf/go/alerts"
)

// Store_SaveListDelete tests that an alerts.Store instance operates as expected.
func Store_SaveListDelete(t *testing.T, a alerts.Store) {
	ctx := context.Background()

	// TODO(jcgregorio) Break up into finer grained tests.
	cfg := alerts.NewConfig()
	cfg.Query = "source_type=svg"
	cfg.DisplayName = "bar"
	err := a.Save(ctx, cfg)
	assert.NoError(t, err)

	// Confirm it appears in the list.
	cfgs, err := a.List(ctx, false)
	assert.NoError(t, err)
	assert.Len(t, cfgs, 1)

	// Delete it.
	err = a.Delete(ctx, int(cfgs[0].ID))
	assert.NoError(t, err)

	// Confirm it is still there if we list deleted configs.
	cfgs, err = a.List(ctx, true)
	assert.NoError(t, err)
	assert.Len(t, cfgs, 1)

	// Confirm it is not there if we don't list deleted configs.
	cfgs, err = a.List(ctx, false)
	assert.NoError(t, err)
	assert.Len(t, cfgs, 0)

	// Store a second config.
	cfg = alerts.NewConfig()
	cfg.Query = "source_type=skp"
	cfg.DisplayName = "foo"
	err = a.Save(ctx, cfg)
	assert.NoError(t, err)

	// Confirm they are both listed when including deleted configs, and they are
	// ordered by DisplayName.
	cfgs, err = a.List(ctx, true)
	assert.NoError(t, err)
	assert.Len(t, cfgs, 2)
	assert.Equal(t, "bar", cfgs[0].DisplayName)
	assert.Equal(t, "foo", cfgs[1].DisplayName)
}
