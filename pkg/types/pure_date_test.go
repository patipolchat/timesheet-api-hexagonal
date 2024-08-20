package types

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type PureDateSuite struct {
	suite.Suite
}

func (s *PureDateSuite) TestString() {
	d := PureDate{Year: 2021, Month: 1, Day: 1}
	s.Equal("2021-01-01", d.String())
}

func (s *PureDateSuite) TestMarshalJSON() {
	d := PureDate{Year: 2021, Month: 1, Day: 1}
	b, err := d.MarshalJSON()
	s.NoError(err)
	s.Equal(`"2021-01-01"`, string(b))
}

func (s *PureDateSuite) TestUnmarshalJSON() {
	d := PureDate{}
	err := d.UnmarshalJSON([]byte(`"2021-01-01"`))
	s.Require().NoError(err)
	s.Equal(2021, d.Year)
	s.Equal(1, d.Month)
	s.Equal(1, d.Day)
}

func (s *PureDateSuite) TestUnmarshalJSON_Invalid() {
	d := PureDate{}
	err := d.UnmarshalJSON([]byte(`"2021-01-32"`))
	s.Error(err)
}

func TestPureDateSuite(t *testing.T) {
	suite.Run(t, new(PureDateSuite))
}
