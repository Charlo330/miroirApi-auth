`GET` /auth/validate

Route pour valider le Bearer token.

**Header :** 

- Authorization - required

```
Authorization : Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTk3MzU3MDAsInN1YnMiOjJ9.i7ft_BCzVX0c_PeizoPxMr3fT_egkcDwSekjjqHfcN
```

**Response Code :** 

200 :

returns the user id.

```json
{
	"id": 1
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