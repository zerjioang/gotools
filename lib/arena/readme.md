# Arena

> This package contains multiple implementations of Arena's in pure Go. The goal of them is to reduce GC pressoure and critical stop the world situations.
## Who is this for

This library may be interesting to you if you wish to reduce garbage collection (e.g. stop-the-world GC) performance issues in your golang programs, allowing you to switch to explicit byte array memory management techniques.

This can be useful, for example, for long-running server programs that manage lots of in-memory data items, such as caches and databases.