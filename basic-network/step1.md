There are several ways to list interfaces in Linux

## A classic way of doing

Run the command bellow

`ifconfig`{{execute}}

## A modern way of doing

`ip addr`{{execute}}

## More options

We can use `grep` and `awk` to further filter the output so we list the names only:

`ifconfig | grep -v -e "^ " -e ^$ | awk '{print $1}'`{{execute}}

`lo` is a loopback interface. It is the default interface and you don''t need a physical or virtual netword card for this. All machines have this address available for local access only. The IP address is 127.0.0.1 and the DNS alias is localhost.

`eth0` is either a physical or virtual network card which is automatically named by Linux. If you add another network card it would be probably named eth1.
