package trayicon

import _ "embed"

//go:embed icon_on_on.ico
var iconOnOn []byte

//go:embed icon_on_off.ico
var iconOnOff []byte

//go:embed icon_off_on.ico
var iconOffOn []byte

//go:embed icon_off_off.ico
var iconOffOff []byte

//go:embed icon_active.ico
var iconActive []byte
