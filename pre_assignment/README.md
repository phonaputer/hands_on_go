# "Hands-on-Go" Pre-assignment

Hi! Thank you for your interest in joining the Hands on Go workshop.
In order to learn the concepts presented in the workshop, attendees need to have a certain minimum knowledge of Go 
(and the other technologies used in the workshop).
This pre-assignment is intended to verify that the attendee has this knowledge.

### Table of Contents
1. [What should I know before doing this assignment?](#pre-knowledge)
2. [What is the assignment?](#assignment)
3. [How to run the integration tests and verify my code?](#verify)
4. [How to submit my assignment?](#submit)

# What should I know before doing this assignment? <a name="pre-knowledge"></a>

The knowledge which this pre-assignment is intended to check is the following:

### Attendees should have...
* ...the ability to write a simple Go program from scratch.
* ...basic familiarity with the Go CLI (`go build`, `go run`, `go test`, etc.).
* ...basic familiarity with Go HTTP standard libraries  (`http.Handler` interface and `http.Server` struct type).
* ...basic familiarity with the Go `encoding/json` standard library.
* ...basic familiarity with at least one data store client library.
* ...the ability to run Docker containers on their local device.
* ...the ability to write simple SQL queries to SELECT and INSERT.
* ...an IDE for Go set up on their local device.

Also, just to add, this is what you don't need to know:
#### Attendees do NOT need to have...

* ...detailed understanding of Go HTTP handlers, JSON serialization, data store client libraries, etc.
* ...knowledge of how to write Dockerfiles or create new Docker images.
* ...detailed knowledge of MySQL or SQL queries.
* ...knowledge of patterns for structuring Go applications (this is the topic you will learn about in the workshop).

## Wait a minute - I don't know some of these things!

No problem! 
If you don't know about any of the topics from the "Attendees should have..." list 
- this assignment is your opportunity to study up on them!

As long as you can complete this assignment on your own without help, you should be ready for the course.
If you don't already know the above, just go ahead and learn it!

Here are some helpful links for learning about some of the more Go related topics above:

* Go HTTP server: https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go
* Go JSON serialization: https://www.sohamkamani.com/golang/json/
* Go standard library database SQL package: http://go-database-sql.org/index.html
* One MySQL driver library you could use with the standard lib SQL package: https://github.com/go-sql-driver/mysql#usage

For IDEs, there are a few different options. Some popular choices are the following:

1. Visual Studio Code with Go plugin (https://code.visualstudio.com/docs/languages/go)
   1. This is a good free option.
2. GoLand (https://www.jetbrains.com/go/promo/)
   1. A Go IDE from JetBrains.
3. IntelliJ with Go plugin (https://plugins.jetbrains.com/plugin/9568-go)
   1. If your team is already paying for your IntelliJ, just add the Go plugin to get the same functions as GoLand.

For Docker and Docker-Compose, you can install them by following the instructions at the following link.
Installing in this way will include both `docker` and `docker-compose`: https://docs.docker.com/get-docker/

# What is the assignment? <a name="assignment"></a>

The pre-assignment for this course is to write a simple HTTP server application with one endpoint for "Get User." 
This endpoint should select a row from a MySQL table by ID and return the data in JSON format in the response body.
This application does not need to handle creation, deletion, or update functionality. Just selects/reads.

To verify your application is working correctly, please run the application 
and then execute the integration tests packaged in this repository.
If the tests are passing, your code works.

You may use any libraries you like and structure the code however you please.
You will not be graded on cleanness/consistency/etc. only correctness and understanding of Go.

The specific requirements for the server are described below. 
If you have any questions about these requirements, please feel free to contact the trainer to ask.

### The HTTP Server

The HTTP server should listen on port 8080 of localhost.

### The Get User Endpoint

The endpoint to get a User should conform to the following spec.

| URL                            | HTTP Method                         |
|--------------------------------|-------------------------------------|
| http://localhost:8080/get-user | any (the integration tests use GET) |

#### Query String Parameters

| Parameter | Format  | Description                          | Optional |
|-----------|---------|--------------------------------------|----------|
| id        | integer | Unique integer ID of a User in MySQL | No       |


#### Example Request
```
curl 'localhost:8080/get-user?id=123'
```

#### Response Code
Success - `200` (with JSON body)

All error cases - `500` (no response body)

#### Success Response Body JSON Format

| JSON Key   | JSON Type | Description                             | Optional |
|------------|-----------|-----------------------------------------|----------|
| id         | Number    | Unique integer ID of a User in MySQL    | No       |
| first_name | String    | Human language given name of the User.  | No       |
| last_name  | String    | Human language family name of the User. | No       |

#### Example Success Response

```
200 OK
{
   "id" : 123,
   "first_name" : "Rakuten",
   "last_name" : "Gopher"
}
```

### The MySQL Database

The MySQL database comes packaged as a Docker container which can be run with Docker Compose.

It contains only one database with one table.

#### How to run MySQL

The MySQL Dockerized setup can be run by navigating to the `deployments/local` directory and running the 
following command:

```
$ docker-compose up
```

This should be enough to get it running. 
If this does not work, please contact the trainer.

Now you should have MySQL running on port `3306` of `localhost`.

#### Schema 

The MySQL database/schema you should use is `hands_on_go`.

#### Tables

There is only one table in the `hands_on_go` schema. This is the `users` table.

It has the following structure:

```sql
CREATE TABLE `hands_on_go`.`users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `first_name` varchar(100) NOT NULL,
  `last_name` varchar(100) NOT NULL,
  `age` int(10) unsigned NOT NULL,
  `phone_number` varchar(255) NOT NULL,
  `phone_verification_status` tinyint(1) NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);
```

For this pre-assignment, the only important columns are `id`, `first_name`
, and `last_name`. These are the columns whose data should be returned in the JSON body.

#### MySQL Users

The Dockerized MySQL comes with one user `dockeruser` for whom the password is `dockerpass`.
This user has the required SELECT permissions for completing this assignment.

# How to run the integration tests and verify my code? <a name="verify"></a>

The integration tests can be run with the following command:

```
# in the root directory of this repository
go test -v test/integration_test.go
```

Note that you need to run your HTTP server application and have it listen on port 8080 of localhost 
before starting the tests. 
Otherwise the tests will not be able to connect to your server and they will fail.

If all tests are passing, then your code should be working!

# How to submit my assignment? <a name="submit"></a>

Please push your code to a branch on the origin of this repository.
The branch name should include your name at the beginning followed by a forward slash, then whatever name you would like: `${YOUR_NAME}/${BRANCH_NAME}`. 
For example: `jane.doe/pre_assignment`.

Once your branch is pushed, open a pull request from your branch into `master` branch.
Finally, please send a link to the pull request to the trainer for verification.

No worries if your code is looking a dirty or unmaintainable, 
improving this is one of the learning targets of the workshop.  