# Task Management System – Low Level Design (LLD)

## 1. Problem Statement

Design a Task Management System similar to Jira/Trello that allows:

- Creating tasks
- Assigning tasks to users
- Updating task status
- Adding comments
- Managing priorities and due dates
- Retrieving tasks efficiently

---

## 2. Functional Requirements

### Core Features

1. Create a task
2. Update a task
3. Delete a task
4. Assign task to a user
5. Change task status (TODO, IN_PROGRESS, DONE)
6. Add comments to task
7. Set priority (LOW, MEDIUM, HIGH)
8. Set due date
9. Fetch task details
10. List tasks by:
   - User
   - Status
   - Priority

---

## 3. Non-Functional Requirements

- Thread-safe
- O(1) task lookup
- Modular & extensible
- Clean separation of concerns
- Easy to scale
- Maintainable code structure

---

## 4. Core Entities

### User
- ID
- Name
- Email

### Task
- ID
- Title
- Description
- Status
- Priority
- Due Date
- Assignee
- Comments
- Created At

### Comment
- ID
- Content
- Created By
- Created At

### TaskManager
- Stores tasks
- Handles business logic
- Provides APIs for operations

---

## 5. System Flow

### Task Creation Flow

1. User provides title, description, priority, due date
2. System generates unique task ID
3. Task status is set to TODO
4. Task stored in memory map
5. Task returned to caller

---

### Task Assignment Flow

1. Task ID is provided
2. User is attached as assignee
3. Task updated

---

### Status Update Flow

1. Task ID is provided
2. Validate task exists
3. Update status
4. Save updated task

---

### Add Comment Flow

1. Task ID provided
2. Validate task exists
3. Create comment object
4. Append comment to task

---

### Fetch Task Flow

1. Task ID provided
2. Lookup in map (O(1))
3. Return task details

---

## 6. Data Structures Used

| Feature        | Data Structure | Reason |
|---------------|---------------|--------|
| Task Storage  | Map[string]*Task | O(1) lookup |
| Concurrency   | RWMutex | Thread safety |
| Comments      | Slice | Ordered storage |

---

## 7. Design Principles Used

### Single Responsibility Principle
- User → user data
- Task → task details
- Comment → comment data
- TaskManager → business logic

### Open/Closed Principle
- Easy to add:
  - Labels
  - Projects
  - Teams
  - Notifications

### Separation of Concerns
- Models separated from Service layer

---

## 8. Time Complexity

| Operation        | Complexity |
|-----------------|-----------|
| Create Task      | O(1) |
| Assign Task      | O(1) |
| Update Status    | O(1) |
| Add Comment      | O(1) |
| Get Task         | O(1) |

---

## 9. Future Enhancements

- Add Projects
- Add Tags/Labels
- Add Notification Service
- Add REST API Layer
- Add Database Persistence (Postgres)
- Add Redis Caching
- Add Role-Based Access Control
- Add Search & Filtering
- Add Pagination
- Add Event-driven Architecture

---

## 10. Scalability Plan

To scale this system:

- Replace in-memory storage with Database
- Add Redis for caching
- Add indexing for search
- Use microservices architecture
- Add message queue for async updates
- Horizontal scaling with load balancer

---

## 11. Interview Explanation Strategy

When explaining in interview:

1. Clarify requirements
2. Define entities
3. Explain relationships
4. Discuss data structures
5. Walk through flows
6. Mention scalability & improvements
7. Mention thread safety
8. Discuss time complexity

---

## 12. Architecture Overview

```
Client
↓
TaskManager (Service Layer)
↓
In-Memory Storage (Map)
↓
Models (User, Task, Comment)
``` 


---

## 13. Why This Design is Production-Ready

- Thread-safe (Mutex)
- Clean separation
- Easy to extend
- O(1) operations
- Modular structure
- Supports scaling

---
