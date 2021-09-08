This is a simple test of nginx. 

run the goServer on docker.

```dockerfile
docker run -p 2222:9999 -e APPID=2222 -d golang_nginx:latest
docker run -p 3333:9999 -e APPID=3333 -d golang_nginx:latest
docker run -p 4444:9999 -e APPID=4444 -d golang_nginx:latest
docker run -p 5555:9999 -e APPID=5555 -d golang_nginx:latest
```

The nginx plays a role as a reverse proxy.

