package config

import (
	"gopkg.in/yaml.v3"
	"strings"
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

func (s *SquirtPattern) UnmarshalJSON(bytes []byte) error {
	str := string(bytes)
	str = strings.TrimPrefix(str, "\"")
	str = strings.TrimSuffix(str, "\"")
	parts := strings.Split(str, ",")
	var durations []time.Duration
	for _, part := range parts {
		if duration, err := time.ParseDuration(strings.TrimSpace(part)); err == nil {
			durations = append(durations, duration)
		}
	}

	*s = durations
	return nil
}

func (s SquirtPattern) MarshalJSON() ([]byte, error) {
	var parts []string
	for _, duration := range s {
		parts = append(parts, duration.String())
	}
	s2 := "\"" + strings.Join(parts, ",") + "\""
	return []byte(s2), nil
}
