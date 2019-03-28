You can use any scripting language supported on the Host to write Plugins. Centreon accepts plugin formats the same as Nagios. The only requirement is that the script has a single line of output and an associated return value.

## Your first plugin

Create a file:

`vi my-plugin.sh`{{execute}}

You can then press `i`{{execute}} or `a`{{execute}} to enter into edit mode of vi. Type the following content:

`echo "OK - service is up"`{{execute}}

To save it just press <kbd>ESC</kbd> or `jj`{{execute}} which is premapped in this environment. This will return vi to command mode. Now, you can type `:wq`{{execute}} to save and exit from vi.

Now run your script:

`sh my-plugin.sh`

The output of your script will go under the **Status information** of Centreon

## How to generate Status: OK, WARNING, CRITICAL, UNKNOWN

The **Status** of the script does not depent on the message that you placed on your script. The return value of your script is used. In linux you can get a return value by using `$?`{{execute}}. Below are the values of the status. 

- OK = 0
- WARNING = 1
- CRITICAL = 2
- UNKNOWN =3

Since you got 0 then the Status will be OK. It does not depend on the Status information.
