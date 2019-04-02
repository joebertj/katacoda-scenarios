On Terminal Host 1 (master), run the receiver Pod:

`kubectl run receiver --rm -i --tty --image gcr.io/translate-watson/mcjoin`{{execute}}

Wait for root prompt and run: `/mcjoin -d`{{execute}}

Open New Terminal by clicking **+** beside Terminal Host 1. Run the sender Pod:

`kubectl run sender --rm -i --tty --image gcr.io/translate-watson/mcjoin`{{execute}}

Wait for root prompt and run: `/mcjoin -s -d`{{execute}}

Go back to Terminal Host 1.

You can hit <kbd>CTRL-C</kbd> to stop receiving or sending.

Observe that packets flow from sender to receiver.

By default `mcjoin` runs recevier mode. `-d` is debug mode. `-s` is sender mode. Use `-h` to display more options.

To leave/join every 5 secs you can use `/mcjoin -d -r 5`{{execute}}

To send 10 packets you can use `/mcjoin -s -d -c 10`{{execute}}

Type `exit`{{execute}} on each terminal when you are done.
