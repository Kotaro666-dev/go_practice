package main

import "testing"

func TestAppName(t *testing.T) {
  expect := "Zoo Application";
  result := AppName();

  if expect != result {
    t.Errorf("%s != %s", expect, result);
  }
}
