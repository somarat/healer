package main

import (
  "log"
  "regexp"
  "strings"
  "os"
  docker "github.com/fsouza/go-dockerclient"
)

var regex, err = regexp.Compile(`^HEALING_ACTION=(STOP|RESTART|NONE)$`)

func getHealingAction(env []string) string {
  for _, element := range env {
    match := regex.FindStringSubmatch(element)
    if len(match) == 2 {
      return match[1]
    }
  }
  return "NONE"
}

func handleEvent(client *docker.Client, event *docker.APIEvents) {
  if event.Type == "container" && strings.Contains(event.Action, "unhealthy") {
    log.Println(event.Status)
    log.Printf("Container %s marked unhealthy\n", event.Actor.ID)
    container, err := client.InspectContainer(event.Actor.ID);
    if err == nil {
      switch getHealingAction(container.Config.Env) {
      case "STOP":
        client.StopContainer(event.Actor.ID, 10)
        log.Printf("Stopped container %s\n", event.Actor.ID)
      case "RESTART":
        client.RestartContainer(event.Actor.ID, 10)
        log.Printf("Restarted container %s\n", event.Actor.ID)
      default:
        log.Printf("Leaving container %s alone\n", event.Actor.ID)
      }
    }
  }
}

func main() {
  dockerHost := os.Getenv("DOCKER_HOST")
  if dockerHost == "" {
    os.Setenv("DOCKER_HOST", "unix:///tmp/docker.sock")
  }
	client, err := docker.NewClientFromEnv()
	if err != nil {
		panic(err)
	}

  listener := make(chan *docker.APIEvents)
  err = client.AddEventListener(listener)
  if err != nil {
      log.Fatal(err)
  }

  log.Println("Monitoring container health")

  for {
    select {
      case event := <-listener:
        go handleEvent(client, event)
    }
  }
}
