package v1

import (
	"fmt"

	"gitlab.com/robotomize/gb-golang/homework/03-04-umanager/pkg/api/apiv1"
)

type serverInterface interface {
	apiv1.ServerInterface
}

var _ serverInterface = (*Handler)(nil)

func New(usersRepository usersClient, linksRepository linksClient, logger Logger) *Handler {
	return &Handler{
		usersHandler: newUsersHandler(usersRepository, logger),
		linksHandler: newLinksHandler(linksRepository, logger),
		logger:       logger,
	}
}

type Logger interface {
	Error(msg string, args ...interface{})
}

type Handler struct {
	*usersHandler
	*linksHandler
	logger Logger
}

func (h *Handler) handleError(err error) {
	if err != nil {
		h.logger.Error(fmt.Sprintf("Error occurred: %s", err.Error()))
	}
}
