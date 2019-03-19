DNS servers resolves names to IP addresses. In Linux it is configured in a file. 

## Let DNS servers resolve the IP of google.com

`ping -c 1 google.com | head -1 | awk '{print $3}'`{{execute}}

The value displayed is one of the IP addresses of google.com

## List DNS Servers

`cat /etc/resolv.conf`{{execute}}

The IP address after the word nameserver is a designated DNS server and will answer to DNS requests.

8.8.8.8 is a DNS server owned by Google.
