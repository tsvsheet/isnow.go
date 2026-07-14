package main

import "testing"

func TestRun(t *testing.T) {
	if code := run([]string{"isnow", "canon", "6"}); code != 0 {
		t.Fatalf("run canon = %d", code)
	}
	if code := run([]string{"isnow", "25"}); code != 2 {
		t.Fatalf("run bad = %d", code)
	}
}

func TestMain_Exits(t *testing.T) {
	old := osExit
	defer func() { osExit = old }()
	got := -1
	osExit = func(code int) { got = code }
	oldArgs := args
	defer func() { args = oldArgs }()
	args = []string{"isnow", "canon", "6"}
	main()
	if got != 0 {
		t.Fatalf("main exit = %d", got)
	}
}
