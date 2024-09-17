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

## Configuration

1. Renommer le fichier ".env.template" en ".env"
2. Modifier le fichier .env pour y mettre la string de connexion à la base de donnée MySql ainsi que la string secrète pour signer les tokens. 

Exemple :

    DB="root:root@tcp(authmysql)/authDb?charset=utf8mb4&parseTime=True&loc=Local"
    Secret="kljfsdlkjKJFHSDFHkndsjkfnJFJASNnjNDuhdJKDN"

## Lancer le projet

  Pour lancer le projet, il suffit de démarrer le docker compose à l'aide de la commande :

    docker compose up
    
## Documentation API

[/auth/Register](./apiRoute/register.md) \
[/auth/Login](./apiRoute/login.md) \
[/auth/validate](./apiRoute/validate.md) \
[/auth/ws](./apiRoute/ws.md) \
[/auth/desktopLogin](./apiRoute/desktopLogin.md)


