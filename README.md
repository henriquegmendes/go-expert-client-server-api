# GO Expert Client Server Challenge

This challenge proposed by GO Expert course

## Install Dependencies

    make install

## Run the server with sqlite3 db (needs docker)

    make run-server-with-sqlite

## Run the client

    make run-client

# Testing the whole application locally

- Clone this project
- Create a .env by copy/pasting .env.example file
- Open two terminals.
- Make sure both terminals is located inside this project's folder
- In the first terminal, run `make run-server-with-sqlite` command and wait until server is UP (listening to port 8000)
- In the second terminal, run `make run-client` command, and it should create/update the `cotacao.txt` file located inside `result` folder
- If you want to check the exchanges saved in database, you make the following request:
`
  curl --request GET \
  --url http://localhost:8000/cotacao
`

# Server Available endpoints

## Create a new Exchange and Return USD-BRL Quote Bid Value

### Request

`POST /cotacao`

    curl --request POST \ --url http://localhost:8080/cotacao

### Response

    {
        "bid": 5.2074
    }

## Get All Saved Exchanges

### Request

`GET /cotacao`

    curl --request GET \ --url http://localhost:8080/cotacao

### Response

    [
	    {
		    "ID": 1,
		    "Type": "Dólar Americano/Real Brasileiro",
		    "Value": 5.1986,
		    "CreatedAt": "2022-12-21T19:48:41.28342572Z",
		    "UpdatedAt": "2022-12-21T19:48:41.28342572Z"
	    },
        {
            "ID": 2,
            "Type": "Dólar Americano/Real Brasileiro",
            "Value": 5.1885,
            "CreatedAt": "2022-12-21T19:49:41.28342522Z",
            "UpdatedAt": "2022-12-21T19:49:41.28342522Z"
        }
    ]
