# ATTENDANCE API
This is contract and api use for employee to check-in using hardhat, golang, fiber, postgresql, sqlx, abigen
### INSTALL
``` sh
cd contract && npm install
cd server && go mod tidy
```
### RUN APP
```sh
cd contract && npm start
cd server && go run main.go
```
### RUN TEST
```sh
cd contract && npm run test
```
### ENV FILE FOR SERVER
create .env in server folder
```
DB_HOST=
DB_PORT=5432
DB_NAME=
DB_USER=
DB_PASSWORD=
```
### DB
<img width="614" alt="Screenshot 2024-06-15 at 14 53 38" src="https://github.com/uncleTen276/greendeco/assets/72869134/fe76b2fb-5abd-498b-8d5c-0c262d301f53">

- <b>users</b> table use to store user's information

- <b>events</b> table use to record events which server interact with contract

### SERVER
##### APIS
Swagger doc
```url
http://localhost:8080/api/v1
```

<img width="1405" alt="Screenshot 2024-06-15 at 15 12 02" src="https://github.com/uncleTen276/greendeco/assets/72869134/aaf79330-b591-4c6a-b9ed-a0e4d8e66c9c">


### CMD
- Compile file .sol
```sh
npx hardhat compile
```
- Run contract server
```sh
npx hardhat run scripts/deploy.js  --network <networkName>
```
- Gen new abi file
```sh
solc --abi <contract>.sol -o build
```
- Get hardhat node list
```sh
npx hardhat node
```

### DOC
- [hardhat](https://hardhat.org)

- [fiber](https://gofiber.io)

- [abigen](https://geth.ethereum.org/docs/tools/abigen)






