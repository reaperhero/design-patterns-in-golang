package Proxy

import (
	"testing"
)

func TestAgentTask_RentHouse(t *testing.T) {
	task := &Task{}
	agent := NewAgentTask(task)
	agent.RentHouse("北京代理", 8000)
}
