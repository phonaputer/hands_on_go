# "Hands-on-Go" User Application

This is the directory in which you will write the code for the User Application you'll be developing in this workshop.

Please create a branch for your code using the following format: `${your_name}/${workshop_start_date_yyyyMM}`.

For example, for an attendee named Jane Doe who attended the September 2022 session, the branch name would be `jane.doe/202209`.

## Dependencies

The User Application uses the same Dockerized MySQL from the pre-assignment.
A Docker Compose file for this can be found in the `deployments/local` subdirectory of this directory.

## How to find the instructor's code

You can see the up-to-date code the instructor is writing in your workshop session by checking out the correct branch for your session.

The branch name will be in the following format: `instructor/${workshop_start_date_yyyyMM}`.

So for a workshop starting in September 2022, the branch would be `instructor/202209`.

Old instructor branches will not be deleted so you can check this anytime.

# Endpoint Specs

The following are the specifications for the HTTP endpoints that will be developed during this course.

## 1. POST /users

Create a new user.

The same email address cannot be registered twice.
An error will be returned if the provided email already exists in the system.

### Request

<table>
<tr><th>Method</th><td>POST</td></tr>
<tr><th>Path</th><td>/users</td></tr>
</table>

#### Body

<table>
<tr><th>Content-Type</th><td>application/json</td></tr>
</table>

| Property          | Nullable | Type   | Description               | Example                  |
| ----------------- | -------- | ------ | ------------------------- | ------------------------ |
| $ (root)          | No       | Object |                           |                          |
| $.**first_name**  | No       | String | The user's given name.    | John                     |
| $.**middle_name** | Yes      | String | The user's middle name.   | Robert                   |
| $.**last_name**   | No       | String | The user's family name.   | Doe                      |
| $.**age**         | No       | Number | User's age in years.      | 30                       |
| $.**email**       | No       | String | The user's email address. | john.r.dough@example.com |

#### Example Request

```
POST /users HTTP/2
Host: 127.0.0.1:8080
Content-Type: application/json
Content-Length: 26

{
    "first_name": "John",
    "middle_name": "Robert",
    "last_name": "Doe",
    "age": 30,
    "email": "john.r.dough@example.com"
}
```

### Response

#### Status

<table>
<tr><th>User was successfully created.</th><td>201</td></tr>
<tr><th>Request parameters are not valid</th><td>400</td></tr>
<tr><th>A user with this email already exists.</th><td>409</td></tr>
<tr><th>Unexpected error</th><td>500</td></tr>
</table>

#### Body

<table>
<tr><th>Content-Type</th><td>application/json</td></tr>
</table>

| Property | Nullable | Type   | Description                          | Example |
| -------- | -------- | ------ | ------------------------------------ | ------- |
| $ (root) | No       | Object |                                      |         |
| $.**id** | No       | Number | Unique ID of the newly created user. | 123     |

#### Example Response

```
HTTP/2 201
content-type: text/event-stream; charset=utf-8

{
    "id": 123
}
```

## 2. GET /users/{user id}

Get a single user's profile data by their ID.

### Request

<table>
<tr><th>Method</th><td>GET</td></tr>
<tr><th>Path</th><td>/users/{user_id}</td></tr>
</table>

#### Path Parameters

| Parameter            | Validation             | Description                                        | Example |
| -------------------- | ---------------------- | -------------------------------------------------- | ------- |
| /users/**{user_id}** | Must be an integer > 0 | ID of the user whose profile you wish to retrieve. | 123     |

#### Example Request

```
GET /users/123 HTTP/2
Host: 127.0.0.1:8080
```

### Response

#### Status

<table>
<tr><th>User profile was successfully retrieved</th><td>200</td></tr>
<tr><th>Request parameters are not valid</th><td>400</td></tr>
<tr><th>User with this ID does not exist</th><td>404</td></tr>
<tr><th>Unexpected error</th><td>500</td></tr>
</table>

#### Body

<table>
<tr><th>Content-Type</th><td>application/json</td></tr>
</table>

| Property          | Nullable | Type   | Description                                                                   | Example                   |
| ----------------- | -------- | ------ | ----------------------------------------------------------------------------- | ------------------------- |
| $ (root)          | No       | Object |                                                                               |                           |
| $.**id**          | No       | Number | Unique ID of the user.                                                        | 123                       |
| $.**first_name**  | No       | String | The user's given name.                                                        | John                      |
| $.**middle_name** | Yes      | String | The user's middle name.                                                       | Robert                    |
| $.**last_name**   | No       | String | The user's family name.                                                       | Doe                       |
| $.**age**         | No       | Number | User's age in years.                                                          | 30                        |
| $.**email**       | No       | String | The user's email address.                                                     | john.r.dough@example.com  |
| $.**created_at**  | No       | String | When this record was created. In RFC3339 format: "yyyy-MM-ddThh:mm:ssZ"       | 2026-09-02T15:45:00+09:00 |
| $.**updated_at**  | No       | String | When this record was last modified. In RFC3339 format: "yyyy-MM-ddThh:mm:ssZ" | 2026-09-02T15:45:00+09:00 |

#### Example Response

```
HTTP/2 200
content-type: text/event-stream; charset=utf-8

{
    "id": 123,
    "first_name": "John",
    "middle_name": "Robert",
    "last_name": "Doe",
    "age": 30,
    "email": "john.r.dough@example.com",
    "created_at": "2026-09-02T15:45:00+09:00",
    "updated_at": "2026-09-02T15:45:00+09:00"
}
```

## 3. PUT /users/{user id}

Update an existing user.

The same email address cannot be registered twice.
An error will be returned if the update gives this user another existing user's email address.

### Request

<table>
<tr><th>Method</th><td>PUT</td></tr>
<tr><th>Path</th><td>/users/{user id}</td></tr>
</table>

#### Path Parameters

| Parameter            | Validation             | Description                                      | Example |
| -------------------- | ---------------------- | ------------------------------------------------ | ------- |
| /users/**{user_id}** | Must be an integer > 0 | ID of the user whose profile you wish to update. | 123     |

#### Body

<table>
<tr><th>Content-Type</th><td>application/json</td></tr>
</table>

| Property          | Nullable | Type   | Description               | Example                  |
| ----------------- | -------- | ------ | ------------------------- | ------------------------ |
| $ (root)          | No       | Object |                           |                          |
| $.**first_name**  | No       | String | The user's given name.    | John                     |
| $.**middle_name** | Yes      | String | The user's middle name.   | Robert                   |
| $.**last_name**   | No       | String | The user's family name.   | Doe                      |
| $.**age**         | No       | Number | User's age in years.      | 30                       |
| $.**email**       | No       | String | The user's email address. | john.r.dough@example.com |

#### Example Request

```
PUT /users/123 HTTP/2
Host: 127.0.0.1:8080
Content-Type: application/json
Content-Length: 26

{
    "first_name": "John",
    "middle_name": "Robert",
    "last_name": "Doe",
    "age": 30,
    "email": "john.r.dough@example.com"
}
```

### Response

#### Status

<table>
<tr><th>User was successfully updated.</th><td>204</td></tr>
<tr><th>Request parameters are not valid</th><td>400</td></tr>
<tr><th>No user with this ID exists.</th><td>404</td></tr>
<tr><th>Email was updated to another user's email.</th><td>409</td></tr>
<tr><th>Unexpected error</th><td>500</td></tr>
</table>

#### Example Response

```
HTTP/2 204
```

## 4. DELETE /users/{user id}

Delete an existing user.

### Request

<table>
<tr><th>Method</th><td>DELETE</td></tr>
<tr><th>Path</th><td>/users/{user id}</td></tr>
</table>

#### Path Parameters

| Parameter            | Validation             | Description                                      | Example |
| -------------------- | ---------------------- | ------------------------------------------------ | ------- |
| /users/**{user_id}** | Must be an integer > 0 | ID of the user whose profile you wish to update. | 123     |

#### Example Request

```
DELETE /users/123 HTTP/2
Host: 127.0.0.1:8080
```

### Response

#### Status

<table>
<tr><th>User was successfully deleted.</th><td>204</td></tr>
<tr><th>Request parameters are not valid</th><td>400</td></tr>
<tr><th>No user with this ID exists.</th><td>404</td></tr>
<tr><th>Unexpected error</th><td>500</td></tr>
</table>

#### Example Response

```
HTTP/2 204
```
