# Summary bank account

This api allows to send by email, the summary of a group of transactions of a person.

## Local deploy

The application has a dockerfile, with which it can be deployed locally for testing or for development, but keep in mind 
that some environment variables are needed for the app to work correctly. 

You have the .env.example file, with which you can create an `.env` file, by filling in the corresponding values, 
these values belong to an authentication based on OAuth 2 of Gmail, with which the smtp management of sending mails is done.

It should be noted that these elements are required, since the application is based precisely on sending emails.

Here are a series of steps to run the application.

## Local deploy steps

From a terminal (preferably on Unix/Linux systems) run the command.

```Bash
$ cp .env.example .env
```
After supplying the login credentials to the smtp handler, you can move on to running the docker image.

In the root of the project execute.

```Bash
$ docker build -t summary-bank-account .
```

Some considerations to keep in mind, first, you must have docker installed on your system; if your system is windows or 
Macos, you need a docker runtime tool, you can use [Docker Desktop](https://www.docker.com/products/docker-desktop/) or 
[Colima](https://github.com/abiosoft/colima) (recommend).

After having the docker image of the project we can run it, in the following command we propose the use of port `8080`.

```Bash
$ docker run -p 8080:8080 summary-bank-account
```
It also has a makefile with which you can run the application just by using the command.

```Bash
$ make run-local
```

## Endpoints

The project has two endpoints, which are described below.

| Http Verb | URI                 | Description                                                                                                                                                                                   |
|-----------|---------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| GET       | /                   | Allows to verify the access to the app, contains the message hello world                                                                                                                      |
| POST      | send-summary/:email | Allows you to send the account summary, based on the `.csv` file stored in the project in the directory `/data`, this endpoint needs as a parameter the email to which the data will be sent. |

The send-data endpoint can respond to error messages due to validations or internal errors with the following structure.

Success case: status code `200 OK`

```JSON
{
  "message": "succcess send summary to address <email>"
}
```

In case of errors, cases `400 Bad Request`, `500 Internal Server Error`

```JSON
{
  "message": "<error message>"
}
```
## Project Architecture

A clean architecture was used, thinking mainly about scalability and the decoupling of the parts that make up the application.

The `Repository Service` pattern was used to be able to separate the service and data management layers, that is why we 
see that all the management of summary calculation and email management are separated from the data retrieval.
and email management are separated from the data collection.

The tests are based on Mock Testing, in which it is taken as an initiative not to use external dependencies that slow 
down the tests and only test the parts of the code for which they are being used. and test only the parts of the code for which they are being used.

If you can run all test in local, you can use this command.

```Bash
$ make test
```

## Online application.

To access the online app you can go to the following address: `http://18.119.134.105:8080/` in which you can test the endpoints in the same way as they are
as they are expressed in this documentation.

## Deploy Description

This application was deployed online using Amazon Web Services, specifically, AWS ECR and AWS ECS, this deployment will be for a limited time only.
limited time.
