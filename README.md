postman documentation:[https://web.postman.co/documentation/37347478-35789e5c-e820-43ef-8bd7-202e68ba56a9/publish?workspaceId=972f6ba3-65ab-4b2c-89ea-d981fe49d911](#)



# Task Management API - README

Welcome to the Task Management API, a RESTful API built using the Go Gin framework. This project provides endpoints for managing tasks, including creating, retrieving, updating, and deleting tasks.

## Table of Contents

1. [Overview](#overview)
2. [Prerequisites](#prerequisites)
3. [Installation](#installation)
4. [Running the API](#running-the-api)
5. [API Endpoints](#api-endpoints)
6. [License](#license)

## Overview

The Task Management API is designed to be simple and efficient, providing a way to manage tasks via a REST API. It uses the Gin framework for handling HTTP requests and responses. The application is now integrated with MongoDB for persistent data storage.

## Prerequisites

Before setting up the project, ensure you have the following installed on your machine:

1. **Go**: Version 1.16 or higher. Download it from [golang.org](https://golang.org/dl/).
2. **Git**: For version control.
3. **MongoDB**: For database storage. You can download and install it from [mongodb.com](https://www.mongodb.com/try/download/community).

## Installation

### 1. Clone the Repository

First, clone the repository to your local machine:

```sh
git clone https://github.com/saleamlakw/BackEnd_Assessment.git
cd BackEnd_Assessment
```

### 2. Install Dependencies

Install the necessary Go packages:

```sh
go mod tidy
```

This command will download all dependencies specified in the `go.mod` file.

### 3. Set Up MongoDB

Ensure MongoDB is running on your local machine or accessible remotely. By default, the application connects to MongoDB running on `localhost` at the default port `27017`. The application expects a database named `taskmanager` and a collection named `tasks`. If these do not exist, they will be created automatically when the API is first accessed.

## Running the API

To start the API server, use the following command:

```sh
go run main.go
```

The server will start on `http://localhost:8080` by default. If you've set a different port using environment variables, it will use that port instead.

## API Endpoints

The following endpoints are available in the Task Management API:
## License

This project is licensed under the MIT License. For more details, refer to the [LICENSE](LICENSE) file in the repository.

