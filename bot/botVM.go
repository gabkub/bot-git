package bot

import (
	"bot-git/config"
	"bot-git/main/connection"
	"bot-git/schedule"
	"github.com/mattermost/mattermost-server/model"
	"log"
	"math/rand"
	"sync"
	"time"
)

var mux = &sync.Mutex{}

func Start() {

	log.Println("Bot has started.")
	seedRand()

	helpHandler.Init(handlers)

	go func() {
		schedule.Start()
		for {
			select {

			case <-connection.Websocket.PingTimeoutChannel:
				mux.Lock()
				log.Println("Websocket PingTimeout.")
				config.ConnectionCfg.Client.Logout()
				connection.Connect()
				mux.Unlock()

			case event := <-connection.Websocket.EventChannel:
				mux.Lock()
				if event != nil {
					if event.IsValid() && isMessage(event.Event) {
						handleEvent(event)
					}
				}
				mux.Unlock()
			}
		}
	}()
	// block to the go function
	select {}
}

func seedRand() {
	rand.Seed(time.Now().UnixNano())
}

func isMessage(eventType string) bool {
	if eventType == model.WEBSOCKET_EVENT_POSTED {
		return true
	}
	return false
}
