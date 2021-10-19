FRWD
====

"frwd" is a tiny TCP/IP port-forwarder written in Go. 

Useful when you want to forward a service but don't have root access or can't bother with iptables.


Install::

	go get github.com/avahidi/frwd/...


Usage::


    # forward TCP  0.0.0.0:8080 <--> 127.0.0.1:80
    frwd :8080 127.0.0.1:80
	
    # one-way forward UDP localhost:5300 -> 8.8.8.8:53
    frwd -u localhost:5300 8.8.8.8:53
