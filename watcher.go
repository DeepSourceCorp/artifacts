package artifacts

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/getsentry/sentry-go"
)

// Watches the broker config for any "WRITE" events
// Exits with status code 1 after any changes to config
func watchBrokerConfigForChanges(filePath string, reloadFunc func() error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println(err)
		sentry.CaptureException(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println(event)
				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Remove == fsnotify.Remove {
					log.Println("Received event:", event)
					log.Println("modified file:", event.Name)

					if strings.HasSuffix(event.Name, filePath) {
						watcher.Close()
						if reloadErr := reloadFunc(); reloadErr != nil {
							log.Fatalln("Failed to reload janus. Exiting.") // skipcq
						}
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Watcher error:", err)
			}
		}
	}()

	err = watcher.Add(filePath)
	if err != nil {
		log.Println(err)
		sentry.CaptureException(err)
	}
	<-done
}
