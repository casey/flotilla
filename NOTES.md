Graph
-----

PUT /A/B -> increment the A -> B directed edge
PUT /A/A -> increment self-edge weight
GET /A/B -> get edge weight, last update time (eventually consistent)
GET /A   -> get total outgoing edge weight from A
GET /    -> get total weight of all edges

This guy will probably need a background task queue. It may seem wierd, but I need it for a service which allows users to edit documents, where it would be nice to track the heritage of documents, like which documents they were forked from. This seems like the minimum needed to support such functionality.

Counter
-------

POST /KEY -> increment counter for KEY, set last access time
GET  /KEY -> get current count, last access time, eventually consistent
GET  /    -> get global counter value

You may also refer to this service as "The Count", or "Count von Count", and imagine that it is a purple muppet. You can probably make the counts overflow, in which case who knows what will happen. The Count is not good with negative numbers.

Enumerate
---------

GET /KEY/VALUE -> get number for VALUE in the KEY namespace
GET /KEY/NUMBER -> get value for NUMBER in the KEY namespace
PUT /KEY/VALUE -> establishes number for VALUE in the KEY namespace, as small as possible, start at 0

Just used to establish small, nice to look at integers for things like users, published pieces of content, etc. Think vimeo video ids.
