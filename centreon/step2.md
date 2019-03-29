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

`killall mysqld`{{execute}}

`systemctl restart mysql`{{execute}}

### Verify open files limit

`ps -ef | grep open-files-limit`{{execute}}

### Apache and PHP 7.1 and modules

`apt-get install -y php7.1 php7.1-opcache libapache2-mod-php7.1 php7.1-mysql php7.1-curl php7.1-json php7.1-gd php7.1-mcrypt php7.1-intl php7.1-mbstring php7.1-xml php7.1-zip php7.1-fpm php7.1-readline php7.1-sqlite3 php-pear php7.1-ldap php7.1-snmp php-db php-date php-xml php7.1-xml`{{execute}}

### Activate Apache PHP-FPM

`a2enmod proxy_fcgi setenvif proxy rewrite`{{execute}}

`a2enconf php7.1-fpm`{{execute}}

`a2dismod php7.1`{{execute}}

`systemctl restart apache2 php7.1-fpm`{{execute}}

### Verify Apache and PHP FPM is running

`ps -ef | grep apache2`{{execute}}

`ps -ef | grep php-fpm`{{execute}}

Access the URL to check:

http://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/

### RRDTool, mailx and PERL and modules

`apt-get install -y rrdtool bsd-mailx libconfig-inifiles-perl libcrypt-des-perl libdigest-hmac-perl libdigest-sha-perl libgd-perl`{{execute}}

For Postfix Configuration leave the setting to Internet Site and press <kbd>TAB</kbd> to OK and press <KBD>ENTER</kbd> twice.

### SNMP MIBS

`apt install -y snmp-mibs-downloader`{{execute}}

## Download and Install

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

## Enable Centreon Components

`systemctl enable centcore`{{execute}}

`systemctl enable centreontrapd`{{execute}}

## Start Centreon Components

`systemctl start centcore`{{execute}}

`systemctl start centreontrapd`{{execute}}

### Verify Centreon Components

`systemctl status centcore`{{execute}}

`systemctl status centreontrapd`{{execute}}

### Manual start

`/usr/bin/perl /usr/local/centreon/bin/centcore --logfile=/var/log/centreon/centcore.log --severity=error --config=/etc/centreon/conf.pm`{{execute}}

`/usr/bin/perl /usr/local/centreon/bin/centreontrapd --logfile=/var/log/centreon/centreontrapd.log --severity=error --config=/etc/centreon/conf.pm --config-extra=/etc/centreon/centreontrapd.pm`{{execute}}

Access the Centreon web interface:

http://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/centreon

For File or Directory 'not found' errors, use /usr/lib as a prefix for directories as what was configured in cmake. Supply the password for MariaDB earlier. Leave the non-required fields as blank. Just click Next until you finish.

## Optional Steps

### Update PEAR

`pear channel-update pear.php.net`{{execute}}

`pear upgrade-all`{{execute}}

### PHP dependencies

`apt install -y php-curl composer`{{execute}}

`cd /usr/local/centreon`{{execute}}

`sed -i '286s/continue/break/' /usr/share/php/Composer/DependencyResolver/RuleSetGenerator.php`{{execute}}

`composer install --no-dev --optimize-autoloader`{{execute}}

### Macro modification

`sed -i -e 's/_CENTREON_PATH_PLACEHOLDER_/centreon/g' /usr/local/centreon/www/index.html`{{execute}}

`sed -i -e 's/@PHP_BIN@/\/usr\/bin\/php/g' /usr/local/centreon/bin/centreon`{{execute}}

`sed -i -e 's/@PHP_BIN@/\/usr\/bin\/php/g' /usr/local/centreon/bin/export-mysql-indexes`{{execute}}

`sed -i -e 's/@PHP_BIN@/\/usr\/bin\/php/g' /usr/local/centreon/bin/generateSqlLite`{{execute}}

`sed -i -e 's/@PHP_BIN@/\/usr\/bin\/php/g' /usr/local/centreon/bin/import-mysql-indexes`{{execute}}

### Javascript dependencies

Return to root directory using `cd`{{execute}}

`cp centreon-web-18.10.4/package* /usr/local/centreon/`{{execute}}

`npm install`{{execute}}

`npm run build`{{execute}}

`npm prune --production`{{execute}}

