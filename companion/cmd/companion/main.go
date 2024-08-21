package main

import (
	"companion/internal"
	"companion/internal/adapter/inbound/streamelements"
	"companion/internal/adapter/inbound/streamlabs"
	"companion/internal/adapter/inbound/trayicon"
	"companion/internal/adapter/inbound/twitchchat"
	"companion/internal/adapter/inbound/ui"
	"companion/internal/adapter/inbound/ui/config"
	"companion/internal/adapter/inbound/ui/discovery"
	"companion/internal/adapter/inbound/ui/status"
	"companion/internal/adapter/inbound/ui/test"
	"companion/internal/adapter/inbound/ui/updates"
	"companion/internal/adapter/inbound/yamlconfig"
	"companion/internal/adapter/outbound/exchangerate"
	"companion/internal/adapter/outbound/scanner"
	"companion/internal/eventprocssor"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	version = "0.0.1"
)

func main() {

	yamlconfig.Init()

	icon := trayicon.New()

	err := setupLogging()
	if err != nil {
		icon.ErrorMessage("%v", err.Error())
		return
	}
	time.Sleep(100 * time.Millisecond)
	log.Println("Companion starting...")

	mux := http.NewServeMux()
	mux.Handle("/api/config", config.NewHandler())
	mux.Handle("/api/status", status.NewHandler())
	mux.Handle("/api/discovery", discovery.NewHandler())
	mux.Handle("/api/test", test.NewHandler())
	mux.Handle("/api/version", updates.NewHandler(version))
	mux.Handle("/", ui.NewHandler())

	srv := http.Server{
		Addr:    ":3080",
		Handler: mux,
	}

	go srv.ListenAndServe()

	events := make(chan internal.StreamEvent)
	messages := make(chan internal.ChatMessage)

	exchangerate.New()
	scanner.New()
	streamlabs.New(events)
	streamelements.New(events)
	twitchchat.New(events, messages)
	eventprocssor.New(events, messages)

	defer catchCrashes()

	<-icon.OnQuit()
	icon.Quit()
	_ = srv.Close()
}

func setupLogging() error {
	file, err := os.OpenFile("companion.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	log.SetOutput(io.MultiWriter(file, os.Stdout))

	return nil
}

func catchCrashes() {
	if r := recover(); r != nil {
		log.Fatalf("Panic: %v", r)
	}
}
