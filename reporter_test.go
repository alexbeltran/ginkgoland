package ginkgoland_test

import (
	"bytes"
	"github.com/alexbeltran/ginkgoland"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	"github.com/onsi/ginkgo/types"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Reporter", func() {
	var (
		reporter ginkgoland.Reporter
		buff     *bytes.Buffer
	)
	BeforeEach(func() {
		reporter = ginkgoland.Reporter{}
		buff = &bytes.Buffer{}
		reporter.SetWriter(buff)
	})
	Context("running a test", func() {
		BeforeEach(func() {
			reporter.SpecWillRun(&types.SpecSummary{
				ComponentTexts: []string{"ROOT", "success layer", "should be success"},
			})
		})
		It("should succeed", func() {
			Expect(buff.String()).To(ContainSubstring("=== RUN"))
		})

		table.DescribeTable("running a spec",
			func(state types.SpecState, expectedSubstr string) {
				spec := &types.SpecSummary{
					ComponentTexts: []string{"ROOT", "success layer", "should be success"},
					ComponentCodeLocations: []types.CodeLocation{
						{
							FileName:       "fake.go",
							LineNumber:     30,
							FullStackTrace: "I am a stack trace!",
						},
					},
					RunTime: time.Millisecond * 500,
					State:   state,
				}
				reporter.SpecDidComplete(spec)
				Expect(buff.String()).To(ContainSubstring(expectedSubstr))
			},
			table.Entry("Pass", types.SpecStatePassed, "PASS"),
			table.Entry("Failed", types.SpecStateFailed, "FAIL"),
			table.Entry("Panic", types.SpecStatePanicked, "FAIL"),
			table.Entry("Time out", types.SpecStateTimedOut, "FAIL"),
			table.Entry("Invalid", types.SpecStateInvalid, "FAIL"),
			table.Entry("Skipped", types.SpecStateSkipped, "SKIP"),
		)
	})
})
