package main

import (
  "testing"
)

func TestGoToRome(t *testing.T) {
  if goToRome(69) != "LXIX" {
    t.Error("expected string is LXIX")
  }
  if goToRome(420) != "CDXX" {
    t.Error("expected string is CDXX")
  }
}
