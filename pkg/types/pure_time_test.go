package types

import (
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"testing"
)

type PureTimeStruct struct {
	Time PureTime `json:"time"`
}

type PureTimeSuite struct {
	suite.Suite
}

func (s *PureTimeSuite) TestString() {
	t := PureTime{Hour: 1, Minute: 30}
	s.Equal("01:30", t.String())
}

func (s *PureTimeSuite) TestToMinutes() {
	t := PureTime{Hour: 1, Minute: 30}
	s.Equal(90, t.ToMinutes())
}

func (s *PureTimeSuite) TestFromMinutes() {
	t := PureTime{}
	t.FromMinutes(90)
	s.Equal(1, t.Hour)
	s.Equal(30, t.Minute)
}

func (s *PureTimeSuite) TestMarshalJSON() {
	t := PureTimeStruct{
		Time: PureTime{Hour: 1, Minute: 30},
	}
	b, err := json.Marshal(t)
	s.NoError(err)
	s.Equal(`{"time":"01:30"}`, string(b))
}

func (s *PureTimeSuite) TestUnmarshalJSON() {
	t := PureTimeStruct{}
	err := json.Unmarshal([]byte(`{"time": "01:30"}`), &t)
	s.Require().NoError(err)
	s.Equal(1, t.Time.Hour)
	s.Equal(30, t.Time.Minute)
}

func (s *PureTimeSuite) TestUnmarshalJSON_Invalid() {
	t := PureTimeStruct{}
	err := json.Unmarshal([]byte(`{"time": "01:60"}`), &t)
	s.Error(err)
}

func TestPureTimeSuite(t *testing.T) {
	suite.Run(t, new(PureTimeSuite))
}
