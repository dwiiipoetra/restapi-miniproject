# REST API Mini Project Golang at CAP Celerates 2022

Golang Rest API with Fiber, PostgreSQL, Gorm & JWT

## Install Fiber Framework:
```
go get github.com/gofiber/fiber/v2
```

## Install GORM with PostgreSQL:
```
go get -u gorm.io/gorm
```
```
go get -u gorm.io/driver/postgres
```

### Using bcrypt Package for Password Hashing
```
go get golang.org/x/crypto/bcrypt
```

## Run the Project
```
go run main.go
```

## Domain Used

- Motorcycle (CRUD)
- User (Register, Login)

## Use Migrate to Make Table in PostgreSQL 

- Uncomment function MigrateMotorcycles & MigrateUsers on models/motorcycles.go & models/users.go
- And also uncomment file which calls the migrate function on the main.go

## There are 7 Endpoints that can be Accessed

### 1. Endpoint Post Motorcycle
```
/api/create_motorcycle
```
### Body
```
{
    "model_name":"Supra Fit",
    "machine_type": "2 silinder",
    "year":2012,
    "color":"matte blue"
}
```
### Result
```
{
    "data": {
        "model_name": "Supra Fit",
        "machine_type": "2 silinder",
        "year": 2012,
        "color": "matte blue"
    },
    "message": "motorcycle succesfully created"
}
```

### 2. Endpoint Get All Motorcycles
```
/api/motorcycles
```
### Result
```
{
    "data": [
        {
            "id": 3,
            "model_name": "CBR 150R Streetfire",
            "machine_type": "4 Langkah, DOHC 4 Katup",
            "year": 2002,
            "color": "red"
        },
        {
            "id": 4,
            "model_name": "CBR 250RR",
            "machine_type": "4 Stroke, 8-Valve, Parallel Twin Cylinder",
            "year": 2016,
            "color": "red"
        },
        {
            "id": 5,
            "model_name": "Genio",
            "machine_type": "4 Langkah, SOHC, eSP",
            "year": 2019,
            "color": "matte blue"
    ],
    "message": "successfully get motorcycles"
}
```

### 3. Endpoint Get Single Motorcycle
```
/api/motorcycle/1
```
### Result
```
{
    "data": {
        "id": 1,
        "model_name": "CBR 150R",
        "machine_type": "4 Langkah, DOHC",
        "year": 2002,
        "color": "blue"
    },
    "message": "successfully get motorcycle"
}
```

### 4. Endpoint Put Motorcycle
```
/api/update_motorcycle/15
```
### Body
```
{
    "model_name":"Supra Fit 125",
    "machine_type": "2 silinder",
    "year":2012,
    "color":"blue"
}
```
### Result
```
{
    "message": "motorcycle successfully updated"
}
```

### 5. Endpoint Delete Motorcycle
```
/api/delete_motorcycle/15
```
### Result
```
{
    "message": "motorcycle successfully deleted"
}
```

### 6. Endpoint Register User
```
/api/register
```
### Body
```
{
    "name":"Moh. Ari",
    "email":"ari@gmail.com",
    "password":"ari"
}
```
### Result
```
{
    "data": {
        "id": 8,
        "name": "Moh. Ari",
        "email": "ari@gmail.com"
    },
    "message": "user succesfully created"
}
```

### 7. Endpoint Login User
```
/api/login
```