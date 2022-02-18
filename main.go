package main

import (
	"github.com/hugolgst/rich-go/client"
	"time"
)

func main() {
	err := client.Login("846539372131385444")
	if err != nil {
		panic(err)
	}

	now := time.Now()
	err = client.SetActivity(client.Activity{
		State:      "Exchanges Connected",
		Details:    "Multi-Exchange Trading",
		LargeImage: "analytic-icon-11449",
		LargeText:  "Trade Crypto Like Never Before",
		SmallImage: "discord_gradient",
		SmallText:  "Cryptonovae",
		Party: &client.Party{
			ID:         "-1",
			Players:    2,
			MaxPlayers: 5,
		},
		Timestamps: &client.Timestamps{
			Start: &now,
		},
		Buttons: []*client.Button{
			&client.Button{
				Label: "Trade Now",
				Url:   "https://cryptonovae.com",
			},
		},
	})

	if err != nil {
		panic(err)
	}

	for {
		time.Sleep(time.Second * 100)
	}
}
