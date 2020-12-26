# Server Arch

* Database: **SQLite**
* Language: **Go**
* ORM: **GORM**

### Data Models

* Projects
* Entities

**Project is the top level model, Entity depends on Projects**

==================================

```
Projects {
    Name        String
    Description String
    ID          Int
}
```

```
Entities {
    Key       String
    Value     String
    Note      String
    ID        Int
    ProjectID Int
}
```