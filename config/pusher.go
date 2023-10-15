package config

import (
	"github.com/pusher/pusher-http-go/v5"
)

func NewPusherClient() *pusher.Client {
	pusherClient := pusher.Client{
		AppID:   "APP_ID",
		Key:     "APP_KEY",
		Secret:  "APP_SECRET",
		Cluster: "APP_CLUSTER",
	}

	return &pusherClient
}