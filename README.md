# Healer
Automatically stops or restarts docker containers that report themselves unhealthy.

## Why is this needed?
The ability to automatically restart failed containers is built in to Docker itself. But what if the container is still running, but
doesn't work properly anymore?
If you're using Kubernetes, you can configure a health check and have Kubernetes restart unhealthy containers for you.
But if you're using AWS ECS, this behavior applies only to containers that are part of an ELB target group. Containers that sit behind this layer (in a typical microservice architecture) do not benefit from this.

## How to use


## Notes
  docker run -it --health-interval=1s  -e "HEALING_ACTION=STOP" --health-cmd="curl -I --silent --fail https://www.google.com || exit 1" centos:7
  
docker run -d --volume=/var/run/docker.sock:/tmp/docker.sock healer  

