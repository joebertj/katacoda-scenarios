It would be much easier if we install Centreon on CentOS as it is a supported distribution. Here we gonna attempt to install it to Ubuntu as it is the available environment in Katacoda.

## Get the Ubuntu version
`lsb_release -a`{{execute}}

When this tutorial was made it is 16.04. When the environment changes some adjustments should be made.

# Centreon Engine

## Install Clib

`apt update`{{execute}}

`apt install -y cmake`{{execute}}

`wget http://files.download.centreon.com/public/centreon-clib/centreon-clib-18.10.0.tar.gz`{{execute}}

`tar xvzf centreon-clib-18.10.0.tar.gz`{{execute}}

`cd centreon-clib-18.10.0/build`{{execute}}

`cmake \
   -DWITH_TESTING=0 \
   -DWITH_PREFIX=/usr \
   -DWITH_PREFIX_LIB=/usr/lib \
   -DWITH_PREFIX_INC=/usr/include/centreon-clib \
   -DWITH_SHARED_LIB=1 \
   -DWITH_STATIC_LIB=0 \
   -DWITH_PKGCONFIG_DIR=/usr/lib/pkgconfig .`{{execute}}

`make`{{execute}}

`make install`{{execute}}

Return to root directory using `cd`{{execute}}

## Adding centreon-engine user

`groupadd centreon-engine`{{execute}}

`useradd -g centreon-engine -m -r -d /var/lib/centreon-engine centreon-engine`{{execute}}

## Download and Install

`wget http://files.download.centreon.com/public/centreon-engine/centreon-engine-18.10.0.tar.gz`{{execute}}

`tar xvzf centreon-engine-18.10.0.tar.gz`{{execute}}

`cd centreon-engine-18.10.0/build`{{execute}}

`cmake \
   -DWITH_PREFIX=/usr \
   -DWITH_PREFIX_BIN=/usr/sbin \
   -DWITH_PREFIX_CONF=/etc/centreon-engine \
   -DWITH_PREFIX_LIB=/usr/lib/centreon-engine \
   -DWITH_USER=centreon-engine \
   -DWITH_GROUP=centreon-engine \
   -DWITH_LOGROTATE_SCRIPT=1 \
   -DWITH_VAR_DIR=/var/log/centreon-engine \
   -DWITH_RW_DIR=/var/lib/centreon-engine/rw \
   -DWITH_STARTUP_DIR=/etc/init.d \
   -DWITH_PKGCONFIG_SCRIPT=1 \
   -DWITH_PKGCONFIG_DIR=/usr/lib/pkgconfig \
   -DWITH_TESTING=0`{{execute}}

`make`{{execute}}

`make install`{{execute}}

Return to root directory using `cd`{{execute}}

# Centreon Broker

## Adding centreon-broker user

`groupadd centreon-broker`{{execute}}

`useradd -g centreon-broker -m -r -d /var/lib/centreon-broker centreon-broker`{{execute}}

## Download and Install

`apt install -y libqt4-dev libqt4-sql-mysql zlib1g-dev librrd-dev liblua5.2-dev libgnutls28-dev`{{execute}}

`wget http://files.download.centreon.com/public/centreon-broker/centreon-broker-18.10.1.tar.gz`{{execute}}

`tar xvzf centreon-broker-18.10.1.tar.gz`{{execute}}

`cd centreon-broker-18.10.1/build`{{execute}}

`cmake \
    -DWITH_DAEMONS='central-broker;central-rrd' \
    -DWITH_GROUP=centreon-broker \
    -DWITH_PREFIX=/usr \
    -DWITH_PREFIX_BIN=/usr/sbin \
    -DWITH_PREFIX_CONF=/etc/centreon-broker \
    -DWITH_PREFIX_INC=/usr/include/centreon-broker \
    -DWITH_PREFIX_LIB=/usr/lib/nagios \
    -DWITH_PREFIX_MODULES=/usr/share/centreon/lib/centreon-broker \
    -DWITH_STARTUP_DIR=/etc/init.d \
    -DWITH_STARTUP_SCRIPT=auto \
    -DWITH_TESTING=0 \
    -DWITH_USER=centreon-broker .`{{execute}}

`make`{{execute}}

`make install`{{execute}}

Return to root directory using `cd`{{execute}}

## Centreon Connector

`apt install -y libperl-dev libssh2-1-dev`{{execute}}

`wget http://files.download.centreon.com/public/centreon-connectors/centreon-connectors-18.10.0.tar.gz`{{execute}}

`tar xvzf centreon-connectors-18.10.0.tar.gz`{{execute}}

`cd centreon-connector-18.10.0/perl/build`{{execute}}

`cmake \
   -DWITH_PREFIX=/usr \
   -DWITH_PREFIX_BINARY=/usr/lib/centreon-connector \
   -DWITH_TESTING=0 .`{{execute}}

`make`{{execute}}

`make install`{{execute}}

Return to root directory using `cd`{{execute}}

`cd centreon-connector-18.10.0/ssh/build`{{execute}}

`cmake \
   -DWITH_PREFIX=/usr \
   -DWITH_PREFIX_BINARY=/usr/lib/centreon-connector \
   -DWITH_TESTING=0 .`{{execute}}

`make`{{execute}}

`make install`{{execute}}

Return to root directory using `cd`{{execute}}

# Centreon Web

## Configure APT sources 

`add-apt-repository -y ppa:ondrej/php`{{execute}}

`add-apt-repository -y ppa:ondrej/apache2`{{execute}}

`apt update`{{execute}}

## Install Packages

### MariaDB

`apt-get install -y mariadb-server`{{execute}}

#### Verify MariaDB is running

`ps -ef | grep mysqld`{{execute}}

Set root password for Centreon Database:

`mysql -u root`{{execute}}

`grant all privileges on centreon.* to 'root'@'localhost' identified by 'centreon';`{{execute}}

`grant all privileges on centreon_storage.* to 'root'@'localhost' identified by 'centreon';`{{execute}}

`exit`{{execute}}

Set open files limit:

`sed -i 's/\[mysqld\]/\[mysqld\]\nopen_files_limit=32000/' /etc/mysql/mariadb.conf.d/50-server.cnf`{{execute}}

`systemctl restart mysql`{{execute}}

### Verify open files limit

`ps -ef | grep open-files-limit`{{execute}}

### Apache and PHP 7.1 and modules

`apt-get install -y php7.1 php7.1-opcache libapache2-mod-php7.1 php7.1-mysql php7.1-curl php7.1-json php7.1-gd php7.1-mcrypt php7.1-intl php7.1-mbstring php7.1-xml php7.1-zip php7.1-fpm php7.1-readline php7.1-sqlite3 php-pear php7.1-ldap php7.1-snmp php-db php-date php-xml php7.1-xml`{{execute}}

### Update PEAR

`pear channel-update pear.php.net`{{execute}}

`pear upgrade-all`{{execute}}

#### Activate Apache PHP-FPM 

`a2enmod proxy_fcgi setenvif proxy rewrite`{{execute}}

`a2enconf php7.1-fpm`{{execute}}

`a2dismod php7.1`{{execute}}

`systemctl restart apache2 php7.1-fpm`{{execute}}

#### Verify Apache and PHP FPM is running

`ps -ef | grep apache2`{{execute}}

`ps -ef | grep php-fpm`{{execute}}

Access the URL to check:

http://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/

### RRDTool, mailx and PERL and modules

`apt-get install -y rrdtool bsd-mailx libconfig-inifiles-perl libcrypt-des-perl libdigest-hmac-perl libdigest-sha-perl libgd-perl`{{execute}}

For Postfix Configuration leave the setting to Internet Site and press <kbd>TAB</kbd> to OK and press <kbd>ENTER</kbd>``{{execute}} twice. 

### SNMP MIBS

`apt install -y snmp-mibs-downloader`{{execute}}

## Adding centreon user

`groupadd -g 6000 centreon`{{execute}}

`useradd -u 6000 -g centreon -m -r -d /var/lib/centreon -c "Centreon Admin" -s /bin/bash centreon`{{execute}}

## Download and Install

`wget http://files.download.centreon.com/public/centreon/centreon-web-18.10.4.tar.gz`{{execute}}

`tar xvzf centreon-web-18.10.4.tar.gz`{{execute}}

`cd centreon-web-18.10.4`{{execute}}

`./install.sh -i`{{execute}}

Press <kbd>ENTER</kbd>``{{execute}} to continue. Use <kbd>SPACEBAR</kbd> to scroll down the license.

For the first 5 questions answer `y`{{execute}}. For PEAR use `/usr/share/php/PEAR.php`{{execute}} as value. For the rest, just use the defaul values by pressing <kbd>ENTER</kbd> and <kbd>y</kbd> when prompted. 

### PHP dependencies

`apt install -y composer`{{execute}}

`cd /usr/lib/centreon`{{execute}}

`composer install --no-dev --optimize-autoloader`{{execute}}

### Macro modification

`sed -i -e 's/_CENTREON_PATH_PLACEHOLDER_/centreon/g' /usr/share/centreon/www/index.html`{{execute}}

`sed -i -e 's/@PHP_BIN@/\/usr\/bin\/php/g' /usr/share/centreon/bin/centreon`{{execute}}

`sed -i -e 's/@PHP_BIN@/\/usr\/bin\/php/g' /usr/share/centreon/bin/export-mysql-indexes`{{execute}}

`sed -i -e 's/@PHP_BIN@/\/usr\/bin\/php/g' /usr/share/centreon/bin/generateSqlLite`{{execute}}

`sed -i -e 's/@PHP_BIN@/\/usr\/bin\/php/g' /usr/share/centreon/bin/import-mysql-indexes`{{execute}}

### Javascript dependencies

Return to root directory using `cd`{{execute}}

`cp centreon-web-18.10.4/package* /usr/lib/centreon/`{{execute}}

`npm install`{{execute}}

`npm run build`{{execute}}

`npm prune --production`{{execute}}

### Configure Timezone

`sed -i 's/;date.timezone =/date.timezone = Asia\/Manila/' /etc/php/7.1/fpm/php.ini`{{execute}}

`grep date.timezone /etc/php/7.1/fpm/php.ini`{{execute}}

`systemctl reload php7.1-fpm`{{execute}}

### Create the Centreon Plugins and Broker Log directories

`mkdir -p /usr/lib/centreon/plugins`{{execute}}

`chown centreon: /usr/lib/centreon/plugins`{{execute}}

`mkdir /var/log/centreon-broker`{{execute}}

`chown centreon-broker: /var/log/centreon-broker`{{execute}}

### Enable Centreon config

`ln -s /etc/apache2/conf-available/centreon.conf /etc/apache2/conf-enabled/`{{execute}}

`systemctl reload apache2`{{execute}}

Access the Centreon web interface:

http://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/centreon

For File or Directory 'not found' errors, use /usr/lib as a prefix for directories as what was configured in cmake. Supply the password for MariaDB earlier. Leave the non-required fields as blank. Just click Next until you finish.
