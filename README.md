# Tasker CLI ✅

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
git clone https://github.com/your-username/task-tracker-cli.git
cd task-tracker-cli
chmod +x task-cli

```

---

## Commands

Run any command using the format:

task-cli <command> <arguments>

Available Commands:

    add "Task description"
    ➜ Add a new task
    Example:

task-cli add "Buy groceries"

update <task_id> "Updated description"
➜ Update task description
Example:

task-cli update 1 "Buy groceries and cook dinner"

delete <task_id>
➜ Delete a task
Example:

task-cli delete 1

mark-in-progress <task_id>
➜ Mark a task as in progress
Example:

task-cli mark-in-progress 2

mark-done <task_id>
➜ Mark a task as done
Example:

task-cli mark-done 2

list
➜ List all tasks
Example:

    task-cli list

    list done
    ➜ List only completed tasks

    list todo
    ➜ List tasks not started

    list in-progress
    ➜ List tasks in progress

Task Structure

Each task is stored in the tasks.json file with the following properties:

{
"id": 1,
"description": "Buy groceries",
"status": "todo", // or "in-progress", "done"
"createdAt": "2025-08-01T12:00:00Z",
"updatedAt": "2025-08-01T12:00:00Z"
}

Error Handling

The CLI gracefully handles common errors such as:

    Missing or incorrect command arguments

    Invalid task IDs

    File not found (automatically created)

    JSON parsing errors
