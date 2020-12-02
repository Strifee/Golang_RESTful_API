# GOLang RESTful API

> RESTful API to create, read, delete and update Users. (database implementation will be added later)

## Bash Start


``` bash
# Installing mux router
go get -u github.com/gorilla/mux
# generate and run .exe
go build
./REST_API
```
## Testing the API 

### Get All Users
``` bash
GET api/user
```
### Get a user
``` bash
GET api/user/{Username}
```

### Create a User
``` bash
POST api/user

# Request 
# {
#   "Username":"meee",
#   "Firstname":"EL Mehdi",
#   "lastname":"SOUMMER",
#   "Email":"test@gmail.com",
# }
```

### Delete a User
``` bash
DELETE api/user/{Username}
```

### Update a User
``` bash
PUT api/user/{Username}

# Request 
# {
#   "Username":"hiiii",
#   "Firstname":"Mehdi",
#   "lastname":"SOUMMMMER",
#   "Email":"test1@gmail.com",
# }
```
