package snowflake

import "testing"

// TestNewNode 测试创建node
func TestNewNode(t *testing.T) {
	//正常测试
	node, err := NewNode(0)
	if err != nil {
		t.Fatalf("Node创建失败, %s", err)
	}

	t.Log(node.Int64())
	t.Log(node.String())

	//超出nodeid范围测试
	_, err = NewNode(5000)
	if err == nil {
		t.Fatalf("超过nodeid被执行")
	}

	_, err = NewNode(-1)
	if err == nil {
		t.Fatalf("超过nodeid被执行")
	}

}
