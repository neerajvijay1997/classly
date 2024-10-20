# Development Guide for Classly

Welcome to the Classly development guide! This document outlines the steps to set up the development environment, install dependencies, and run the application.

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- **Go**: Version 1.22.5
  - You can download and install Go from the official [Go website](https://golang.org/dl/).

## Getting Started

Follow these steps to set up the project on your local machine:

### Clone the Repository

Open your terminal and run the following command to clone the repository:

```bash
git clone git@github.com:neerajvijay1997/classly.git
```

### Navigate to the Project Directory

Change to the project directory:

```bash
cd classly
```

### Install Dependencies

Use the following command to install the necessary dependencies:

```bash
go mod tidy
```

### Build the Application

To build the application, run:

```bash
go build -o classly-app
```

### Run the Application

You can run the application with the following command:

```bash
./classly-app
```
The server is currently running on port 8080.

# Classly API Usage

Here are the steps to interact with the Classly API:

### Check Classly Version

To see the current version of Classly, run:

```bash
curl -X GET http://localhost:8080/version

# Expected Output:
# {"version":"Classly-v0.1.0"}
```

### Register User

You can register a user as follows:

```bash
curl -X POST http://localhost:8080/signup -H "Content-Type: application/json" -d '{"name": "John Doe", "email": "john.doe@example.com"}'

# Expected Output:
# {"message":"User registered successfully","user_name":"tA3Qj"}
```

### Check User Registration

To check whether a user is registered, run the following command, replacing tA3Qj with the desired user_name:

```bash
curl -X GET http://localhost:8080/user/tA3Qj

# Expected Output:
# {"user_name":"tA3Qj","name":"John Doe","email":"john.doe@example.com"}
```

### Create a Class

User can create a class with the following command:

```bash
curl -X POST http://localhost:8080/classes -H "Content-Type: application/json" -d '{
  "user_name": "tA3Qj",
  "class_name": "Zumba Class",
  "start_date": "2024-10-20",
  "end_date": "2024-10-25",
  "capacity": 30
}'

# Expected Output:
# {"message":"Class created successfully","class_id":"u2GWN"}
```

### Check Available Classes

Users can check all available classes by running:

```bash
curl -X GET http://localhost:8080/all-classes

# Expected Output:
# [{"id":"u2GWN","class_name":"Zumba Class","description":"","class_provider_user_name":"tA3Qj","start_date":"2024-10-20T00:00:00Z","end_date":"2024-10-25T00:00:00Z","capacity":30}]
```

### Book a Class

User can book the class by running:

```bash
curl -X POST http://localhost:8080/bookings -H "Content-Type: application/json" -d '{
  "user_name": "qH8OU",
  "class_id": "u2GWN",
  "booking_date": "2024-10-21"
}'

# Expected Output:
# {"message":"Class booked successfully","class_session_id":"u2GWN#2024-10-21"}
```

### View Booked Classes

Users can view their booked classes by using the following command, substituting tA3Qj with the appropriate user_name:

```bash
curl -X GET http://localhost:8080/booked-classes/tA3Qj

# Expected Output:
# [{"id":"u2GWN","class_name":"Zumba Class","description":"","class_provider_user_name":"tA3Qj","start_date":"2024-10-20T00:00:00Z","end_date":"2024-10-25T00:00:00Z","capacity":30,"Sessions":["2024-10-21T00:00:00Z"]}]
```

### Check Class Status

Class creators can check the status of their created classes by using the following command, replacing tA3Qj with the relevant user_name:

```bash
curl -X GET http://localhost:8080/classes-status/tA3Qj

# Expected Output:
# [{"id":"u2GWN","class_name":"Zumba Class","description":"","class_provider_user_name":"tA3Qj","start_date":"2024-10-20T00:00:00Z","end_date":"2024-10-25T00:00:00Z","capacity":30,"Sessions":{"2024-10-21T00:00:00Z":[{"user_name":"qH8OU","name":"Jane Smith","email":"jane.smith@example.com"}]}}]
```