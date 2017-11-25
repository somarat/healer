# Healer
Automatically heal docker containers that report themselves unhealthy

## Notes
  docker run -it --health-interval=1s  -e "HEALING_ACTION=STOP" --health-cmd="curl -I --silent --fail https://www.google.com || exit 1" centos:7
  
docker run -d --volume=/var/run/docker.sock:/tmp/docker.sock healer  

