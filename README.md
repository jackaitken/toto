# Toto 🚽

Toto is a lightweight command-line application that allows you to create and manage todo lists right in your terminal. It keeps your todo lists in a local JSON file for speed and simplicity. 

```
➜ toto l
Todo list for today
1: welcome
2: to
3: toto!
```

### Subcommands:

```
list | l: list today's todo list
add | a <new todo text>: add a new todo item for today. If there is no list for today, one will be created first
edit | e <todo_id> <new todo text>: edit a todo item
complete | c <todo_id>: mark a todo item as complete
incomplete | i <todo_id>: mark a todo as incomplete
delete | d <todo_id>: delete a todo item
```

- `list | l`: list today's todo list
- `add | a <new todo text>`: add a new todo item for today. If there is no list for today, one will be created first
- `edit | e <todo_id> <new todo text>`: edit a todo item
- `complete | c <todo_id>`: mark a todo item as complete
- `incomplete | i <todo_id>`: mark a todo as incomplete
- `delete | d <todo_id>`: delete a todo item

This was written with [Cobra](https://github.com/spf13/cobra).

### Installation:

```
go install github.com/jackaitken/toto@latest
```