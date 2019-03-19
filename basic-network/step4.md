

## Let DNS servers resolve the IP of google.com

`ifconfig eth0`{{execute}}

The value displayed is one of the IP addresses of google.com

## List DNS Servers

`cat /etc/resolv.conf`{{execute}}

The IP address after the word nameserver is a designated DNS server and will answer to DNS requests.
