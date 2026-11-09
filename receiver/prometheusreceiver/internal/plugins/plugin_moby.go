// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !remove_all_sd || enable_moby_sd

package plugins // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver/internal/plugins"

import (
	_ "github.com/prometheus/prometheus/discovery/moby" // Register moby plugin.
)
