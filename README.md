flotilla
========

A collection of itty-bitty services.

The intention is that these random, disparate pieces can be assembled into something cool. (And maybe once I'm done writing nano services I'll actually get to that something!) The advantage of writing them independently is that they become much easier to reason about and be audited for resource usage, correctness, and security.

There is also some shared code that I factored out after it appeared in a few services. It's not really a library, as such, and breaking changes may be introduced without warning.


Cast of Characters
------------------

* [Share](http://github.com/casey/share) - permissive content addressed storage
* [WORM](http://github.com/casey/worm) - write-once key value storage
* [Timestamp](http://github.com/casey/timestamp) - timestamping
* [Static](http://github.com/casey/static) - static file serving
* [Okay](http://github.com/casey/ok) - always copacetic

They all:

* HAVE NO WARRANTY
* Are written in Go
* Run on App Engine
* Do one very simple thing
* Are self contained
* Try to be good REST citizens
* Have test instances running
* Require no authentication
* Should be abuse resistant
* Are released under a simple all-permissive license

The test instances are running on the GAE free tier, so feel free to try to break them. And let me know if you are able to! Be aware that they might disappear or lose all their data at any time.


Yet to be Written
-----------------

* Graph - A directed graph builder thingy
* Enumerate - A service for allocating sequential small numbers


Halp!
-----

I am sure that I did everything wrong. Suggestions, issues, and pull requests are all welcome. Please keep in mind that following conventions is not a priority for me.
