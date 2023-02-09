package stores

import (
	"github.com/gorilla/sessions"
)

type Store interface {
	sessions.Store
	Options(sessions.Options)
}
