package junitxml

import (
	"encoding/xml"
	"errors"
	"io"
	"regexp"
)

// Parse reads XML-encoded data from the given reader and decodes it into
// a [TestReport] struct.
func Parse(r io.Reader) (TestReport, error) {
	var report TestReport

	if err := xml.NewDecoder(r).Decode(&report); err != nil {
		return TestReport{}, err
	}
	return report, nil
}

// ParseMany works like [Parse] but supports reading multiple <testsuites>
// root elements and returns a [TestReport] for each of them.
func ParseMany(r io.Reader) ([]TestReport, error) {
	var (
		reports []TestReport
		suites  []TestSuite
	)

	decoder := xml.NewDecoder(r)
	for {
		tok, err := decoder.Token()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, err
		}

		if elem, ok := tok.(xml.StartElement); ok {
			var report TestReport
			if err := decoder.DecodeElement(&report, &elem); err == nil {
				reports = append(reports, report)
			} else if !errors.Is(err, io.EOF) {
				var suite TestSuite
				if err := decoder.DecodeElement(&suite, &elem); err != nil {
					return nil, err
				}
				suites = append(suites, suite)
			} else {
				return nil, err
			}
		}
	}

	if len(suites) > 0 {
		reports = append(reports, newTestReport(suites))
	}

	return reports, nil
}

const (
	AttachmentRegex = `\[\[ATTACHMENT\|(?<path>[^\[\]\|]+?)\]\]`
	PropertyRegex   = `\[\[PROPERTY\|(?<name>[^\[\]\|=]+)=(?<value>[^\[\]\|=]+)\]\]`

	// PropertyTextValueRegex = `\[\[PROPERTY\|(?<name>[^\[\]\|=]+)\]\]\n(?<text>(?:.*\n)+)\[\[/PROPERTY\]\]`
)

func ParseTextAttachments(s string) []string {
	var paths []string

	re := regexp.MustCompile(AttachmentRegex)
	for _, match := range re.FindAllStringSubmatch(s, -1) {
		paths = append(paths, match[1])
	}

	return paths
}

func ParseTextProperties(s string) []Property {
	var properties []Property

	re := regexp.MustCompile(PropertyRegex)
	for _, match := range re.FindAllStringSubmatch(s, -1) {
		properties = append(properties, Property{
			Name:  match[1],
			Value: match[2],
		})
	}

	return properties
}
