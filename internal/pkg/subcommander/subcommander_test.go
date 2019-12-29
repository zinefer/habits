package subcommander_test

import (
	"bytes"
	"flag"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/zinefer/habits/internal/pkg/subcommander"
)

type TestSuite struct {
	suite.Suite
}

var update = flag.Bool("update", false, "update .golden files")

var fakeSubcommandCalls []string

func addFakeSubcommandCall(name string) bool {
	fakeSubcommandCalls = append(fakeSubcommandCalls, name)
	return true
}

func (suite *TestSuite) TestSubcommander() {
	sc := subcommander.New()
	sc.Register("a", "a command", &FakeSubcommandA{})
	sc.Register("b", "b command", &FakeSubcommandB{})
	sc.Register("c", "c command", &FakeSubcommandC{})

	sc.Execute("a")
	sc.Execute("a:b")
	sc.Execute("c:a")
	sc.Execute("b")
	sc.Execute("c:b:b")
	sc.Execute("c:b")

	assert.Equal(suite.T(), []string{"A", "AB", "CA", "B", "CBB", "CB"}, fakeSubcommandCalls)

	actualOutput := captureOutput(func() {
		sc.Execute("")
	})

	expectedOutput := suite.getGoldenfile(actualOutput)

	assert.Equal(suite.T(), expectedOutput, actualOutput)
}

func TestSubcommanderTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) getGoldenfile(data string) string {
	gp := "testdata.golden"
	if *update {
		suite.T().Log("update golden file")
		if err := ioutil.WriteFile(gp, []byte(data), 0644); err != nil {
			suite.T().Fatalf("failed to update golden file: %s", err)
		}
	}
	g, err := ioutil.ReadFile(gp)
	if err != nil {
		suite.T().Fatalf("failed reading .golden: %s", err)
	}
	return string(g)
}

func captureOutput(f func()) string {
	orig := os.Stdout // Capture old stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// copy the output in a separate goroutine so printing can't block indefinitely
	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	f()

	w.Close()
	os.Stdout = orig
	return <-outC
}

// Mock subcommand structs below
type FakeSubcommandA struct{}
type FakeSubcommandAA struct{}
type FakeSubcommandAB struct{}

func (*FakeSubcommandA) Subcommander() *subcommander.Subcommander {
	sc := subcommander.New()
	sc.Register("a", "aanother command", &FakeSubcommandAA{})
	sc.Register("b", "abnother command", &FakeSubcommandAB{})
	return sc
}

func (*FakeSubcommandA) Run() bool {
	return addFakeSubcommandCall("A")
}

func (*FakeSubcommandAA) Subcommander() *subcommander.Subcommander {
	return subcommander.New()
}

func (*FakeSubcommandAA) Run() bool {
	return addFakeSubcommandCall("AA")
}

func (*FakeSubcommandAB) Subcommander() *subcommander.Subcommander {
	return subcommander.New()
}

func (*FakeSubcommandAB) Run() bool {
	return addFakeSubcommandCall("AB")
}

type FakeSubcommandB struct{}

func (*FakeSubcommandB) Subcommander() *subcommander.Subcommander {
	return subcommander.New()
}

func (*FakeSubcommandB) Run() bool {
	return addFakeSubcommandCall("B")
}

type FakeSubcommandC struct{}
type FakeSubcommandCA struct{}
type FakeSubcommandCB struct{}
type FakeSubcommandCBA struct{}
type FakeSubcommandCBB struct{}

func (*FakeSubcommandC) Subcommander() *subcommander.Subcommander {
	sc := subcommander.New()
	sc.Register("a", "ACommand", &FakeSubcommandCA{})
	sc.Register("b", "Already chewed command", &FakeSubcommandCB{})
	return sc
}

func (*FakeSubcommandC) Run() bool {
	return addFakeSubcommandCall("C")
}

func (*FakeSubcommandCA) Subcommander() *subcommander.Subcommander {
	return subcommander.New()
}

func (*FakeSubcommandCA) Run() bool {
	return addFakeSubcommandCall("CA")
}

func (*FakeSubcommandCB) Subcommander() *subcommander.Subcommander {
	sc := subcommander.New()
	sc.Register("a", "CBA Command", &FakeSubcommandCBA{})
	sc.Register("b", "CBB command", &FakeSubcommandCBB{})
	return sc
}

func (*FakeSubcommandCB) Run() bool {
	return addFakeSubcommandCall("CB")
}

func (*FakeSubcommandCBA) Subcommander() *subcommander.Subcommander {
	return subcommander.New()
}

func (*FakeSubcommandCBA) Run() bool {
	return addFakeSubcommandCall("CBA")
}

func (*FakeSubcommandCBB) Subcommander() *subcommander.Subcommander {
	return subcommander.New()
}

func (*FakeSubcommandCBB) Run() bool {
	return addFakeSubcommandCall("CBB")
}
