package provider

import (
	"testing"

	"github.com/hadenlabs/terraform-supabase/internal/testutil/config"
)

func zapForTest() (*ZapLog, func()) {
	conf := config.MustLoadEnvWithFilename("./mocking/zap.env")
	log := NewZap(*conf)

	return log, func() {}
}

func TestZapDebugSuccess(t *testing.T) {
	log, tearDown := zapForTest()
	defer tearDown()

	log.Debugf("test subject")
}
