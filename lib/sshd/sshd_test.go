package sshd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoteRun(t *testing.T) {
	t.Run("ls", func(t *testing.T) {
		c := Client{}
		c.EnrollWithPassword("tecnalia", "Tecnalia#0000")

		err := c.Connect("172.26.252.101", "22")
		assert.Nil(t, err)
		data, cmdErr := c.SyncCommand([]string{"ls -alh"})
		assert.Nil(t, cmdErr)
		t.Log(string(data))
		data, cmdErr = c.SyncCommand([]string{"netstat -napo | grep :53"})
		assert.Nil(t, cmdErr)
		t.Log(string(data))
	})
}
