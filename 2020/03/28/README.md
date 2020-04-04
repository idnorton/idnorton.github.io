### SSH ProxyCommand for subnets

It's common in modern cloud environments to prevent ssh access to hosts
within the environment and only allow access to a single well maintained
bastion host with suitable protections against attack.

If you have multiple environments with non overlapping IP address
assignments then it would be neat to be able to configure the ssh client
to be able to use a different bastion host based on a CIDR match.

Let's take some examples:

| AWS Account | Bastion host                  | CIDR          |
| ----------- | ----------------------------- | ------------- |
| Account A   | bastion.account-a.example.com | 10.0.0.0/16   |
| Account B   | bastion.account-b.example.com | 10.1.0.0/16   |
| Account C   | bastion.account-c.example.com | 10.2.0.0/24   |
| Account D   | bastion.account-d.example.com | 10.3.0.0/25   |
| Account E   | bastion.account-e.example.com | 10.3.0.128/25 |

In the table above, we can use a traditional string match for the first
three:

```
Match 10.0.*
    ProxyCommand ssh -q bastion.account-a.example.com -W %h:%p

Match 10.1.*
    ProxyCommand ssh -q bastion.account-b.example.com -W %h:%p

Match 10.2.*
    ProxyCommand ssh -q bastion.account-c.example.com -W %h:%p
```

However that doesn't work for the last two. So how do we do this?
The answer is that we cannot achieve this with the a standard
configuration options.

This is something that's had questions online [https://serverfault.com/questions/803902/](https://serverfault.com/questions/803902/)

There's even a feature request [http://bugzilla.mindrot.org/show_bug.cgi?id=1169](http://bugzilla.mindrot.org/show_bug.cgi?id=1169)

There isn't currently a way to do this, however there is a neat feature in
`Match exec` which allows us to run a command and use the exit codes:

```
The exec keyword executes the specified command under the user's shell.  If the
command returns a zero exit status then the condition is considered true.
```

Docs at [https://linux.die.net/man/5/ssh_config](https://linux.die.net/man/5/ssh_config)

Having written some GoLang to do the network matching, we can now use this syntax:

```
  Match exec "~/bin/ssh-test-network --cidr=10.3.0.0/25 --host=%n
    ProxyCommand ssh -q bastion.account-d.example.com -W %h:%p

  Match exec "~/bin/ssh-test-network --cidr=10.3.0.128/25 --host=%n
    ProxyCommand ssh -q bastion.account-e.example.com -W %h:%p
```

Here's the command line options for the tool:

```
$ ./ssh-test-network --help
Usage of ./ssh-test-network:

  -cidr string
        CIDR network to test against
  -debug
        Enable debugging
  -host string
        Hostname to test

Use this command to route subnet ssh via a jumphost in your ssh config:

  Match exec "~/bin/ssh-test-network --cidr=10.0.0.0/8 --host=%n
    ProxyCommand ssh -q jumpbox.example.com -W %h:%p
    ForwardAgent yes
```

### Files

* [src/ssh-test-network (binary)](src/ssh-test-network)
* [src/ssh-test-network.go (source)](src/ssh-test-network.go)
