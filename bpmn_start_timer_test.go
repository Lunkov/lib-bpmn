package bpmn

import (
  "flag"
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestCheckParseTime(t *testing.T) {
  flag.Set("alsologtostderr", "true")
  flag.Set("log_dir", ".")
  flag.Set("v", "9")
  flag.Parse()

  assert.Equal(t, "25h0m0s", parseDuration("P1DT1H").String())
  assert.Equal(t, "15m0s", parseDuration("PT15M").String())
  assert.Equal(t, "30m0s", parseDuration("PT0.5H").String())
  assert.Equal(t, float64(88200), parseDuration("P1DT0.5H").Seconds())
}


