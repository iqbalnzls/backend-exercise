# backend-exercise<br>


## 🚀 Getting Started

1. Run docker in local computer, then type the following command in the terminal:
    ```sh
    make run
    ```

2. Run Migrations

   Once the containers are up, run the following command to connect to the database:
   ```sh
   make migrate
   ```
   Then type the SQL command in migrations folder (init_schema.sql) to set up the database.


3. At this point, you can access the application at `http://localhost:6000`.


4. Stop the containers

   To stop the containers by running:

    ```sh
    make stop
    ```
   

## 📝 API Documentation

For listings and users apis, there is a security layer to ensure that only authenticated users can access the endpoints.
1. Using the allowlist IP address with `127.192.1.1` as the only allowed IP address.
2. Using basic authentication with username `admin` and password `password123`.


### ✅ Listings

- **Create**

    ```sh
    curl --location --request POST 'localhost:6000/listings' \
    --header 'Ip-Address: 127.192.1.1' \
    --header 'Content-Type: application/x-www-form-urlencoded' \
    --header 'Authorization: Basic YWRtaW46cGFzc3dvcmQxMjM=' \
    --data-urlencode 'user_id=1' \
    --data-urlencode 'listing_type=rent' \
    --data-urlencode 'price=190000'
    ```

- **Get All**

    ```sh
    curl --location --request GET 'localhost:6000/listings?page_num=1&page_size=3&user_id=1' \
    --header 'Ip-Address: 127.192.1.1' \
    --header 'Authorization: Basic YWRtaW46cGFzc3dvcmQxMjM=' \
    --data ''
    ```  

### ✅ Users

- **Create**

    ```sh
    curl --location --request POST 'localhost:6000/users' \
    --header 'Ip-Address: 127.192.1.1' \
    --header 'Content-Type: application/x-www-form-urlencoded' \
    --header 'Authorization: Basic YWRtaW46cGFzc3dvcmQxMjM=' \
    --data-urlencode 'name=user3'
    ```

- **Get By ID**

    ```sh
    curl --location --request GET 'localhost:6000/users/90' \
    --header 'Ip-Address: 127.192.1.1' \
    --header 'Authorization: Basic YWRtaW46cGFzc3dvcmQxMjM=' \
    --data ''
    ```

- **Get All**

    ```sh
    curl --location --request GET 'localhost:6000/users?page_num=1&page_size=1' \
    --header 'Ip-Address: 127.192.1.1' \
    --header 'Authorization: Basic YWRtaW46cGFzc3dvcmQxMjM=' \
    --data ''
    ```


### ✅ Public API

- **Get All Listings**

    ```sh
    curl --location --request GET 'localhost:6000/public-api/listings?page_num=1&page_size=10&user_id=3'
    ```

- **Create Listings**

    ```sh
    curl --location --request POST 'localhost:6000/public-api/listings' \
    --header 'Content-Type: application/json' \
    --data '{
    "user_id": 3,
    "listing_type": "sale",
    "price": 8900077
    }'
    ```

- **Create Users**

    ```sh
    curl --location --request POST 'localhost:6000/public-api/users' \
    --header 'Content-Type: application/json' \
    --data '{
    "name": "Lorel Ipsum"
    }'
    ```