# docker-volume-rsync

- An rsync docker volume plugin just for the fun of it.
- This is my first Go program (I am learning).
- I don't have a master plan for this project.

#### How does it work
The volume driver parses the `--volume` flag wich has the format `<remote host>/<remote path>:<container path>`. The driver synchronizes (pulls) the data so it can be bind-mounted to the container. When the container is terminated the driver synchronizes (pushes) the data back to the remote host.

#### Install
```
go get github.com/h0tbird/docker-volume-rsync
go install github.com/h0tbird/docker-volume-rsync
```

#### Run the driver
```
sudo docker-volume-rsync \
--archive \
--compress \
--delete \
--ssh-private-key /root/.ssh/data.key
```

#### Run a container
```
docker run -it --volume-driver rsync -v remote.host.org/foo:/foo alpine sh
```

#### Acknowledgments
I learn by copying and pasting someone else's code and adjusting it to my needs (stackoverflow also helps). This wouldn't have been possible without the code I borrowed from David Calavera and Matthias Kadenbach so thank you!

#### Devel stuff:
[Docker Plugins](https://github.com/docker/docker/blob/master/docs/extend/index.md)
