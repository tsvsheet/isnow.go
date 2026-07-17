package app

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/tsvsheet/isnow.go/internal/constants"
)

func TestParseZone(t *testing.T) {
	if loc, err := ParseZone(""); err != nil || loc != time.Local {
		t.Fatalf("ParseZone(\"\") = %v, %v", loc, err)
	}
	if _, err := ParseZone("America/New_York"); err != nil {
		t.Fatalf("ParseZone(NY) = %v", err)
	}
	if _, err := ParseZone("Bogus/Zone"); !errors.Is(err, constants.ErrBadZone) {
		t.Fatalf("ParseZone(bad) = %v", err)
	}
}

func TestParseInstant(t *testing.T) {
	for _, s := range []string{
		"2026-07-14T12:00:00Z",
		"2026-07-14T12:00:00",
		"2026-07-14 12:00:00",
		"2026-07-14",
	} {
		if _, err := ParseInstant(s, time.UTC); err != nil {
			t.Fatalf("ParseInstant(%q) = %v", s, err)
		}
	}
	if _, err := ParseInstant("not-a-time", time.UTC); !errors.Is(err, constants.ErrBadTime) {
		t.Fatalf("ParseInstant(bad) = %v", err)
	}
}

func TestRealSleep(t *testing.T) {
	if err := RealSleep(context.Background(), time.Millisecond); err != nil {
		t.Fatalf("RealSleep = %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := RealSleep(ctx, time.Hour); !errors.Is(err, context.Canceled) {
		t.Fatalf("RealSleep(cancelled) = %v", err)
	}
}

func TestRealSpawn(t *testing.T) {
	if err := RealSpawn(context.Background(), "true", nil); err != nil {
		t.Fatalf("RealSpawn(true) = %v", err)
	}
	if err := RealSpawn(context.Background(), "this-command-does-not-exist-isnow", nil); err == nil {
		t.Fatal("RealSpawn(missing) should error")
	}
}
