FRWD
====

A tiny port-forwarder written in Go.


Usage::

    go build
    # forward TCP  0.0.0.0:8080 <--> 127.0.0.1:80
    ./frwd :8080 127.0.0.1:80
    # one-way forward UDP localhost:5300 -> 8.8.8.8:53
    ./frwd -u localhost:5300 8.8.8.8:53


Why?
----

To give you access to ports that would otherwise be hidden or
require some work to access.

I wrote this to emulate EXPOSE command of Docker in LXC. The alternative would
have been be a cryptic iptables command that would require root access.

Why not?
--------

As it is with most things in networking, you can probably already do it with netcat :(

License
-------

I throw a dart on the Wikipedia FOSS license page and it landed on zlib v0.7
