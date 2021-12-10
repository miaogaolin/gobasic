package unittest

import (
	"testing"
	"time"
)

func TestA(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second)
}

func TestB(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second)
}

func TestC(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second)
}
