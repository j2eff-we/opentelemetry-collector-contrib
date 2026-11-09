// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !remove_all_sd || enable_stackit_sd

package plugins // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver/internal/plugins"

import (
	_ "github.com/prometheus/prometheus/discovery/stackit" // Register stackit plugin.
)
