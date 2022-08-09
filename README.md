# Golang Fiber JWT Authentication

### We use golang jwt authentication for backend
### Check & setup the vuejs jwt authentication project in the following link.
[Vue.js JWT Authentication](https://github.com/RakibSiddiquee/vuejs-jwt-auth) 

## Install Fiber Framework:
```
go get github.com/gofiber/fiber/v2
```

## Install GORM with MySQL:
```
go get -u gorm.io/gorm
```
```
go get -u gorm.io/driver/mysql
```

### We used bcrypt for password hashing

## Run the project
```
go run main.go
```

### Registration URL:
```
http://localhost:8000/api/register
```

### Login URL:
```
http://localhost:8000/api/login
```

### Get logged in user after login:
```
http://localhost:8000/api/user
```

### Logout URL:
```
http://localhost:8000/api/logout
```



# REST API Mini Project Golang at CAP Celerates 2022

Golang Rest API with Fiber, PostgreSQL, Gorm & JWT

## Domain Used

- Motorcycle (CRUD)
- User (Register, Login)

// func MigrateUsers(db *gorm.DB) error {

## Use Migrate to Make Table in PostgreSQL 

- Uncomment function MigrateMotorcycles & MigrateUsers on models/motorcycles.go & models/users.go
- And also uncomment file which calls the migrate function on the main.go

## There are 7 endpoints that can be accessed

### Endpoint Post Motorcycle
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

### Endpoint Get All Motorcycles
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

### Endpoint Get Single Motorcycle
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

### Endpoint Put Motorcycle
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

### Endpoint Delete Motorcycle
```
/api/delete_motorcycle/15
```

### Result
```
{
    "message": "motorcycle successfully deleted"
}

```

### Endpoint Register User
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