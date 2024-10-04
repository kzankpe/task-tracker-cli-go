# Task Tracker Console Application
Project Task URL : https://roadmap.sh/projects/task-tracker

Golang CLI solution for the task-tracker [challenge](https://roadmap.sh/projects/task-tracker) from [roadmap.sh](https://roadmap.sh/).

This CLI allows you to add, update, delete, and list tasks with various statuses such as "To-Do," "In-Progress," and "Done."

## Features

- **Add a New Task**: Create new tasks with a simple command.
- **Update Task**: Modify the description of an existing task.
- **Delete Task**: Remove tasks by their ID.
- **List Tasks**: Display all tasks or filter tasks by status.
- **Change Task Status**: Mark tasks as "To-Do," "In-Progress," or "Done."
- **Help Command**: Display a list of available commands.
- **Clear Console**: Clear the console and display the welcome message again.

## Installation

To run this application, follow these steps:

```bash
git clone https://github.com/kzankpe/task-tracker-cli-go.git
```
Run the following command to build and run the project:

```bash
go build -o task-cli
./task-cli --help # To see the list of available commands
```

## Usage

```bash
task-cli [command]
```

###  Available Commands

- **add [description]**: Adds a new task with the provided description.
- **update [id] [new description]**: Updates the task with the given ID.
- **delete [id]**: Deletes the task with the given ID.
- **list**: Lists all tasks.
- **list [status]**: Lists tasks filtered by status ("todo", "in-progress", "done").
- **markInProgress [id]**: Marks the task with the given ID as "In-Progress".
- **markDone [id]**: Marks the task with the given ID as "Done".
- **help**: Display help about the cli.

### Example

```bash
# Adding a new task
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
task-cli update 1 "Buy groceries and cook dinner"
task-cli delete 1

# Marking a task as in progress or done
task-cli mark-in-progress 1
task-cli mark-done 1

# Listing all tasks
task-cli list

# Listing tasks by status
task-cli list done
task-cli list todo
task-cli list in-progress
```