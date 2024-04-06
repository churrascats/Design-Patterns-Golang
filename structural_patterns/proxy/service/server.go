package service

type Server interface {
	HandleRequest(string, string) (int, string)
}
