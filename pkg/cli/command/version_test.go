package command_test

import (
	"bytes"
	"testing"

	"github.com/phss/ogs-notifier/pkg/cli/command"
	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	output, err := runCommand("version")

	assert.NoError(t, err)
	assert.Equal(t, "ogscli v0.1\n", output)
}

func runCommand(args ...string) (output string, err error) {
	buf := new(bytes.Buffer)
	command.Root.SetOutput(buf)
	command.Root.SetArgs(args)
	_, err = command.Root.ExecuteC()
	return buf.String(), err
}
