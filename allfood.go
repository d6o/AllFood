package main

import (
	"flag"
)

func main() {
	a := App{}

	radius := flag.Int("radius", 0, "Radius")
	lat := flag.String("lat", "", "Latitude")
	lng := flag.String("lng", "", "Longitute")
	gKey := flag.String("google-key", "", "Google Api Key")
	fbID := flag.String("fb-app-id", "", "Facebook APP ID")
	fbSecret := flag.String("fb-app-secret", "", "Facebook APP Secret")
	fsID := flag.String("fs-app-id", "", "FourSquare APP ID")
	fsSecret := flag.String("fs-app-secret", "", "FourSquare APP Secret")

	flag.Parse()

	a.Initialize(*radius, *lat, *lng)

	if *gKey != "" {
		google := NewGoogle(*gKey)
		a.AddProvider(google)
	}

	if *fbID != "" && *fbSecret != "" {
		fb := NewFacebook(*fbID, *fbSecret)
		fb.GetToken()
		a.AddProvider(fb)
	}

	if *fsID != "" && *fsSecret != "" {
		fs := NewFourSquare(*fsID, *fsSecret)
		a.AddProvider(fs)
	}

	a.Run()
}
