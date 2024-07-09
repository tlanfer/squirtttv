package eventprocssor

import (
	"companion/internal/config"
	"strconv"
	"testing"
)

func TestFloatingPointParsing(t *testing.T) {
	str := "1"
	amount, err := strconv.ParseFloat(str, 64)
	if err != nil {
		t.Fatalf("Error parsing float: %s", err)
	}
	if amount != 1.0 {
		t.Fatalf("Expected 1.0, got %f", amount)
	}
}

func TestFindMatch(t *testing.T) {
	amount := 12.34
	events := []config.Event{
		{Match: "exact", Amount: 12.34},
		{Match: "minimum", Amount: 10.00},
	}

	event := findMatch(amount, events)

	if event == nil {
		t.Fatal("Expected event to be found")
	}
	if event.Match != "exact" {
		t.Fatalf("Expected event to be exact, got %s", event.Match)
	}
}

func TestFindMatchNoMatch(t *testing.T) {
	amount := 12.34
	events := []config.Event{
		{Match: "minimum", Amount: 15.00},
	}

	event := findMatch(amount, events)

	if event != nil {
		t.Fatal("Expected no event to be found")
	}
}

func TestFindMatchBestMatch(t *testing.T) {
	amount := 12.34
	events := []config.Event{
		{Match: "minimum", Amount: 10.00},
		{Match: "minimum", Amount: 15.00},
	}

	event := findMatch(amount, events)

	if event == nil {
		t.Fatal("Expected event to be found")
	}
	if event.Amount != 10.00 {
		t.Fatalf("Expected event amount to be 10.00, got %f", event.Amount)
	}
}

func TestFindMatchBestMatchMultiple(t *testing.T) {
	amount := 12.34
	events := []config.Event{
		{Match: "minimum", Amount: 10.00},
		{Match: "minimum", Amount: 15.00},
		{Match: "minimum", Amount: 12.00},
	}

	event := findMatch(amount, events)

	if event == nil {
		t.Fatal("Expected event to be found")
	}
	if event.Amount != 12.00 {
		t.Fatalf("Expected event amount to be 12.00, got %f", event.Amount)
	}
}

func TestFindMatchBestMatchExact(t *testing.T) {
	amount := 12.34
	events := []config.Event{
		{Match: "minimum", Amount: 10.00},
		{Match: "minimum", Amount: 15.00},
		{Match: "exact", Amount: 12.34},
	}

	event := findMatch(amount, events)

	if event == nil {
		t.Fatal("Expected event to be found")
	}
	if event.Match != "exact" {
		t.Fatalf("Expected event to be exact, got %s", event.Match)
	}
}

func TestFindMatchBestMatchExactMultiple(t *testing.T) {
	amount := 12.34
	events := []config.Event{
		{Match: "minimum", Amount: 10.00},
		{Match: "minimum", Amount: 15.00},
		{Match: "exact", Amount: 12.34},
		{Match: "exact", Amount: 12.34},
	}

	event := findMatch(amount, events)

	if event == nil {
		t.Fatal("Expected event to be found")
	}
	if event.Match != "exact" {
		t.Fatalf("Expected event to be exact, got %s", event.Match)
	}
}

func TestFindMatchBestMatchExactMultipleDifferentAmounts(t *testing.T) {
	amount := 12.34
	events := []config.Event{
		{Match: "minimum", Amount: 10.00},
		{Match: "minimum", Amount: 15.00},
		{Match: "exact", Amount: 12.34},
		{Match: "exact", Amount: 12.35},
	}

	event := findMatch(amount, events)

	if event == nil {
		t.Fatal("Expected event to be found")
	}
	if event.Amount != 12.34 {
		t.Fatalf("Expected event amount to be 12.34, got %f", event.Amount)
	}
}

func TestFindMatchBestMatchMinimumMultipleDifferentAmounts(t *testing.T) {
	amount := 12.36
	events := []config.Event{
		{Match: "minimum", Amount: 10.00},
		{Match: "minimum", Amount: 15.00},
		{Match: "minimum", Amount: 12.35},
	}

	event := findMatch(amount, events)

	if event == nil {
		t.Fatal("Expected event to be found")
	}
	if event.Amount != 12.35 {
		t.Fatalf("Expected event amount to be 12.35, got %f", event.Amount)
	}
}
