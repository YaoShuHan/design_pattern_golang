package Adapter

import (
	"testing"
)

func TestAdapter(t *testing.T) {
	forward := NewForwards("Forward")
	guards := NewGuards("Guards")
	center := NewTranslator("Center")

	forward.Attack()
	guards.Attack()
	center.Attack()
}
