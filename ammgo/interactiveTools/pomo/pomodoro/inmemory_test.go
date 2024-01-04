package pomodoro_test

import (
	"testing"

	"agmmtoo.me/ammgo/interactiveTools/pomo/pomodoro"
	"agmmtoo.me/ammgo/interactiveTools/pomo/pomodoro/repository"
)

func getRepo(t *testing.T) (pomodoro.Repository, func()) {
	t.Helper()
	return repository.NewInMemoryRepo(), func() {}
}
