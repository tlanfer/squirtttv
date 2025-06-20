package internal

import (
	"fmt"
)

type ChatRole string

const (
	ChatRolePleb ChatRole = "pleb"
	ChatRoleSub  ChatRole = "sub"
	ChatRoleMod  ChatRole = "mod"
)

type ChatMessage struct {
	User    string
	Role    ChatRole
	Message string
}

func (s ChatMessage) String() string {
	return fmt.Sprintf("[ %v (%v) | %v ]", s.User, s.Role, s.Message)
}
