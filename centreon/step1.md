It would be much easier if we install Centreon on CentOS as it is a supported distribution. Here we gonna attempt to install it to Ubuntu as it is the available environment in Katacoda.

## Get the Ubuntu version
`lsb_release -a`{{execute}}

When this tutorial was made it is 16.04. When the environment changes some adjustments should be made.

## Configure APT sources 

`add-apt-repository -y ppa:ondrej/php`{{execute}}

`add-apt-repository -y ppa:ondrej/apache2`{{execute}}

`apt update`{{execute}}

## Install Packages

### MariaDB

`apt-get install -y mariadb-server`{{execute}}

For Postfix Configuration leave the setting to Internet Site and press `TAB`{{execute}} to OK and press `ENTER`{{execute}}

### Apache and PHP 7.1 and modules

`apt-get install -y php7.1 php7.1-opcache libapache2-mod-php7.1 php7.1-mysql php7.1-curl php7.1-json php7.1-gd php7.1-mcrypt php7.1-intl php7.1-mbstring php7.1-xml php7.1-zip php7.1-fpm php7.1-readline php7.1-sqlite3 php-pear php7.1-ldap php7.1-snmp php-db php-date php7.1-xml`{{execute}}

#### Activate Apache PHP-FPM 

`a2enmod proxy_fcgi setenvif proxy rewrite`{{execute}}

`a2enconf php7.1-fpm`{{execute}}

`a2dismod php7.1`{{execute}}

`systemctl restart apache2 php7.1-fpm`{{execute}}

### RRDTool, mailx and PERL and modules

`apt-get install -y rrdtool bsd-mailx libconfig-inifiles-perl libcrypt-des-perl libdigest-hmac-perl libdigest-sha-perl libgd-perl`{{execute}}

### SNMP MIBS

`apt install -y snmp-mibs-downloader`{{execute}}

#### Verify MariaDB, Apache and PHP FPM is running

`ps -ef | grep mysqld`{{execute}}

`ps -ef | grep apache2`{{execute}}

`ps -ef | grep php-fpm`{{execute}}

## Adding centreon user

`groupadd -g 6000 centreon`{{execute}}

`useradd -u 6000 -g centreon -m -r -d /var/lib/centreon -c "Centreon Admin" -s /bin/bash centreon`{{execute}}

## Centreon Web

`http://files.download.centreon.com/public/centreon/centreon-web-18.10.4.tar.gz`{{execute}}

`tar xvzf centreon-web-18.10.4.tar.gz`{{execute}}

`cd centreon-web-18.10.4`{{execute}}

`./install.sh -i`{{execute}}

Press `ENTER`{{execute}} to continue. Use `SPACEBAR`{{execute}} to scroll down the license.

For the first 5 questions answer `y`{{execute}}. For the rest, just use the defaul values by pressing `ENTER`{{execute}} when prompted.
