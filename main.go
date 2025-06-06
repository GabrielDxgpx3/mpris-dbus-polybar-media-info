package main

import (
	"fmt"
	"os"

	"github.com/godbus/dbus/v5"
)

type MprisMetadata struct {
	artist []string
	title  string
}

func newMprisMetadata(variantData map[string]dbus.Variant) MprisMetadata {
	return MprisMetadata{
		variantData["xesam:artist"].Value().([]string),
		variantData["xesam:title"].Value().(string),
	}
}

func getPlayerProperty(object *dbus.BusObject, property string) interface{} {
	path := "org.mpris.MediaPlayer2.Player." + property
	variant, err := (*object).GetProperty(path)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(0)
	}

	return variant.Value()
}

func getFormattedOuput(metadata MprisMetadata) string {
	return metadata.title + " - " + metadata.artist[0]
}

func main() {
	conn, err := dbus.ConnectSessionBus()

	if err != nil {
		panic(err)
	}

	defer conn.Close()
	obj := conn.Object("org.mpris.MediaPlayer2.playerctld", "/org/mpris/MediaPlayer2")

	isPlayerIdentified := getPlayerProperty(&obj, "CanPlay").(bool)
	palybackStatus := getPlayerProperty(&obj, "PlaybackStatus").(string)

	if !isPlayerIdentified || palybackStatus != "Playing" {
		fmt.Print(" ")
		os.Exit(0)
	}

	response := getPlayerProperty(&obj, "Metadata").(map[string]dbus.Variant)
	metadata := newMprisMetadata(response)

	fmt.Print(getFormattedOuput(metadata))
}
