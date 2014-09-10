flotilla
========

A collection of itty-bitty services.

The intention is that these random, disparate pieces can be assembled into something cool. (And maybe once I'm done writing nano services I'll actually get to that something!) The advantage of writing them independantly is that they become much easier to reason about and be audited for resource usage, correctness, and security.


Cast of Characters
------------------

* [WORM - write-once key value storage](github.com/casey/worm)
* [Timestamp - timestamp service](github.com/casey/timestamp)
* [Static - static file serving](github.com/casey/static)
* [Okay - always OK](github.com/casey/ok)

The test instances are all running on the GAE free tier, so feel free to try to hose them. Be aware that they might disappear or lose all their data at any time.

They all:

* HAVE NO WARRANTY
* Do one very simple thing
* Are self contained
* Run on App Engine
* Are written in Go
* Try to be good REST citizens
* Have test instances running
* Require no authentication
* Should be resistant to abuse (SHOULD be)
* Are released under an all-permissive license


Yet to be Written
-----------------

* Publish - Content-addressable storage service
* Ephemera - Like WORM but allows overwriting, possibly with versioning
* Graph - A directed graph builder thingy
* Counter - A counter service. Think for hit counting, but might actually be a subset of Graph
* Enumerate - A service for assigning small numbers from a sequence. Could be used to assign small unique numbers to user accounts or pieces of published content


To Do
----

* More tests! Find edge cases!
* Add CORS headers so they can all be used from the browser
* Figure out if they're all using the right caching strategy
* Write some simple example consumers. Ideas include a pastebin service, a link shortener, a URL hit counter, and more.


Halp!
-----

I am sure that I made lots of mistakes. Suggestions, issues, and pull requests are all welcome. My coding style is probably a little weird, so sorry about that.
