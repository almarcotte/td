# td - cli todo

`td` is a command-line todo / task management app that allows adding tasks on a per-project basis or globally. Project
task lists are stored in a specific folder on disk whereas the global list is typically stored in the user's home
directory.

## Usage

Assuming `td` is in your `$PATH`:

```text
td <command> [params]
    |
    a       Add a new item
    c       Mark an item as completed
    p       Mark an item as in progress
    r       Revert item status to incomplete
    t       Add tags to an item
    f       Find an item
    ls      List all items for the current project
```

To manage the global list, most of the commands above are available, simply prefix them with `g` (i.e. `ga` adds an item
to the global list, `gf` searches the global list, and so on)