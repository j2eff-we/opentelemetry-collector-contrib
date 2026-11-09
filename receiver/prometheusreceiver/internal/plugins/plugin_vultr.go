// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !remove_all_sd || enable_vultr_sd

package plugins // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver/internal/plugins"

import (
	_ "github.com/prometheus/prometheus/discovery/vultr" // Register vultr plugin.
)
