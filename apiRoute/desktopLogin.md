`POST` /auth/desktopLogin

Route pour le login de l'application desktop.

**Body :** 

- id - required

```json
{
    "id" : "000355"
}
```

**Response Code :** 

200

400 :

```json
{
	"error": "Invalid request"
}
```

```json
{
	"error": "Invalid id"
}
```

```json
{
	"error": "The id is expired"
}
```

401:

```json
{
	"error": "Invalid token"
}
```

500:

```json
{
	"error": "Internal server error",
}
```

</aside>