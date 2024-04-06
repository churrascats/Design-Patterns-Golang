package prototype

type Cloneable interface {
	Print(string)
	Clone() Cloneable
}
