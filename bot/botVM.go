package bot

import (
	"bot-git/config"
	"bot-git/footballDatabase"
	"bot-git/logg"
	"bot-git/main/connection"
	"bot-git/schedule"
	"github.com/mattermost/mattermost-server/model"
	"log"
	"sync"
)

var mux = &sync.Mutex{}

func Start() {

	logg.WriteToFile("Bot has started.")
	log.Println("Bot has started.")

	go func() {
		schedule.Start()
		footballDatabase.CreateTableDB()
		for {
			select {

			case <-connection.Websocket.PingTimeoutChannel:
				mux.Lock()
				logg.WriteToFile("Websocket PingTimeout.")
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

func isMessage(eventType string) bool {
	if eventType == model.WEBSOCKET_EVENT_POSTED {
		return true
	}
	return false
}
