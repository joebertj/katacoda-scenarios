## Update PEAR

`pear channel-update pear.php.net`{{execute}}

`pear upgrade-all`{{execute}}

## PHP dependencies

`apt install -y php-curl composer`{{execute}}

`cd /usr/local/centreon`{{execute}}

`sed -i '286s/continue/break/' /usr/share/php/Composer/DependencyResolver/RuleSetGenerator.php`{{execute}}

`composer install --no-dev --optimize-autoloader`{{execute}}

## Macro modification

`sed -i -e 's/_CENTREON_PATH_PLACEHOLDER_/centreon/g' /usr/local/centreon/www/index.html`{{execute}}

`sed -i -e 's/@PHP_BIN@/\/usr\/bin\/php/g' /usr/local/centreon/bin/centreon`{{execute}}

`sed -i -e 's/@PHP_BIN@/\/usr\/bin\/php/g' /usr/local/centreon/bin/export-mysql-indexes`{{execute}}

`sed -i -e 's/@PHP_BIN@/\/usr\/bin\/php/g' /usr/local/centreon/bin/generateSqlLite`{{execute}}

`sed -i -e 's/@PHP_BIN@/\/usr\/bin\/php/g' /usr/local/centreon/bin/import-mysql-indexes`{{execute}}

## Javascript dependencies

Return to root directory using `cd`{{execute}}

`cp centreon-web-18.10.4/package* /usr/local/centreon/`{{execute}}

`npm install`{{execute}}

`npm run build`{{execute}}

`npm prune --production`{{execute}}
