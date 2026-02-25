package test

import (
	"strings"
	"testing"
	"time"

	"github.com/keypad/clock/src/clock"
)

func TestRenderutc(t *testing.T) {
	base := time.Date(2026, 2, 25, 10, 30, 45, 0, time.UTC)
	text, err := clock.Render("utc", base)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}
	if !strings.Contains(text, "utc  2026-02-25 10:30:45") {
		t.Fatalf("unexpected text: %s", text)
	}
	if !strings.Contains(text, "(+00:00)") {
		t.Fatalf("missing offset: %s", text)
	}
}

func TestRenderest(t *testing.T) {
	base := time.Date(2026, 2, 25, 10, 30, 45, 0, time.UTC)
	text, err := clock.Render("est", base)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}
	if !strings.Contains(text, "est  2026-02-25 05:30:45") {
		t.Fatalf("unexpected text: %s", text)
	}
}

func TestError(t *testing.T) {
	_, err := clock.Render("bad", time.Now())
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestZones(t *testing.T) {
	text := clock.Zones()
	if text != "cet cst est ist jst local mst pst utc" {
		t.Fatalf("unexpected zones: %s", text)
	}
}
