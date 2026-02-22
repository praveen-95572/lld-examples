# LRU Cache – Low Level Design (LLD)

## 1. Problem Statement

Design a **Least Recently Used (LRU) Cache** that:

- Stores key-value pairs in memory
- Has a fixed capacity
- Evicts the least recently used item when capacity is exceeded
- Supports O(1) get and set operations
- Thread-safe for concurrent access

---

## 2. Functional Requirements

1. `Get(key)` – returns value if exists; else returns -1 or nil
2. `Put(key, value)` – insert or update key-value
3. Evict **least recently used** key when capacity is full
4. Optional: support TTL for each key

---

## 3. Non-Functional Requirements

- Thread-safe
- High performance (O(1) for get/put)
- Modular and maintainable
- Easy to extend (e.g., distributed caching later)

---

## 4. Core Entities

### LRUCache
- Capacity: maximum number of items
- HashMap: key → node (O(1) lookup)
- Doubly Linked List: stores keys in usage order
  - Head → most recently used
  - Tail → least recently used

### Node
- Key
- Value
- Pointers to previous and next nodes

---

## 5. Data Structures

| Component | Data Structure | Purpose |
|-----------|----------------|---------|
| Cache     | map[key]*Node  | O(1) lookup |
| Usage     | Doubly Linked List | Track LRU order |
| Node      | struct         | Stores key & value |

---

## 6. LRU Cache Flow

### Get(key)
1. Lookup key in hashmap
2. If found:
   - Move node to **head** of linked list
   - Return value
3. If not found:
   - Return -1 / nil

### Put(key, value)
1. Lookup key
2. If exists:
   - Update value
   - Move node to **head**
3. If not exists:
   - If capacity reached:
     - Remove **tail node**
     - Delete key from hashmap
   - Insert new node at **head**
   - Add to hashmap

---

## 7. Thread Safety

- Use `sync.RWMutex` to:
  - Allow concurrent reads
  - Prevent race conditions during writes

---

## 8. Time Complexity

| Operation | Complexity |
|-----------|------------|
| Get       | O(1)      |
| Put       | O(1)      |

---

## 9. Future Enhancements

- TTL (time-to-live) support
- Persistence to disk or Redis
- Distributed cache (e.g., using consistent hashing)
- Metrics / monitoring
- Generic types using Go 1.18+ generics