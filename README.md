# Task Tracker CLI

A simple task tracker to manage your tasks from the command line. Allows you to add, update, delete and list tasks, as well as mark their status.

## Requirements

- [Go](https://golang.org/doc/install) (1.18 o superior)

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/tu-usuario/task-tracker.git
   cd task-tracker
   ```

2. **Compile proyect:**

```bash
   make build
```

# Usage

The CLI supports the following commands:

a. **Add a Task**

```bash
./task-cli add "Task description"
./task-cli add "Buy groceries"
```

b. **Update a Task**

```bash
./task-cli update <ID> "New task description"
./task-cli update 1 "Buy groceries and cook dinner"
```

c. **Delete a Task**

```bash
./task-cli delete <ID>
./task-cli delete 1
```

d. **Mark a Task as In Progress**

```bash
./task-cli mark-in-progress <ID>
./task-cli mark-in-progress 1
```

e. **Mark a Task as Done**

```bash
./task-cli mark-done <ID>
./task-cli mark-done 1
```

f. **List All Tasks**

```bash
./task-cli list
```

g. **List Tasks by Status**

```bash
./task-cli list <status>
./task-cli list todo
./task-cli list done
./task-cli list in-progress
```

# Project Structure

- `main.go`: The main CLI code handling command logic.
- `/task`: Package containing the logic for managing tasks (add, update, delete, list).
- `tasks.json`: JSON file where tasks are stored. It will be automatically created if it does not exist.

# Challange roadmap

- https://roadmap.sh/projects/task-tracker