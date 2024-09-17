# Projet miroir Api Authentification

Cette Api à été conçue pour un projet scolaire de miroir intelligent qui avait
besoin d'une API d'authentification. L'API utilise les JWT.

L'API comporte un controlleur (usersController) permettant d'effectuer :

- création de compte
- Connexion au compte
- Validation de token JWT

Elle comporte aussi controlleur (desktopAppController) permettant de se login à l'aide d'un code qui est généré par l'application. Une connexion websocket est nécessaire.


## Prérequis

Vous devez avoir docker d'installer sur votre machine.

https://docs.docker.com/engine/install/

## Lancer le projet

  Pour lancer le projet, il suffit de démarrer le docker compose à l'aide de la commande :

    docker compose up
    
## Documentation API

[/auth/Register](./apiRoute/register.md) \
[/auth/Login](./apiRoute/login.md) \
[/auth/validate](./apiRoute/validate.md) \
[/auth/ws](./apiRoute/ws.md) \
[/auth/desktopLogin](./apiRoute/desktopLogin.md)


