On Terminal Host 1 (master), run the sender Pod:

`kubectl run sender --rm -i --tty --image gcr.io/translate-watson/mcjoin`{{execute}}

Open New Terminal by clicking **+** beside Terminal Host 1. Run the receiver Pod:

`kubectl run receiver --rm -i --tty --image gcr.io/translate-watson/mcjoin`{{execute}}

Run on sender: `/mcjoin/mcjoin -s -d`{{execute}}

Run on receiver: `/mcjoin/mcjoin -d`{{execute}}

When you confirm that packets flow from sender to receiver you can hit <kdb>CTRL-C</kbd> and type exit on each terminal.
