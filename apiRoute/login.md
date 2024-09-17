
`POST` **auth/login**

Route pour se login à l'application.

**Body :** 

- email - required
- password - required, min 6, max

```json
{
    "email": "test124@hotmail.com",
    "password":"ALLO"
}
```

**Response Code :** 

200 :

```json
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTk3MzE0MTksInN1YnMiOjF9.pgeiWXURK4K4Nbvbh1RR45oOB6kSASeyIkE9S03apI0"
}
```

400 :

```json
{
	"error": "Invalid request"
}
```

401: 

```json
{
	"error": "Invalid email or password"
}
```

500:

```json
{
	"error": "Internal server error",
}
```