
`POST` **auth/register**

Route pour se créer un compte.

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
	"email": "test124@hotmail.com",
  "password":"ALLO"
}
```

400 :

```json
{
	"error": "Invalid request"
}
```

```json
{
	"error": "The email is already in use"
}
```

500:

```json
{
	"error": "Internal server error",
}
```

</aside>