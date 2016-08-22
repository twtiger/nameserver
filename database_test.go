package nameserver

import . "gopkg.in/check.v1"

type DatabaseSuite struct{}

var _ = Suite(&DatabaseSuite{})

func (s *DatabaseSuite) Test_RetrievalOfTwoARecordsForOurAuthoritativeDomain(c *C) {
	labels := []label{"twtiger", "com"}

	records := retrieve(labels)

	c.Assert(len(records), Equals, 2)
	c.Assert(records[0], DeepEquals,
		&record{
			Name:  "twtiger.com.",
			Type:  qtypeA,
			Class: qclassIN,
			TTL:   oneHour,
			RData: "123.123.7.8",
		})
	c.Assert(records[1], DeepEquals,
		&record{
			Name:  "twtiger.com.",
			Type:  qtypeA,
			Class: qclassIN,
			TTL:   oneHour,
			RData: "78.78.90.1",
		})
}

func (s *DatabaseSuite) Test_ReturnsEmptyRecordSliceIfDatabaseDoesNotContainDomain(c *C) {
	labels := []label{"www", "nothere", "orthere"}

	records := retrieve(labels)

	c.Assert(len(records), Equals, 0)
}
