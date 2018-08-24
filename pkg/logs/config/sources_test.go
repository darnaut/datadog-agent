// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2018 Datadog, Inc.

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddSource(t *testing.T) {
	sources := NewLogSources()
	assert.Equal(t, 0, len(sources.GetSources()))
	sources.AddSource(NewLogSource("foo", nil, SourceOriginConfig))
	assert.Equal(t, 1, len(sources.GetSources()))
	sources.AddSource(NewLogSource("bar", nil, SourceOriginConfig))
	assert.Equal(t, 2, len(sources.GetSources()))
}

func TestRemoveSource(t *testing.T) {
	sources := NewLogSources()
	source1 := NewLogSource("foo", nil, SourceOriginConfig)
	sources.AddSource(source1)
	source2 := NewLogSource("bar", nil, SourceOriginConfig)
	sources.AddSource(source2)
	assert.Equal(t, 2, len(sources.GetSources()))
	sources.RemoveSource(source1)
	assert.Equal(t, 1, len(sources.GetSources()))
	assert.Equal(t, source2, sources.GetSources()[0])
	sources.RemoveSource(source2)
	assert.Equal(t, 0, len(sources.GetSources()))
}

func TestGetSources(t *testing.T) {
	sources := NewLogSources()
	assert.Equal(t, 0, len(sources.GetSources()))
	sources.AddSource(NewLogSource("", nil, SourceOriginConfig))
	assert.Equal(t, 1, len(sources.GetSources()))
}
