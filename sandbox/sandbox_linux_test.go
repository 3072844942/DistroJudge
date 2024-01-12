//go:build linux

package sandbox

import "testing"

func TestCommand(t *testing.T) {
	name := "cmd"
	cmd := Command(name)

	if err := cmd.Run(); err != nil {
		t.Errorf("command %s running err. err: %v", name, err)
	}
}
