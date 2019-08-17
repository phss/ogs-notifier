package cli

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	output, err := runCommand("version")

	assert.NoError(t, err)
	assert.Equal(t, "ogscli v0.1\n", output)
}

func runCommand(args ...string) (output string, err error) {
	buf := new(bytes.Buffer)
	rootCmd.SetOutput(buf)
	rootCmd.SetArgs(args)
	_, err = rootCmd.ExecuteC()
	return buf.String(), err
}
