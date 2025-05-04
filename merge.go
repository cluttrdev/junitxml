package junitxml

// Merge combines multiple [TestReport] items into one.
// The timestamp of final report will be left empty.
func Merge(reports []TestReport) TestReport {
	var report TestReport

	report.XMLName.Local = "testsuites"

	for _, r := range reports {
		report.Tests += r.Tests
		report.Failures += r.Failures
		report.Errors += r.Errors
		report.Skipped += r.Skipped
		report.Time += r.Time

		report.TestSuites = append(report.TestSuites, r.TestSuites...)
	}

	return report
}
