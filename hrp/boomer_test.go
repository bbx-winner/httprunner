package hrp

import (
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/httprunner/httprunner/v4/hrp/pkg/boomer"
)

func TestBoomerStandaloneRun(t *testing.T) {
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.WarnLevel)
	buildHashicorpGoPlugin()
	defer removeHashicorpGoPlugin()
	testcase1 := &TestCase{
		Config: NewConfig("TestCase1").SetBaseURL("http://httpbin.org"),
		TestSteps: []IStep{
			NewStep("headers").
				GET("/headers").
				Validate().
				AssertEqual("status_code", 200, "check status code").
				AssertEqual("headers.\"Content-Type\"", "application/json", "check http response Content-Type"),
			NewStep("user-agent").
				GET("/user-agent").
				Validate().
				AssertEqual("status_code", 200, "check status code").
				AssertEqual("headers.\"Content-Type\"", "application/json", "check http response Content-Type"),
			NewStep("TestCase3").CallRefCase(&TestCase{Config: NewConfig("TestCase3")}),
		},
	}
	testcase2 := TestCasePath(demoTestCaseWithPluginJSONPath)

	b := NewStandaloneBoomer(10, 10)
	b.AddOutput(boomer.NewConsoleOutput())
	go b.Run(testcase1, &testcase2)
	time.Sleep(60 * 60 * time.Second)
	b.Quit()
}
