# Golang REST API

A small rest api in golang

## API Endpoints

The sample app is hosted in heroku https://go-chef-rest.herokuapp.com

* POST /user/signup

```
curl -X POST \
  https://go-chef-rest.herokuapp.com/user/signup \
  -H 'Content-Type: application/json' \
  -d '{
	"email":ENTER EMAIL,
    "password":ENTER PASSWORD(more than 8 char)
}'

Returns
   {
       authToken: <AUTH_TOKEN>,
       success: true
   }
```

* POST /user/signin

```
curl -X POST \
  https://go-chef-rest.herokuapp.com/user/signin \
  -H 'Content-Type: application/json' \
  -d '{
	"email":ENTER EMAIL,
    "password":ENTER PASSWORD(more than 8 char)
}'

Returns
   {
       authToken: <AUTH_TOKEN>,
       success: true
   }
```

* GET /user/profile

```
curl -X GET \
  https://go-chef-rest.herokuapp.com/user/profile \
  -H 'Authorization: Bearer <AUTH_TOKEN>'

Returns:
    {
        success: true,
        user: {
            // USER DATA
        }
    }
```

* PATCH /user/profile/update

```
curl -X PATCH \
  https://go-chef-rest.herokuapp.com/user/profile/update \
  -H 'Authorization: Bearer <AUTH_TOKEN>' \
  -d '{
	"name" : <UPDATE NAME>,
	"password": <UPDATE PASSWORD>
}'

Returns:

    {
        "success": true,
        "user": "Succesfully updated"
    }
```

## Install locally

* Clone the repo on go workspace `git clone https://github.com/droidnoob/golang-rest.git`
* Go ot the directory 
* Create `.env` file and fill in using the `env_example` file
