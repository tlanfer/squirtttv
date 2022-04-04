package companion

import (
	"gopkg.in/yaml.v3"
	"time"
)

type SquirtPattern []time.Duration

func (s SquirtPattern) MarshalYAML() (interface{}, error) {
	p := s
	switch len(p) {
	case 0:
		return nil, nil
	case 1:
		return p[0], nil
	default:
		return []time.Duration(p), nil
	}
}

func (s *SquirtPattern) UnmarshalYAML(value *yaml.Node) error {
	var p []time.Duration
	switch value.Kind {
	case yaml.ScalarNode:
		if duration, err := time.ParseDuration(value.Value); err == nil {
			p = append(p, duration)
		}
	case yaml.SequenceNode:
		for _, node := range value.Content {
			if duration, err := time.ParseDuration(node.Value); err == nil {
				p = append(p, duration)
			}
		}
	}
	*s = p
	return nil
}
