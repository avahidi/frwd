
FRWD
====

A tiny port-forwarder written in golang.


Usage::

    # build and forward from high-http to http on local machine
    go build
    ./frwd :8080 127.0.0.1:80




Why
---

Frwd can give you access to ports that would otherwise be hidden or require
significant work to access.

For example, unlike Docker LXD does not have an EXPOSE command.
So to expose ports you would need to set iptables rules.
Frwd can achieve similar results without needing root access or requiring you to
remember that odd iptables syntax.



Why not?
--------

As it is with most things in networking, you can already do it with netcat::

    netcat -L 127.0.0.1:80 -p 8080

:(


License
-------

I throw a dart on the Wikipedia FOSS license page and it landed on zlib v0.7
