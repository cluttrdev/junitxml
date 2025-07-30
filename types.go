package junitxml

import (
	"encoding/xml"
)

type TestReport struct {
	XMLName    xml.Name    `xml:"testsuites"`
	Tests      int         `xml:"tests,attr,omitempty"`
	Failures   int         `xml:"failures,attr,omitempty"`
	Errors     int         `xml:"errors,attr,omitempty"`
	Skipped    int         `xml:"skipped,attr,omitempty"`
	Time       float64     `xml:"time,attr,omitempty"`
	Timestamp  string      `xml:"timestamp,attr,omitempty"`
	TestSuites []TestSuite `xml:"testsuite"`
}

type TestSuite struct {
	// XMLName    xml.Name   `xml:"testsuite"`
	Name       string     `xml:"name,attr,omitempty"`
	Tests      int        `xml:"tests,attr,omitempty"`
	Failures   int        `xml:"failures,attr,omitempty"`
	Errors     int        `xml:"errors,attr,omitempty"`
	Skipped    int        `xml:"skipped,attr,omitempty"`
	Time       float64    `xml:"time,attr,omitempty"`
	Timestamp  string     `xml:"timestamp,attr,omitempty"`
	File       string     `xml:"file,attr,omitempty"`
	Properties []Property `xml:"properties>property,omitempty"`
	SystemOut  *SystemOut `xml:"system-out"`
	SystemErr  *SystemErr `xml:"system-err"`
	TestCases  []TestCase `xml:"testcase"`
}

type TestCase struct {
	// XMLName   xml.Name  `xml:"testcase"`
	Name       string     `xml:"name,attr,omitempty"`
	Classname  string     `xml:"classname,attr,omitempty"`
	Tests      int        `xml:"tests,attr,omitempty"`
	Time       float64    `xml:"time,attr,omitempty"`
	File       string     `xml:"file,attr,omitempty"`
	Line       int        `xml:"line,attr,omitempty"`
	Failure    *Failure   `xml:"failure"`
	Error      *Error     `xml:"error"`
	Skipped    *Skipped   `xml:"skipped"`
	Properties []Property `xml:"properties>property,omitempty"`
	SystemOut  *SystemOut `xml:"system-out"`
	SystemErr  *SystemErr `xml:"system-err"`
}

type Failure struct {
	// XMLName xml.Name `xml:"failure"`
	Message string `xml:"message,attr,omitempty"`
	Type    string `xml:"type,attr,omitempty"`
	Text    string `xml:",innerxml"`
}

type Error struct {
	// XMLName xml.Name `xml:"error"`
	Message string `xml:"message,attr,omitempty"`
	Type    string `xml:"type,attr,omitempty"`
	Text    string `xml:",innerxml"`
}

type Skipped struct {
	// XMLName xml.Name `xml:"skipped"`
	Message string `xml:"message,attr,omitempty"`
}

type Property struct {
	// XMLName xml.Name `xml:"property"`
	Name  string `xml:"name,attr,omitempty"`
	Value string `xml:"value,attr,omitempty"`
	Text  string `xml:",innerxml"`
}

type SystemOut struct {
	// XMLName xml.Name `xml:"system-out"`
	Text string `xml:",innerxml"`
}

type SystemErr struct {
	// XMLName xml.Name `xml:"system-err"`
	Text string `xml:",innerxml"`
}

func newTestReport(suites []TestSuite) TestReport {
	report := TestReport{
		TestSuites: suites,
	}
	report.XMLName.Local = "testsuites"

	for i := 0; i < len(report.TestSuites); i++ {
		report.Tests += report.TestSuites[i].Tests
		report.Failures += report.TestSuites[i].Failures
		report.Errors += report.TestSuites[i].Errors
		report.Skipped += report.TestSuites[i].Skipped
		report.Time += report.TestSuites[i].Time
	}

	return report
}
