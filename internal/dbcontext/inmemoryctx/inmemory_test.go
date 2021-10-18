package inmemoryctx

import (
	"testing"
)

func TestNew(t *testing.T) {

	ctx := New()

	if ctx == nil {
		t.Fatalf("ctx == nil")
	}

}

func TestCreateList(t *testing.T) {

	ctx := New()

	err := ctx.CreateList("foo")

	if err != nil {
		t.Fatalf("Expected nil value in err, val = %v", err)
	}

	err = ctx.CreateList("foo")

	if err == nil {
		t.Fatalf("Expected err 'There is already a receivers list for user %v ', val = %v", "foo", err)
	}

}

func TestDeleteList(t *testing.T) {

	ctx := New()

	err := ctx.CreateList("foo")

	if err != nil {
		t.Fatal(err)
	}

	err = ctx.DeleteList("foo")

	if err != nil {
		t.Fatal(err)
	}

}

func TestAddToList(t *testing.T) {

	ctx := New()

	err := ctx.CreateList("foo")

	if err != nil {
		t.Fatal(err)
	}

	ctx.AddToList("foo", "bar1")
	ctx.AddToList("foo", "bar2")
	ctx.AddToList("foo", "bar3")

}
