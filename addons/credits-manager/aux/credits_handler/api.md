# List credits

executable *list* *asc*/*desc* [*search string*]

* full list 
``` bash
./credits-manager list asc
```

* list filter/search by Name or Author
``` bash
./credits-manager list asc foo
```

# edit credit
* add credit
``` bash
./credits-manager add {"name":"Work", "filename":"res://file", "author":"Joe", "link":"http://...", "type":"Music", "licence":"MIT"}
```

* update credit
``` bash
./credits-manager update {_id:1, "name":"Work", "filename":"res://file", "author":"Joe", "link":"http://...", "type":"Music", "licence":"MIT"}
```

* check credit is already designet to a file
``` bash
./credits-manager file-exists res://file
```

# delete credit
``` bash
./credits-manager delete  1
```

# list licences
``` bash
./credits-manager licences asc/desc
```

# list types
``` bash
./credits-manager types asc/desc
```

# auto-complete for author list
``` bash
./credits-manager auto-complete-author [characters here]
```



