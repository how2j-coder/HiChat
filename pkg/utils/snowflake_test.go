package utils

import "testing"

func TestSnowflake(t *testing.T)  {
	s, _ := NewSnowflake(32)
	generate, err := s.Generate()
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(generate)
}