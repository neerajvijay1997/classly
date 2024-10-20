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