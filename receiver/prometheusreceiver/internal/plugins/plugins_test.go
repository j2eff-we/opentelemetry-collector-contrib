// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !remove_all_sd

package plugins

import (
	"testing"

	"github.com/prometheus/common/promslog"
	promconfig "github.com/prometheus/prometheus/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServiceDiscoveryRegistration(t *testing.T) {
	tests := []struct {
		name    string
		sd      string
		wantErr string
	}{
		{name: "file", sd: `file_sd_configs: [{files: [targets.json]}]`},
		{name: "kubernetes", sd: `kubernetes_sd_configs: [{role: pod}]`},
		{name: "ec2", sd: `ec2_sd_configs: [{region: us-east-1}]`},
		{name: "outscale is not registered", sd: `outscale_sd_configs: [{}]`, wantErr: "field outscale_sd_configs not found"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := "scrape_configs:\n  - job_name: test\n    " + tt.sd + "\n"
			_, err := promconfig.Load(cfg, promslog.NewNopLogger())
			if tt.wantErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
		})
	}
}
