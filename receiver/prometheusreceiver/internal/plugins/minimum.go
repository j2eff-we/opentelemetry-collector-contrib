// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package plugins registers the Prometheus service discovery mechanisms that
// the receiver supports. It mirrors github.com/prometheus/prometheus/plugins,
// including its build tags, but leaves out Outscale SD because
// github.com/outscale/osc-sdk-go/v2 depends on github.com/aws/aws-sdk-go (v1).
package plugins // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver/internal/plugins"

import (
	_ "github.com/prometheus/prometheus/discovery/file" // Register file plugin.
	_ "github.com/prometheus/prometheus/discovery/http" // Register http plugin.
)
