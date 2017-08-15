
FRWD
====

A tiny port-forwarder written in golang.


Usage::

    # build and forward from high-http to http on local machine
    go build
    ./frwd :8080 127.0.0.1:80


Why?
----

To give you access to ports that would otherwise be hidden or require
some work to access.

I wrote this when moving containers from Docker to LXD and noticing that
LXD was missing an EXPOSE command. The alternative to using frwd would
have been be a cryptic iptables command that would require root access.


Why not?
--------

As it is with most things in networking, you can already do it with netcat::

    netcat -L 127.0.0.1:80 -p 8080

:(


License
-------

I throw a dart on the Wikipedia FOSS license page and it landed on zlib v0.7
