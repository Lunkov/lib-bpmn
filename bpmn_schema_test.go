package bpmn

import (
  "flag"
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/golang/glog"
)

func TestCheckBPMN(t *testing.T) {
  flag.Set("alsologtostderr", "true")
  flag.Set("log_dir", ".")
  flag.Set("v", "9")
  flag.Parse()

  glog.Info("Logging configured: TestCheckBPMN")

  var b Info
  ok, msg := b.LoadXML("etc.test/wrong_diagram_1.bpmn")
  assert.Equal(t, true, ok)
  assert.Equal(t, "OK", msg)
  
  ok, msg = b.Validate(nil)
  assert.Equal(t, false, ok)
  assert.Equal(t, "ERR: SendTask(Task_1lock54) no has outgoing", msg)

  ok, msg = b.LoadXML("etc.test/wrong_diagram_2.bpmn")
  assert.Equal(t, true, ok)
  assert.Equal(t, "OK", msg)
  
  ok, msg = b.Validate(nil)
  assert.Equal(t, false, ok)
  assert.Equal(t, "ERR: SendTask(Task_1lock54) no has outgoing", msg)

  srvc := NewServices()
  srvc.LoadFromFiles("etc.test")

  ok, msg = b.LoadXML("etc.test/new_user.bpmn")
  assert.Equal(t, true, ok)
  assert.Equal(t, "OK", msg)
  
  ok, msg = b.Validate(srvc)
  assert.Equal(t, true, ok)
  assert.Equal(t, "OK", msg)

}

