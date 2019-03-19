We already know that the DNS Servers are configured in `/etc/resolv.conf`

## Add 1.1.1.1 as a new DNS server

1.1.1.1 is a DNS server owned by Cloudflare.

`echo "nameserver 1.1.1.1" >> /etc/resolv.conf`{{execute}}

The above command will add 1.1.1.1 as a DNS server. 
