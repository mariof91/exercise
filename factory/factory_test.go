package factory

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type factoryUnitTestSuite struct {
	suite.Suite
	adapter *Factory
}

func (s *factoryUnitTestSuite) SetupSuite() {

	s.adapter = New()
}

func TestFactoryUnitTestSuite(t *testing.T) {
	suite.Run(t, &factoryUnitTestSuite{})
}

func (s *factoryUnitTestSuite) TestSample() {
	//code here
	ch := s.adapter.StartAssemblingProcess(2)

	s.Assert().NotNil(ch)

	var cars [] int
	c1 := <- ch
	c2 := <- ch
	cars = append(cars, c1.Id)
	cars = append(cars, c2.Id)

	s.Assert().Contains( cars, 0)
	s.Assert().Contains( cars, 1)
	s.Assert().NotContains(cars, 2)
	s.Assert().NotEmpty(c1.TestingLog)
	s.Assert().NotEmpty(c1.AssembleLog)


}
