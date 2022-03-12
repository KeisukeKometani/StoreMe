# Online Shop

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

## Model rule

Always include gorm.Model
```go
type Product struct {
	gorm.Model
	Name	  string  `json:"name" `
	Price	  int  `json:"price" `
	Description string  `json:"description" `
	Image	  string  `json:"image" `
}

// gorm.Modelの定義
type Model struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}
```

