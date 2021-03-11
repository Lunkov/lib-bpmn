package bpmn

import (
  "flag"
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
  flag.Set("alsologtostderr", "true")
  flag.Set("log_dir", ".")
  // flag.Set("v", "9")
  flag.Parse()
  
  s := NewServices()
  
  s.LoadFromFiles("etc.test")
  
  assert.Equal(t, int64(2), s.Count())
}
