package config

import (
	"testing"

	"github.com/stretchr/testify/require"

	coreconfig "github.com/hadenlabs/terraform-supabase/config"
)

func TestConfigLoadEnvSuccess(t *testing.T) {
	conf, err := LoadEnvWithFilename("./mocking/config.env")
	require.NoError(t, err, "unexpected error: %v", err)
	require.IsType(t, &coreconfig.Config{}, conf)
	require.Equal(t, "zap", conf.Log.Provider, "unexpected log provider")
}

func TestConfigMustLoadEnvWithPanic(t *testing.T) {
	require.Panics(
		t,
		func() { MustLoadEnvWithFilename("./mocking/notfound.env") },
		"The code did not panic as expected",
	)
}

func TestConfigLoadEnvFailed(t *testing.T) {
	conf, err := LoadEnvWithFilename("./mocking/notfound.env")
	require.Error(t, err, "expected an error but got none")
	require.Nil(t, conf, "expected config to be nil when loading fails")
}
