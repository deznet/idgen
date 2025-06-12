package idgen

// Provider id生成器Provider接口
type Provider interface {
	//Int64 生成Int64 id
	Int64() (int64, error)

	//String 生层字符串 id
	String() (string, error)
}

type IdGen struct {
	Provider Provider
}

func NewIdGen(p Provider) *IdGen {
	return &IdGen{Provider: p}
}

func (c *IdGen) Int64() (int64, error) {
	return c.Provider.Int64()
}

func (c *IdGen) String() (string, error) {
	return c.Provider.String()
}
