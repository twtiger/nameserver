package nameserver

import . "gopkg.in/check.v1"

type DatabaseSuite struct{}

var _ = Suite(&DatabaseSuite{})

func (s *DatabaseSuite) Test_RetrievalOfTwoARecordsForOurAuthoritativeDomain(c *C) {
	labels := createLabelsFor("twtiger.com.")

	records := retrieve(labels)

	c.Assert(len(records), Equals, 2)
	c.Assert(records, DeepEquals,
		[]*record{
			tigerRecord1,
			tigerRecord2,
		})
}

func (s *DatabaseSuite) Test_ReturnsEmptyRecordSliceIfDatabaseDoesNotContainDomain(c *C) {
	labels := createLabelsFor("www.nothere.orthere.")

	records := retrieve(labels)

	c.Assert(len(records), Equals, 0)
}
