You can use any scripting language supported on the Host to write Plugins. Centreon accepts plugin formats the same as Nagios. The only requirement is that the script has a single line of output and an associated return value.

## Your first plugin

Create a file:

`vi my-plugin.sh`{{execute}}

You can then press `i`{{execute}} or `a`{{execute}} to enter into edit mode of vi. Type the following content:

`echo "OK - service is up"`{{execute}}

To save it just press <kbd>ESC</kbd> or `jj`{{execute}} which is premapped in this environment. This will return `vi` to command mode. Now, you can type `:wq`{{execute}} to save and exit from `vi`.

Now run your script:

`sh my-plugin.sh`{{execute}}

You can also add execute permission to your script:

`chmod +x my-plugin.sh`{{execute}}

You can now run it this way:

`./my-plugin.sh`{{execute}}

The output of your script will go under the **Status information** of Centreon.

### How to generate Status: OK, WARNING, CRITICAL, UNKNOWN

The **Status** of the script does not depent on the message that you placed on your script. Instead, the return value of your script is used. In Linux you can get a return value by using `echo $?`{{execute}}. Below are the values of the **Status**. 

- **OK** = 0
- **WARNING** = 1
- **CRITICAL** = 2
- **UNKNOWN** =3

Since you got `0` then the Status will be **OK**. It does not depend on the Status information.

To demonstrate return values run the following command:

`rm /`{{execute}} 

It throws an error. Now to see the return value use this:

`echo $?`{{execute}}

Observe that the value is no longer `0`.

Open your script again:

`vi my-plugin.sh`{{execute}}

Switch to edit mode by using `i`{{execute}} or `a`{{execute}}

Add the following line after `echo` line:

`exit 1`{{execute}}

It should look like this:

`echo "OK - service is up"
exit 1`

Switch to command mode by hitting <kbd>ESC</kbd> or `jj`{{execute}} and save and exit `:wq`{{execute}}

Now run your script again:

`./my-plugin.sh`{{execute}}

To check the return value use the command `echo $?`{{execute}}`.

Although it says **OK** the actual status would be **WARNING** because the return value is `1`.

### Performance data

Performance data is a feature of the plugin used to graph the values obtained if it is practical to do so. One use is case is if you need to count something and check if the value falls within a range then you can decide the **Status** from that. What you need to do is place a `|` and a `variable` and a `value` in the form of `variable=value` in the end of the output.

To demonstrate, open your script again:

`vi my-plugin.sh`{{execute}}

Switch to edit mode by using `i`{{execute}} or `a`{{execute}}

Add the following at the end of the `echo` line:

` | users=3`{{execute}}

It should look like this:

`echo "OK - service is up | users=3"
exit 1`

Switch to command mode by hitting <kbd>ESC</kbd> or `jj`{{execute}} and save and exit `:wq`{{execute}}

Now run your script again:

`./my-plugin.sh`{{execute}}

Let us apply what we have learned so far to count the number of logged in users in the system. Then we can change **Status** if it is more than `1`.

To do this, open your script again:

`vi my-plugin.sh`{{execute}}

Switch to edit mode by using `i`{{execute}} or `a`{{execute}}

Modify the script to look like this:

`status=0
users=$(who | wc -l)
if [ $users -gt 1 ]; then
    echo "CRITICAL - $users are logged in | users=$users"
    status = 2
else
    echo "OK - $users are logged in | users=$users"
fi
exit $status`

Switch to command mode by hitting <kbd>ESC</kbd> or `jj`{{execute}} and save and exit `:wq`{{execute}}

Now run your script again:

`./my-plugin.sh`{{execute}}

Check the return value again:

`echo $?`{{execute}}
