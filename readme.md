# go-graphql-demo

- Simple User management service created by using [go-graphql]() and [gorm]().

## Installation

**Clonning the repo**

```bash
git clone https://github.com/Besufikad17/go-graphql-demo.git 
```

**Installing libraries**

```bash
go get
```

**Setting up configurations**

```
// storing configs in .env file
DB_CONNECTION="postgres://username:password@localhost:5432/DB_NAME"
JWT_SECRET=""
``` 

**Running**

```bash
go run main.go
```

## Usage

- `baseUrl` = http://localhost:8000/graphql

### Authentication

**SignUp**

```bash
curl {{baseUrl}}?query=mutation{signup(firstName: "John", lastName: "Doe", email: "johndoe@gmail.com", phoneNumber: "+2519123456", password: "123456"){message, token}}
```

**Login**

```bash
curl {{baseUrl}}?query=mutation{login(loginText: "+2519123456", password: "123456"){message, token}}
```

### CRUD

**AddUser**

```bash
curl {{baseUrl}}?query=mutation{add(firstName: "Jane", lastName: "Doe", email: "janedoe@gmail.com", phoneNumber: "+2519876543"){firstName, lastName, email, phoneNumber}}
```

**GetUsers**

- Getting all users from the DB.

```bash
curl {{baseUrl}}?query={users{id,%20firstName,%20lastName,%20email,%20phoneNumber}}
```

- Getting the first 10 users from the DB using `skip` and `take`.

```bash
curl {{baseUrl}}?query={users(skip: 0, take: 10){id,%20firstName,%20lastName,%20email,%20phoneNumber}} 
```

- Searching users from the DB using specific `text`.

```bash
curl {{baseUrl}}?query={users(text: "Jane"){id,%20firstName,%20lastName,%20email,%20phoneNumber}}
```

**GetUserById**

```bash
curl {{baseUrl}}?query={user(id:1){id,%20firstName,%20lastName,%20email,%20phoneNumber}}
```

**UpdateUser**

```bash
curl {{baseUrl}}?query=mutation{update(id:1, firstName: "Abebe", lastName: "Kebede", email: "abekebe@gmail.com", phoneNumber: "0987654321"){id, firstName, lastName, email, phoneNumber}}
```

**DeleteUser**

```bash
curl {{baseUrl}}?query=mutation{delete(id:2){text}}
```
