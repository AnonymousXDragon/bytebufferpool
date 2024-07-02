# ByteBufferPool with Anti-Memory Waste Protection

* anti-memory waste protection ensures that pool doesn't grow beyond the given size and waste memory
* bytebufferpool used to manage and reuse the buffer efficiently , and doesn't allow buffer to grow bigger
* sync/pool
  * we can reuse bytebuffers instead of creating one each time when required
