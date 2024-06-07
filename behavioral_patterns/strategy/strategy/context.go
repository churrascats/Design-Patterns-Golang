package strategy

type Context struct {
	strategy Strategy
}

func NewContext() Context {
	return Context{}
}

func (c *Context) SetStrategy(s Strategy) {
	c.strategy = s
}

func (c *Context) ShowArrayItems(array []int) {
	c.strategy.ShowArrayItems(array)
}
