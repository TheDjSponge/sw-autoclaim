<!-- omit in toc -->
# Redemption service


The redemption service is the backend service responsible for user management, coupon storage and redemptions of coupons. It interacts with a 
PostGres database and is aimed to work along with a discord bot as frontend (even thought the project could also serve a frontend page in the future)

The only authorization constraints are currently defined by the 
discord id of the user making a request through the frontend (discord bot). I make the current assumption that the service is only accessible and that a user is only able to modify or get information
about an account that was registered by himself through the discord bot. Currently I assume that no sensitive information is held
in the database (we could argue for hive_id though). 

If the threat model were to be modified and include calls to the 
api from something else than the discord bot (the api being exposed
to the internet for some reason), the implementation of authorization
through a designated token would be necessary. 

- [Backend api documentation](#backend-api-documentation)
  - [coupon management](#coupon-management)
    - [/v1/coupons](#v1coupons)
      - [POST](#post)
      - [GET](#get)
  - [user management](#user-management)
    - [/v1/users](#v1users)
      - [POST](#post-1)
      - [GET](#get-1)
      - [DELETE](#delete)

# Backend api documentation


## coupon management

### /v1/coupons

#### POST

Allows to add a new coupon

Parameters: -
request body: application/json
```json
// example body
{"coupon_codes" : ["code1","code2",]}
```

response: application/json
```json
// example response
{
    "status": 201,
    "message": "Inserted successfully"
}

```
codes:
- 201: successfully deleted user
- 400: invalid inputs, json could not be decoded
- 500: error when trying to add coupons to database

#### GET

allows to display all coupons

Parameters: 
- code: returns information about specific coupon

request body: application/json
```json
// example body
null
```

response: application/json
```json
// example response
[
{
    "Code": "code1",
    "Status": "active",
    "FirstSeenAt": "06-01-2026",
    "UpdatedAt": "06-01-2026"
},
// {
//     ...
// }
]

```
codes:
- 200: successful request
- 500: error when trying to fetch coupons in database


## user management

### /v1/users
#### POST

Allows to register a new user. Before registration, the hive_id is
checked to ensure the user currently exists.

Parameters: -

request body: application/json
```json
// example body
{
    "hive_id": "user_012vc32",
    "server": "europe", // one of europe, global, asia, japan
    "discord_id": "01231231232123",
    "discord_username": "TheRealBob25"
}
```

response: application/json
```json
// If successful
null
// If an error is thrown
{
    "status": 400,
    "message": "Invalid JSON"
}

```
codes:
- 201: user creation succeeded (user already exists / was created)
- 400: invalid inputs, json could not be decoded
- 500: error during user validation/registration

#### GET

Allows to fetch the list of users currently registered.

Parameters: -
request body: application/json
```json
// example body
null
```

response: application/json
```json
// If successful
[
{
    "hive_id": "user_012vc32",
    "server": "europe", // one of europe, global, asia, japan
    "discord_id": "01231231232123",
    "discord_username": "TheRealBob25"
},
// {
//     ...
// }
]

// If an error is thrown
{
    "status": 400,
    "message": "Invalid JSON"
}
```
codes:
- 200: request successful
- 500: error when fetching 

#### DELETE

Allows to remove a registed user

Parameters: -
request body: application/json
```json
// example body
{
    "discord_id":"01231231232123",
    "hive_id":"user_012vc32",
    "server":"europe",
}
```

response: application/json
```json
// example response
{
    "status": 200,
    "message": "User deleted successfully"
}

```
codes:
- 200: successfully deleted user
- 400: invalid inputs, json could not be decoded
- 500: error when trying to delete user

