package ginkgoland

import (
	"fmt"
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/types"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

type Reporter struct {
	T         *testing.T
	suiteName string
	out       io.Writer
}

func (r *Reporter) SetWriter(w io.Writer) {
	r.out = w
}

func (r *Reporter) setDefaultWriter() {
	if r.out == nil {
		r.out = os.Stdout
	}
}

func (r *Reporter) SpecSuiteWillBegin(config config.GinkgoConfigType, summary *types.SuiteSummary) {
	r.setDefaultWriter()
	if r.T == nil {
		r.suiteName = "GinkgoTest"
	} else {
		r.suiteName = r.T.Name()
	}
}

func (r *Reporter) BeforeSuiteDidRun(setupSummary *types.SetupSummary) {
	r.setDefaultWriter()
}

func (r *Reporter) SpecWillRun(specSummary *types.SpecSummary) {
	r.setDefaultWriter()
	r.printGoRun(nameFromSummary(specSummary))
}

func (r *Reporter) SpecDidComplete(specSummary *types.SpecSummary) {
	r.setDefaultWriter()

	var state string
	switch specSummary.State {
	case types.SpecStatePassed:
		state = "PASS"
	case types.SpecStateFailed, types.SpecStateInvalid, types.SpecStateTimedOut, types.SpecStatePanicked:
		state = "FAIL"
		r.printGoResult(state, nameFromSummary(specSummary), specSummary.RunTime)

		// print where and why we failed
		_, _ = fmt.Fprintf(r.out, "  %s: %s\n", specSummary.Failure.Location.String(), specSummary.Failure.Message)
		return
	case types.SpecStateSkipped:
		state = "SKIP"
		r.printGoResult(state, nameFromSummary(specSummary), specSummary.RunTime)

		// print what was skipped
		_, _ = fmt.Fprintf(r.out, "  %+v is skipped\n", specSummary.ComponentCodeLocations[len(specSummary.ComponentCodeLocations)-1])
		return
	default:
		return
	}
	r.printGoResult(state, nameFromSummary(specSummary), specSummary.RunTime)
}

func (r *Reporter) AfterSuiteDidRun(setupSummary *types.SetupSummary) {
}

func (r *Reporter) SpecSuiteDidEnd(summary *types.SuiteSummary) {
	r.setDefaultWriter()
	state := "PASS"
	if summary.NumberOfFailedSpecs+summary.NumberOfSkippedSpecs > 0 {
		state = "FAIL"
	}
	_, _ = fmt.Fprintf(r.out, "--- %s: %s (%0.3f)\n", state, r.suiteName, summary.RunTime.Seconds())
}

func nameFromSummary(summary *types.SpecSummary) string {
	return strings.Join(summary.ComponentTexts[1:], "/")
}

func (r *Reporter) printGoResult(state string, name string, duration time.Duration) {
	_, _ = fmt.Fprintf(r.out, "--- %s: %s/%s (%0.3f)\n", state, r.suiteName, name, duration.Seconds())
}

func (r *Reporter) printGoRun(name string) {
	_, _ = fmt.Fprintf(r.out, "=== RUN   %s/%s\n", r.suiteName, name)
}
