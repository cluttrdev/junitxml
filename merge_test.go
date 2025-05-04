package junitxml_test

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"

	"go.cluttr.dev/junitxml"
)

func TestMerge(t *testing.T) {
	reports := []junitxml.TestReport{
		{
			XMLName: xml.Name{Local: "testsuites"},
			Time:    21.0,
			TestSuites: []junitxml.TestSuite{
				{
					Name: "Tests.Registration",
					Time: 6.0,
					TestCases: []junitxml.TestCase{
						{Name: "testCase1", Classname: "Tests.Registration", Time: 1.0},
						{Name: "testCase2", Classname: "Tests.Registration", Time: 2.0},
						{Name: "testCase3", Classname: "Tests.Registration", Time: 3.0},
					},
				},
				{
					Name: "Tests.Authentication",
					Time: 15,
					TestCases: []junitxml.TestCase{
						{Name: "testCase4", Classname: "Tests.Authentication", Time: 4.0},
						{Name: "testCase5", Classname: "Tests.Authentication", Time: 5.0},
						{Name: "testCase6", Classname: "Tests.Authentication", Time: 6.0, Failure: &junitxml.Failure{
							Message: "Assertion error message", Type: "AssertionError", Text: `Call stack printed here`,
						}},
					},
				},
			},
		},
		{
			XMLName:   xml.Name{Local: "testsuites"},
			Time:      1.0,
			Timestamp: "2006-01-02T15:04:05+0700",
			TestSuites: []junitxml.TestSuite{
				{
					Name: "Tests.Login",
					Time: 1.0,
					TestCases: []junitxml.TestCase{
						{Name: "testCase1", Classname: "Tests.Login", Time: 1.0},
					},
				},
			},
		},
		{
			XMLName:   xml.Name{Local: "testsuites"},
			Time:      1.0,
			Timestamp: "2025-05-04T12:16:04+0000",
			TestSuites: []junitxml.TestSuite{
				{
					Name: "Tests.Logout",
					Time: 1.0,
					TestCases: []junitxml.TestCase{
						{Name: "testCase1", Classname: "Tests.Logout", Time: 1.0},
					},
				},
			},
		},
	}

	report := junitxml.Merge(reports)

	want := junitxml.TestReport{
		XMLName:   xml.Name{Local: "testsuites"},
		Time:      23.0,
		Timestamp: "",
		TestSuites: []junitxml.TestSuite{
			{
				Name: "Tests.Registration",
				Time: 6.0,
				TestCases: []junitxml.TestCase{
					{Name: "testCase1", Classname: "Tests.Registration", Time: 1.0},
					{Name: "testCase2", Classname: "Tests.Registration", Time: 2.0},
					{Name: "testCase3", Classname: "Tests.Registration", Time: 3.0},
				},
			},
			{
				Name: "Tests.Authentication",
				Time: 15,
				TestCases: []junitxml.TestCase{
					{Name: "testCase4", Classname: "Tests.Authentication", Time: 4.0},
					{Name: "testCase5", Classname: "Tests.Authentication", Time: 5.0},
					{Name: "testCase6", Classname: "Tests.Authentication", Time: 6.0, Failure: &junitxml.Failure{
						Message: "Assertion error message", Type: "AssertionError", Text: `Call stack printed here`,
					}},
				},
			},
			{
				Name: "Tests.Login",
				Time: 1.0,
				TestCases: []junitxml.TestCase{
					{Name: "testCase1", Classname: "Tests.Login", Time: 1.0},
				},
			},
			{
				Name: "Tests.Logout",
				Time: 1.0,
				TestCases: []junitxml.TestCase{
					{Name: "testCase1", Classname: "Tests.Logout", Time: 1.0},
				},
			},
		},
	}

	if diff := cmp.Diff(want, report); diff != "" {
		t.Errorf("Mismatch (-want +got):\n%s", diff)
	}
}
