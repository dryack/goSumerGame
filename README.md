# goSumerGame
A game based on ideas from [The Sumerian Game](https://en.wikipedia.org/wiki/The_Sumerian_Game), [Hamurabi](https://en.wikipedia.org/wiki/Hamurabi_(video_game)), [Dukedom](https://en.wikipedia.org/wiki/Dukedom_(video_game)), and [Manor](https://web.archive.org/web/20141204104645/http://mmreference.com/product/manor/)

## Components
### Server
A web API that permits user registration and login; game session (savegame)  management, and gameplay via JSON requests.

#### Setup
1. Install Postgres
2. Create `.env.local` file under the program's home directory, modeled after `server.env`
3. Create `DB_USER` with password `DB_PASSWORD`  in Postgres
4. Create `DB_NAME` owned by `DB_USER` in Postgres
5. Define `TOKEN_TTL` (I suggest a large number for development or localhost-only servers)
6. Define `JWT_PRIVATE_KEY` (Making this a secure key won't matter much for development or localhost-only servers)
7. `$ go run server/server.go`
8. Register a user, for example: `$ curl -i -H "Content-Type: application/json" -X POST -d '{"username":"someuser","password":"foobar"}' http://localhost:80/auth/register`. You will receive an auth token from the server following registration, this will be required to perform actions as your user on the server
9. Create a game, for example: `$ curl -d '{"debug":0}' -H "Content-Type: application/json" -H "Authorization: Bearer <token received during registration or login>" -X POST http://localhost:80/api/game` (the debug parameter will ultimately decide how much or how little randomization will be involved in a given savegame)

#### Login
After `TOKEN_TTL` hours has passed, your auth token will have expired, and you'll need to login again.

`$ curl -i -H "Content-Type: application/json" -X POST -d '{"username":"someuser","password":"foobar"}' http://localhost:80/auth/login`

#### Taking a Turn
Currently there's not much going on here, but:
`$  curl -d '{"gameid":2,"purchaseacres":100}' -H "Content-Type: application/json" -H "Authorization: Bearer <token received during registration or login>" -X POST http://localhost:80/api/game/play`

### Client
(Currently unimplemented)

Although the game can be played directly using cURL for example, it's ideally meant to work with a client. The intent is to allow players to implement their own clients as they desire, although I will supply a default client.

