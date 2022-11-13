package max31865_test

import (
	"errors"
	"github.com/a-clap/iot/pkg/max31865"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MaxSuite struct {
	suite.Suite
}

type MaxTransferMock struct {
	mock.Mock
}

func (m *MaxTransferMock) Close() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MaxTransferMock) ReadWrite(write []byte) (read []byte, err error) {
	args := m.Called(write)
	return args.Get(0).([]byte), args.Error(1)
}

var (
	mocker *MaxTransferMock
)

func TestMax31865(t *testing.T) {
	suite.Run(t, new(MaxSuite))
}

func (m *MaxSuite) SetupTest() {
	mocker = new(MaxTransferMock)
}

func (m *MaxSuite) TestMaxInterfaceError() {
	mocker.On("ReadWrite", mock.Anything).Return([]byte{}, errors.New("interface broken"))
	max, err := max31865.New(mocker)
	m.Equal(max31865.ErrInterface, err)
	m.Equal(nil, max)
}

func (m *MaxSuite) TestMaxInterfaceReturnsZeroes() {
	mocker.On("ReadWrite", mock.Anything).Return([]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, nil)
	max, err := max31865.New(mocker)
	m.Equal(max31865.ErrReadZeroes, err)
	m.Equal(nil, max)
}

func (m *MaxSuite) TestMaxInterfaceReturnsFF() {
	mocker.On("ReadWrite", mock.Anything).Return([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, nil)
	max, err := max31865.New(mocker)
	m.Equal(max31865.ErrReadFF, err)
	m.Equal(nil, max)
}
