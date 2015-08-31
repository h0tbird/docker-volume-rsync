# docker-volume-rsync

- An rsync docker volume plugin.
- This is my first Go program.
- I am just learning Go and having fun with Docker.
- I have not even a clear idea of what I'll do ;-)

I run this thing like this in one console:
```
go get github.com/h0tbird/docker-volume-rsync
go install github.com/h0tbird/docker-volume-rsync
sudo docker-volume-rsync
```

And then I run docker like this in another console:
```
docker run -it --volume-driver rsync -v src.host.org/foo:/foo alpine sh
```

[Docker Plugins](https://github.com/docker/docker/blob/master/docs/extend/index.md)
