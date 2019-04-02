On Terminal Host 1 (master), run the sender Pod:

`kubectl run sender --rm -i --tty --image gcr.io/translate-watson/mcjoin`{{execute}}

Open New Terminal by clicking **+** beside Terminal Host 1. Run the receiver Pod:

`kubectl run receover --rm -i --tty --image gcr.io/translate-watson/mcjoin`{{execute}}

Run on sender: `/mcjoin/mcjoin -s`{{execute}}

Run on reciver: `/mcjoin/mcjoin`{{execute}}
