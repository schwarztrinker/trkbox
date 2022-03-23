# Database

Running a database for the test environment

```
docker run --detach --name some-mariadb --env MARIADB_USER=gorm --env MARIADB_PASSWORD=gorm --env MARIADB_ROOT_PASSWORD=gorm  -e MARIADB_DATABASE=gorm -p 3306:3306 mariadb:latest
```