Graph
-----

* PUT /A/B -> increment the A -> B directed edge
* PUT /A/A -> increment A->A self-edge weight
* GET /A/B -> get edge weight, last update time (eventually consistent)
* GET /A   -> get total outgoing edge weight from A
* GET /    -> get total weight of all edges

This guy will probably need a background task queue to avoid datastore write contention. It might seem weird like a weird service, but I need something to track document revision history, and this seems like the minimum needed to support that.

Enumerate
---------

* GET /KEY/VALUE -> get number for VALUE in the KEY namespace
* GET /KEY/NUMBER -> get value for NUMBER in the KEY namespace
* PUT /KEY/VALUE -> establishes number for VALUE in the KEY namespace, as small as possible, start at 0

Just used to establish small, nice to look at integers for things like users, published pieces of content, etc. Think vimeo video ids.
