About go-run-as
===============
This program allow you run some program as some user.

Usage
=====

```
go-run-as <user> <program> [args...]
```

Example:
=========

```
[root@1f890e77ad08 go-run-as]# ./go-run-as ftp sh -c 'echo I am `whoami`'
I am ftp
```

Pre-Conditions
==============

You must have the privileges to run `<program>` as `<user>`. This program will NOT try to login.

So this program is almost only for the `root` user -- my aim is use it in docker.

Comparing With `sudo`
===================

`sudo -u <user> <program> [args...]` can also do the same thing. But `sudo` needs a tty. This program don't need it.

Comparing With `su`
===================

`su - <user> -c "<program> [args...]"` can also do the same thing. But it requires pass program and args in one string. 
If there are a lot of args or if there are special chars in args, it is difficult to pass them.



