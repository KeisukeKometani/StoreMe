# Online Shop

```go
// gorm.Modelの定義
type Model struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}
```

## Build a Development Env

### Requirements

- Docker
- docker-compose

### Alternative requirement

- make
- dip (Ruby gem)
    - ruby

## Startup of local environment

### Using make

```sh
make provision
make up
```

Or, if you want it to run in the background, use the command below.

```sh
make provision
make upd
```

### Using dip (*not working at moment*)

```sh
dip provision
dip up
```

## Migration DB

Migration files are located at db/migrate directory.

### Execute Migrates

#### Migrate up (CREATE TABLE etc..)

```sh
make migrate_up
```

#### Migrate down (DROP TABLE etc...)

```sh
make migrate_down
```

#### Migrate force

example:
```sh
make migrate_force VER=20220312013850
```

### Create new Migrate file

example:
```sh
make migrate_g NAME=create_products_table
```

## Access to the application

Go to the following URL with your browser.

http://localhost:8080/

## Shutdown or cleanup of local environment

### down

```sh
make down
```

### Cleanup all docker images, volumes and orphans

```sh
make downall
```
