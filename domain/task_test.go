package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CreateTask(t *testing.T) {
	var task *Task = NewTask("buy egg")

	assert.NotNil(t, task)
	assert.Equal(t, "buy egg", task.Title())
}
