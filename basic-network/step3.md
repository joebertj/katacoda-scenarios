DNS servers resolves names to IP addresses. In Linux it is configured in a file. 

## Let DNS servers resolve the IP of google.com

`ping -c 1 google.com | head -1 | awk '{print $3}'`{{execute}}

The value displayed is one of the IPs of google.com

## List DNS Servers

`cat /etc/resolv.conf`{{execute}}


