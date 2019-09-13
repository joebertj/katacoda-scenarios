## Install MariaDB Package

`apt-get install -y mariadb-server`{{execute}}

## Verify MariaDB is running

`ps -ef | grep mysqld`{{execute}}

Set root password for Centreon Database:

`mysql -u root`{{execute}}

`grant all privileges on centreon.* to 'root'@'localhost' identified by 'centreon';`{{execute}}

`grant all privileges on centreon_storage.* to 'root'@'localhost' identified by 'centreon';`{{execute}}

`exit`{{execute}}

## Set open files limit

`sed -i 's/\[mysqld\]/\[mysqld\]\nopen_files_limit=32000/' /etc/mysql/mariadb.conf.d/50-server.cnf`{{execute}}

`killall mysqld`{{execute}}

## Verify if it is no longer running

`ps -ef | grep mysqld`{{execute}}

`systemctl restart mysql`{{execute}}

## Verify open files limit

`ps -ef | grep open-files-limit`{{execute}}

## RRDTool, mailx and PERL and modules

`apt-get install -y rrdtool bsd-mailx libconfig-inifiles-perl libcrypt-des-perl libdigest-hmac-perl libdigest-sha-perl libgd-perl libdbd-mysql-perl librrd-simple-perl`{{execute}}

For Postfix Configuration leave the setting to Internet Site and press <kbd>TAB</kbd> to OK and press <KBD>ENTER</kbd> twice.

## SNMP MIBS

`apt install -y snmp-mibs-downloader`{{execute}}

