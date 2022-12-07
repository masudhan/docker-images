** If you want to use this image, do `docker pull chmadhus/web-server-3` **

OR it you want to tweak and make you own image follow below steps

## How to run

`docker build -t chmadhus/web-server-3:bangalore .`
`docker run -dit --name my-running-app -p 8080:80 chmadhus/web-server-3:bangalore`

Access from : http://localhost:8080

## To stop and remove the running container

`docker container rm -f my-running-app`

## To remove docker image

`docker rmi chmadhus/web-server-3:bangalore`


