package service

type App struct {
}

func (a *App) HandleRequest(url, method string) (int, string) {

	if url == APP_STATUS_URL && method == GET {
		return 200, "Ok"
	}

	if url == CREATE_USER_URL && method == POST {
		return 201, "Created"
	}

	return 404, "Not found"
}
