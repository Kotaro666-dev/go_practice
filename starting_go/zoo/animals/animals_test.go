package animals

import "testing"

func TestElephantFeed(t *testing.T) {
  expect := "Grass";
  result := ElephantFeed();

  if expect != result {
    t.Errorf("%s != %s", expect, result);
    }
  }

func TestMonkeyFeed(t *testing.T) {
  expect := "Banana";
  result := MonkeyFeed();

  if expect != result {
    t.Errorf("%s != %s", expect, result);
    }
  }

func TestRabbitFeed(t *testing.T) {
  expect := "Carrot";
  result := RabbitFeed();

  if expect != result {
    t.Errorf("%s != %s", expect, result);
    }
  }
