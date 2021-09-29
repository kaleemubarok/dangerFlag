# dangerFlag

Simple micro service to update dummy data on a json file and serve as an endpoint to cosume the data generated.

## Service Endpoint

    - `GET /data` -> Get latest data

### How to Run

- Go to project folder

  Usually, it would be `cd go/src/github.com/kaleemubarok/dangerFlag`.

- Fill in the environment variables

  Copy the sample env file.
    ```
    cp env.sample .env
    ```
  Then, fill the values according to your setting in `.env` file.

- Run the database

  - Make sure to run MySQL and create a DB.
  - or you can easily run attached docker compose with db initiation
  ```
  docker compose up
  ```

- Download the dependencies

    ```
    go mod download 
    ```

- Run the application

    ```
    go run main.go
  
###Sample Request : Body JSON
- POST
 ```
  {
    "orderAt": "2021-09-17T21:18:57+07:00",
    "customerName": "Adam Benjamin",
    "items": [
      {
        "itemCode": "ADXM1",
        "description": "Tortoise baquet",
        "quantity": 1
      }
    ]
  }
 ```
- PUT
 ```
  {
    "orderId": 1,
    "orderAt": "2021-09-17T21:18:57+07:00",
    "customerName": "Adam Benjamin",
    "items": [
      {
        "itemId": 1,
        "itemCode": "ADXM1",
        "description": "Tortoise baquet",
        "quantity": 1
      }
    ]
  }
 ```
