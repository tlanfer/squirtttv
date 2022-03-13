package companion

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"strings"
)

type ChatRole int

const (
	ChatRolePleb ChatRole = 0
	ChatRoleSub  ChatRole = 1
	ChatRoleMod  ChatRole = 3
)

type ChatMessage struct {
	User    string
	Role    ChatRole
	Message string
}

func (s ChatMessage) String() string {
	return fmt.Sprintf("[ %v (%v) | %v ]", s.User, s.Role, s.Message)
}

func (c ChatRole) MarshalYAML() (interface{}, error) {
	switch c {
	case ChatRolePleb:
		return "pleb", nil
	case ChatRoleSub:
		return "sub", nil
	case ChatRoleMod:
		return "mod", nil
	}
	return "pleb", nil
}

func (c *ChatRole) UnmarshalYAML(value *yaml.Node) error {
	v := strings.ToLower(value.Value)
	switch v {
	case "mod":
		*c = ChatRoleMod
	case "sub":
		*c = ChatRoleSub
	default:
		*c = ChatRolePleb
	}
	return nil
}
