# Server Arch

* Database: **SQLite**
* Language: **Go**
* ORM: **GORM**

### Data Models

* Projects
* Entities
* Entitity Values

**Project is the top level model, Entity depends on Projects and Entitity Values depend on Entitity**

==================================

```
Projects {
    Name: String;
    Description: String;
    Id: Int;
}
```

```
Entities {
    Name: String;
    Id: Int;
    ProjectId: Int;
}
```

```
EntityValues {
    Name: String;
    Id: Int;
    EntityId: Int;
}
```