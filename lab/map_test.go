package lab

import (
	"fmt"
	"sort"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/sirupsen/logrus"
)

func (s *LabSuite) TestMap() {
	var m map[string]int
	s.Nil(m)
	s.EqualValues(0, m["A"])
	s.Panics(func() {
		m["b"] = 1
	})
}

// https://blog.golang.org/go-maps-in-action
// https://www.callicoder.com/golang-maps/
func TestMapBasic(t *testing.T) {
	m := make(map[string]string)
	m["name"] = "geek"
	assert.Equal(t, "geek", m["name"])
	delete(m, "name")
	assert.Equal(t, 0, len(m))
	// test existence
	_, ok := m["name"]
	assert.Equal(t, ok, false)

	m["country"] = "China"
	for key, value := range m {
		logrus.Infoln("Key:", key, "Value:", value)
	}
}

func TestMapKeysInOrder(t *testing.T) {
	//When iterating over A map with A range loop, the iteration order is not specified and is not guaranteed to
	//be the same from one iteration to the next. If you require A stable iteration order you must maintain A separate data structure
	//that specifies that order. This example uses A separate sorted slice of keys to print A map[int]string in key order:
	m := map[int]string{
		3: "jian",
		1: "cao",
		2: "rui",
	}
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k])
	}
}

func TestMapOfArray(t *testing.T) {
	m := make(map[int][]int)
	a := m[0]
	a = append(a, 1)
	a = append(a, 2)
	m[0] = a
	assert.Equal(t, 2, len(a))
	assert.Equal(t, 2, a[1])
	logrus.Infof("origin map %+v", m)
}
