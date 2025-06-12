package snowflake

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

// 说明
// 0 + 41位时间戳（毫秒）+ 序列 + 节点
const (
	//节点位数
	nodeBits uint8 = 10

	//序列位数
	seqBits uint8 = 12

	//节点ID的最大值，用于防止溢出
	nodeMask int64 = -1 ^ (-1 << nodeBits)

	//序列数最大值，用于防止溢出
	seqMask int64 = -1 ^ (-1 << seqBits)

	//起始时间戳，单位：毫秒
	epoch int64 = 1577808000000

	//时间左偏移量
	timeShift = nodeBits + seqBits

	//序列左偏移量
	seqShift = nodeBits
)

type Node struct {
	//锁
	mu sync.Mutex

	//节点id
	nodeId int64

	//上次id生成的时间戳（毫秒）
	lastTime int64

	//毫秒中序列
	seq int64
}

// NewNode 初始化一个节点
func NewNode(nodeId int64) (*Node, error) {
	if nodeId < 0 || nodeId > nodeMask {
		return nil, errors.New("invalid nodeId")
	}
	return &Node{
		nodeId:   nodeId,
		lastTime: 0,
		seq:      0,
	}, nil
}

// Generate 生成id
func (n *Node) Generate() int64 {
	n.mu.Lock()
	defer n.mu.Unlock()
	now := time.Now().UnixMilli()
	//时间回拨
	if now < n.lastTime {
		//先看lastTime的seq是否有余量
		n.seq = (n.seq + 1) & seqMask
		//如果没有余量，毫秒+1
		if n.seq == 0 {
			n.lastTime = n.lastTime + 1
		}
	} else {
		if now == n.lastTime {
			n.seq = (n.seq + 1) & seqMask
			//seq溢出，等待时间更新
			if n.seq == 0 {
				for now <= n.lastTime {
					now = time.Now().UnixMilli()
				}
			}
		} else {
			n.seq = 0
		}
		n.lastTime = now
	}
	return (n.lastTime-epoch)<<timeShift | n.seq<<seqShift | n.nodeId
}

// Int64 同Generate
func (n *Node) Int64() (int64, error) {
	return n.Generate(), nil
}

// String 字符串
func (n *Node) String() (string, error) {
	return strconv.FormatInt(n.Generate(), 10), nil
}
