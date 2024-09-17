
`Websocket` /auth/ws

Route pour se connecter au websocket pour le login sur l'application desktop.

 **Message Command :**

- /generateNewId
    
    <aside>
    ❓ This commands allows you to regenerate a new 6 digits id.
    
    </aside>
    
    <aside>
    💬 Response : id: 1
    
    </aside>
    

**Received Message :** 

<aside>
💬 The id to login to the app.
id: 123456

</aside>

<aside>
💬 When the token expires.
id expired

</aside>

<aside>
💬 When there is an error
error

</aside>

**Response Code :** 

500:

```json
{
	"error": "Internal server error",
}
```

</aside>