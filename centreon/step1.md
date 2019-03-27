It would be much easier if we install Centreon on CentOS as it is a supported distribution. Here we gonna attempt to install it to Ubuntu as it is the available environment in Katacoda.

## Get the Ubuntu version
`lsb_release -a`{{execute}}

When this tutorial was made it is 16.04. When the environment changes some adjustments should be made.

## Configure APT sources 

`add-apt-repository -y ppa:ondrej/php`{{execute}}
`add-apt-repository -y ppa:ondrej/apache2`{{execute}}
`apt update`{{execute}}

## Install Packages

### PHP 7.1 and modules

`apt-get install -y php7.1 php7.1-opcache libapache2-mod-php7.1 php7.1-mysql php7.1-curl php7.1-json php7.1-gd php7.1-mcrypt php7.1-intl php7.1-mbstring php7.1-xml php7.1-zip php7.1-fpm php7.1-readline php7.1-sqlite3 php-pear php7.1-ldap php7.1-snmp php-db php-date`{{execute}}

### PERL and modules

`apt-get install -y libconfig-inifiles-perl libcrypt-des-perl libdigest-hmac-perl libdigest-sha-perl libgd-perl`{{execute}}

### MariaDB and mailx

`apt-get install -y bsd-mailx mariadb-server`{{execute}}

For Postfix Configuration leave the setting to Internet Site and press `TAB`{{execute}} to OK and press `ENTER`{{execute}}
