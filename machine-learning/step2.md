The gateway server is most probably a NAT Gateway which will route our outgoing access to Internet 

## Display the routing table

`route`{{execute}}

## Filter and display the gateway column

`route | grep default | awk '{print $2}'`{{execute}}

The IP address displayed here is the gateway server
