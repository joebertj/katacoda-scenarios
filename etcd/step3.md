## Configure APT sources 

`add-apt-repository -y ppa:ondrej/php`{{execute}}

`add-apt-repository -y ppa:ondrej/apache2`{{execute}}

`apt update`{{execute}}

## Apache and PHP 7.1 and modules

`apt-get install -y php7.1 php7.1-opcache libapache2-mod-php7.1 php7.1-mysql php7.1-curl php7.1-json php7.1-gd php7.1-mcrypt php7.1-intl php7.1-mbstring php7.1-xml php7.1-zip php7.1-fpm php7.1-readline php7.1-sqlite3 php-pear php7.1-ldap php7.1-snmp php-db php-date php-xml php7.1-xml php-mysql php-rrd`{{execute}}

## Activate Apache PHP-FPM

`a2enmod proxy_fcgi setenvif proxy rewrite`{{execute}}

`a2enconf php7.1-fpm`{{execute}}

`a2dismod php7.1`{{execute}}

`systemctl restart apache2 php7.1-fpm`{{execute}}

## Verify Apache and PHP FPM is running

`ps -ef | grep apache2`{{execute}}

`ps -ef | grep php-fpm`{{execute}}

Access the URL to check:

http://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/

## Centreon Web

Return to root directory using `cd`{{execute}}

`wget http://files.download.centreon.com/public/centreon/centreon-web-18.10.4.tar.gz`{{execute}}

`tar xvzf centreon-web-18.10.4.tar.gz`{{execute}}

`cd centreon-web-18.10.4`{{execute}}

`./install.sh -i`{{execute}}

Press <kbd>ENTER</kbd> to continue. Use <kbd>SPACEBAR</kbd> to scroll down the license.

For the first 5 questions answer `y`{{execute}}. For PEAR use `/usr/share/php/PEAR.php`{{execute}} as value. For the rest, just use the defaul values by pressing <kbd>ENTER</kbd> and <kbd>y</kbd> when prompted. 

### Configure Timezone

`sed -i 's/;date.timezone =/date.timezone = Asia\/Manila/' /etc/php/7.1/fpm/php.ini`{{execute}}

`grep date.timezone /etc/php/7.1/fpm/php.ini`{{execute}}

`systemctl reload php7.1-fpm`{{execute}}

### Enable Centreon config

`ln -s /etc/apache2/conf-available/centreon.conf /etc/apache2/conf-enabled/`{{execute}}

`systemctl reload apache2`{{execute}}

Access the Centreon web interface:

http://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/centreon

For File or Directory 'not found' errors, use /usr/lib as a prefix for directories as what was configured in cmake. Supply the password for MariaDB earlier. Leave the non-required fields as blank. Just click Next until you finish.

Before logging in:

## Enable Centreon Components

`chown -R centreon: /usr/local/centreon`{{execute}}

`systemctl enable centengine`{{execute}}

`systemctl enable cbd`{{execute}}

`systemctl enable centcore`{{execute}}

`systemctl enable centreontrapd`{{execute}}

## Start Centreon Components

`systemctl start centengine`{{execute}}

`systemctl start cbd`{{execute}}

`systemctl start centcore`{{execute}}

`systemctl start centreontrapd`{{execute}}

### Verify Centreon Components

`systemctl status centengine`{{execute}}

`systemctl status cbd`{{execute}}

`systemctl status centcore`{{execute}}

`systemctl status centreontrapd`{{execute}}
