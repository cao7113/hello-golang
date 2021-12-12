package file

import (
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func (s *FileSuite) TestReadWrite() {
	pwd, err := os.Getwd()
	s.Nil(err)
	log.Println("pwd = ", pwd)

	fl := "../README.md"
	bs, err := ioutil.ReadFile(fl)
	s.Nil(err)
	log.Println(fl, string(bs))
}

func TestFileSuite(t *testing.T) {
	suite.Run(t, &FileSuite{})
}

type FileSuite struct {
	suite.Suite
}
