package hospital

type Handler interface {
	Execute(p *Patient)
	SetNext(h Handler)
}
