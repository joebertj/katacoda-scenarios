These are the modules that we need for this course
- scipy
- numpy
- matplotlib
- pandas
- sklearn

To check if we already have these go to Python 3 console:

`python3`{{execute}}

Run the following commands: 

`import sys`{{execute}}
`print('Python: {}'.format(sys.version))`{{execute}}

`import scipy`{{execute}}
`print('scipy: {}'.format(scipy.__version__))`{{execute}}

`import numpy`{{execute}}
`print('numpy: {}'.format(numpy.__version__))`{{execute}}

`import matplotlib`{{execute}}
`print('matplotlib: {}'.format(matplotlib.__version__))`{{execute}}

`import pandas`{{execute}}
`print('pandas: {}'.format(pandas.__version__))`{{execute}}

`import sklearn`{{execute}}
`print('sklearn: {}'.format(sklearn.__version__))`{{execute}}

Exit the console using `quit()`{{execute}}

For any missing modules install it using pip. e.g. if scipy is missing

`pip install scipy`{{execute}}

If pip is not installed, install it first:

`apt install -y pyton-pip`{{execute}}

So that we can view the images that we will generate later, we also need to spin up an http server like nginx.

`apt upgrade`{{execute}}

`apt install nginx`{{execute}}

`sed -i "s/root \/var\/www\/html;/root \/home\/scrapbook\/tutorial;/" /etc/nginx/sites-enabled/default`{{execute}}

`systemctl restart nginx`

Check if everything works by accessing nginx. 

http://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com](Test nginx)
