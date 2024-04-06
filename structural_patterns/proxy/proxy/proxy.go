package proxy

import "proxy/service"

type ServiceProxy struct {
	app               service.App
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func NewServiceProxy() *ServiceProxy {
	return &ServiceProxy{
		app:               service.App{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (s *ServiceProxy) HandleRequest(url, method string) (int, string) {
	isAllowed := s.checkRateLimit(url)

	if !isAllowed {
		return 403, "Forbidden"
	}

	return s.app.HandleRequest(url, method)
}

func (s *ServiceProxy) checkRateLimit(url string) bool {
	if s.rateLimiter[url] >= s.maxAllowedRequest {
		return false
	}

	s.rateLimiter[url]++

	return true
}
