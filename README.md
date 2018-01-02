# Healer
Automatically stops or restarts docker containers that report themselves unhealthy.

## Why is this needed?
The ability to automatically restart failed containers is built in to Docker itself. But what if the container is still running but
doesn't work properly anymore?

If you're using Kubernetes, you can configure a health check and have Kubernetes restart unhealthy containers for you.

But if you're using AWS ECS, this behavior applies only to containers that are part of an ELB target group. Containers that sit behind this layer (in a typical microservice architecture) do not benefit from this.

## How to use
Healer uses Docker's built-in health functionality (see the [docs](https://docs.docker.com/engine/reference/commandline/run/)). When it sees that a container is unhealthy, it will either ignore the event, stop the container, or restart the container. The action taken depends on the value of the `HEALING_ACTION` environment variable set for the monitored container. `HEALING_ACTION` can be set to `NONE`, `STOP`, or `RESTART`.

## Example usage
Start Healer:
```
docker run -d --volume=/var/run/docker.sock:/tmp/docker.sock somarat/healer:latest  
```
Start monitored container:
```
docker run --health-interval=10s  -e "HEALING_ACTION=STOP" --health-cmd="curl -I --silent --fail localhost:8080 || exit 1" myimage:latest
```
