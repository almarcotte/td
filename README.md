# td - cli todo

[![Build Status](https://travis-ci.org/gnumast/td.svg?branch=master)](https://travis-ci.org/gnumast/td)

`td` is a command-line todo / task management app that allows adding tasks on a per-project basis or globally. Project
task lists are stored in a specific folder on disk whereas the global list is typically stored in the user's home
directory.

## Usage

Assuming `td` is in your `$PATH`:

```text
td <command> [params]
    a <item description> [!tag1 ... !tagn] [@due date]   
            Add a new item with the given tags and due date. Tags should be single words starting with !. Due date is
            can be in common formats or @today, @tomorrow, @nextweek. Returns the item id.
    c <item id>     
            Mark an item as completed.
    p <item id>
            Mark an item as in progress
    r <item id>
            Revert item status to incomplete
    t <item id> [!tag1 ... !tagn]
            Add tags to an item. If no tags are specified, all tags are removed from the item.
    f <needle>
            Find an item by description or tag. To search by tag, needle should start with a !.
    ls      
            List all items for the current project.
```

To manage the global list, most of the commands above are available, simply prefix them with `g` (i.e. `ga` adds an item
to the global list, `gf` searches the global list, and so on)