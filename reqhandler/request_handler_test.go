package reqhandler

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type RequestHandlerSuite struct{}

var _ = Suite(&RequestHandlerSuite{})

func (s *RequestHandlerSuite) TestHandlesConnection(c *C) {
	output := HandleConnection(new(mockConn))
	c.Assert(true, Equals, output)
}
