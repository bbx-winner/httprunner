package hrp

import (
	"testing"
	"time"

	"github.com/rs/zerolog"

	"github.com/httprunner/httprunner/v4/hrp/pkg/boomer"
)

func TestMyBoomerStandaloneRun(t *testing.T) {
	// buildHashicorpPyPlugin()
	// defer removeHashicorpPyPlugin()
	path := "demo-flat/testcases/requests.json"
	testcase2 := TestCasePath(path)
	profile := &boomer.Profile{}
	profile.SpawnCount = 2
	profile.SpawnRate = 1
	zerolog.SetGlobalLevel(zerolog.WarnLevel)
	b := NewStandaloneBoomer(profile.SpawnCount, profile.SpawnRate)
	b.SetProfile(profile)
	b.InitBoomer()
	go b.Run(&testcase2)
	time.Sleep(1 * 20 * time.Second)
	b.Quit()
	time.Sleep(10 * time.Second)
}
