# User-Service
This service handles all routes which require access to the users account.

## Setup 
To run this service independently, you have to install all necessary dependencies, like so:
```sh
go get ./...
```
To run the app, run:
```sh
go run main.go
```

## REST API
This service has the following routes:
- */api/users/change-lang* changes the users language in the db to the language specified in the Lang header. If the language wasn't specified in the SUPPORTED_LANGS env variable it returns a status code 400 (Bad Request).
- *api/users/delete-acc* deletes the user account with the specified uid.
- */api/users/sell-stock* sells a specified amount of stocks from the specified symbol if the user has bought them before.
- */api/users/buy-stock* adds stocks to the users portfolio if the user has enough money to buy the specified amount of shares. It also expects the stock price to be set by the user, because the user recieves the newest price in real-time. In the future the route will get the stock price from the *Stock-Service* since the current implementation's kinda dumb.
- */api/users/cash* returns the amount of cash a user has.
- */api/users/owned-stocks* returns all stock information of all stocks the user has bought.

**Important**: All routes require a UID header.

## Contributing 
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
MIT License. Click [here](https://choosealicense.com/licenses/mit/) or see the LICENSE file for details.