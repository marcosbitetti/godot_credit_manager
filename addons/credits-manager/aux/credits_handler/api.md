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
./credits-manager add {"name":"Work", "filename":"res://file", "author":"Joe", "link":"http://...", "type_id":1, "licence_id":1}
```

* update credit
``` bash
./credits-manager update {_id:1, "name":"Work", "filename":"res://file", "author":"Joe", "link":"http://...", "type_id":1, "licence_id":1}
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



