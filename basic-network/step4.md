Since we are using a non-root terminal some of the commands will no longer work. You can test it on an actual Linux machine.

## Manually setting an IP address of interface eth0 to 192.168.1.1

`ifconfig eth0 192.168.1.1`{{execute}}

or

`ip addr add 192.168.1.1 dev eth0`{{execute}}
