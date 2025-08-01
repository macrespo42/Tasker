# [Tasker CLI](https://roadmap.sh/projects/task-tracker) ✅

**A simple command-line tool to manage your tasks efficiently.**  
Track what you need to do, what you're working on, and what you've completed — all from your terminal.

---

## Requirements

Before using **Task Tracker CLI**, make sure your environment meets the following requirements:

- A Unix-like system (Linux, macOS)
- A programming language runtime (e.g., Python, Node.js, etc.)
- A terminal with support for command-line tools
- No third-party libraries required (uses only native modules)

---

## Installation

To install or run locally, clone the repo and make the script executable:

```bash
git clone https://github.com/macrespo42/Tasker
cd Tasker
go run . <command> <arguments>

```

or to installed it globally

```bash
go install github.com/macrespo42/Tasker@latest
# if $GOPATH isn't in your .zshrc/.bashrc
export PATH=$PATH:$(go env GOPATH)/bin > ~/.zshrc
source ~/.zshrc
# Then run the program from anywhere
Tasker <command> <arguments>
```

---

## Commands

Run any command using the format:

```bash
Tasker <command> <arguments>
```

### Available Commands:

add "Task description"
➜ Add a new task
Example:

```bash
Tasker add "Buy groceries"
```

update <task_id> "Updated description"
➜ Update task description
Example:

```bash
Tasker update 1 "Buy groceries and cook dinner"
```

delete <task_id>
➜ Delete a task
Example:

```bash
Tasker delete 1
```

mark-in-progress <task_id>
➜ Mark a task as in progress
Example:

```bash
Tasker mark-in-progress 2
```

mark-done <task_id>
➜ Mark a task as done
Example:

```bash
Tasker mark-done 2
```

list
➜ List all tasks
Example:

```bash
Tasker list
```

- Tasker list done
  ➜ List only completed tasks

- Tasker list todo
  ➜ List tasks not started

- Tasker list in-progress
  ➜ List tasks in progress

Task Structure

Each task is stored in the tasks.json file with the following properties:

```json
{
  "id": 1,
  "description": "Buy groceries",
  "status": "todo", // or "in-progress", "done"
  "createdAt": "2025-08-01T12:00:00Z",
  "updatedAt": "2025-08-01T12:00:00Z"
}
```

## Error Handling

The CLI gracefully handles common errors such as:
Missing or incorrect command arguments
Invalid task IDs
File not found (automatically created)
JSON parsing errors
