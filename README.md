# task-tracker
A simple [task tracker](https://roadmap.sh/projects/task-tracker) CLI tool built with Go

## Installation
```bash
go build
```

## Usage
### Add
`task-tracker add "Task description"`
### Update
`task-tracker update id "Task description"`
### Delete
`task-tracker delete id`
### List
- `task-tracker list`
- `task-tracker list todo`
- `task-tracker list in-progress`
- `task-tracker list done`
### Mark as
- `task-tracker mark-in-progress id`
- `task-tracker mark-done id`