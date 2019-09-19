## Libraries

By default compilation generated a dynamically linked library. To verify run:
`file hello-world`{{execute}}

We can convert it to statically linked library by using the following command:
`CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"'`{{execute}}

We can now verify again:
`file hello-world`{{execute}}
