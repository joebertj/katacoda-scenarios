It would be much easier if we install Centreon on CentOS as it is a supported distribution. Here we gonna attempt to install it to Ubuntu as it is the available environment in Katacoda.

## Configure vi

Map `jj` to <kbd>ESC</kbd> button

`echo "inoremap jj <Esc>" > .vimrc`{{execute}}

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

## Adding centreon user

`groupadd -g 6000 centreon`{{execute}}

`useradd -u 6000 -g centreon -m -r -d /var/lib/centreon -c "Centreon Admin" -s /bin/bash centreon`{{execute}}

### Create the Centreon Plugins and Log directories

`mkdir -p /usr/lib/centreon/plugins`{{execute}}

`chown centreon: /usr/lib/centreon/plugins`{{execute}}

`mkdir /var/log/centreon`{{execute}}

`chown centreon: /var/log/centreon`{{execute}}

`mkdir /var/log/centreon-broker`{{execute}}

`chown centreon-broker: /var/log/centreon-broker`{{execute}}

## Configure systemd

`vi /etc/systemd/system/centengine.service`{{execute}}

`i`{{execute}} or `a`{{execute}} to enter edit mode

```
[Unit]
Description=Centreon Engine

[Service]
ExecStart=/usr/sbin/centengine /etc/centreon-engine/centengine.cfg
ExecReload=/bin/kill -HUP $MAINPID
Type=simple
User=centreon-engine

[Install]
WantedBy=multi-user.target
```{{execute}}

<kbd>ESC</kbd> or `jj`{{execute}} to enter command mode

`:wq`{{execute}}

## Enable Centreon Components

`systemctl enable centengine`{{execute}}

`systemctl enable cbd`{{execute}}

## Start Centreon Components

`systemctl start centengine`{{execute}}

`systemctl start cbd`{{execute}}

### Verify Centreon Components

`systemctl status centengine`{{execute}}

`systemctl status cbd`{{execute}}
