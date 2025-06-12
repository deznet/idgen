package idgen

import (
	"github.com/deznet/idgen/provider/snowflake"
	"testing"
)

func TestIdGen(t *testing.T) {
	p, _ := snowflake.NewNode(0)
	c := NewIdGen(p)
	i, err := c.Int64()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(i)
	s, err := c.String()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
}
