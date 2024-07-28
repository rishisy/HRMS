<div style="text-align: center;">
    <h1>Hospital Management System</h1>
    <img src="https://img.icons8.com/clouds/100/doctors-bag.png" alt="Logo" />
</div>

## Table of Contents
1. [Tech Stack](#tech-stack)
2. [API Reference](#api-reference)
3. [Project Structure](#project-structure)
4. [Installation](#installation)
5. [Docker](#docker)
6. [Additional Notes](#additional-notes)
7. [Conclusion](#conclusion)

The Hospital Management System is a web application built using Go (Golang) with the Gin framework for routing and GORM as the ORM library for Postgre Database interactions. This system manages doctors and patients, allowing for the addition, updating, searching, and retrieval of records.

## Tech Stack

**Server:** Golang , Gin , GORM

**Database:** PostgresSQL

**Testing:** Postman , Insomnia , SoapUI

## API Reference

### Create Doctor

```http
POST /doctor/
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name`    | `string` | **Required**. Name of the doctor  |
| `contact_no` | `string` | **Required**. Contact number of the doctor |

### Get Doctor

```http
GET /doctor/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of the doctor to fetch |

### Update Doctor

```http
PATCH /doctor/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of the doctor to update |
| `contact_no` | `string` | **Required**. New contact number of the doctor |

### Create Patient

```http
POST /patient/
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name`    | `string` | **Required**. Name of the patient |
| `contact_no` | `string` | **Required**. Contact number of the patient |
| `address` | `string` | **Required**. Address of the patient |
| `doctor_id` | `string` | **Required**. Id of the doctor assigned to the patient |

### Get Patient

```http
GET /patient/:id  
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of the patient to fetch |

### Update Patient

```http
PATCH /patient/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of the patient to update |
| `contact_no` | `string` | **Optional**. New contact number of the patient |
| `address` | `string` | **Optional**. New address of the patient |
| `doctor_id` | `string` | **Optional**. New doctor id assigned to the patient |

### Get Patients by Doctor Id

```http
GET /fetchPatientByDoctorId/:docterId
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `docterId`| `string` | **Required**. Id of the doctor to fetch patients for |

This endpoint returns a list of patients being monitored by the specified doctor.

## Project Structure

```bash
├── db
│   ├── dbConfig.go
│   └── uuid.go
├── doctors
│   ├── doctors.go
│   ├── handler.go
│   └── routes.go
├── patients
│   ├── handler.go
│   ├── patients.go
│   └── routes.go
├── go.mod
├── go.sum
├── main.go
├── readme.md
└── test.txt
```

## Installation

### Prerequisites
Before you begin, ensure you have the following installed on your machine:
- **Go**: Version 1.18 or later
- **Docker**: For containerization
- **Database**: PostgreSQL

### Steps to Install

-  **Clone the Repository**
   Open your terminal and run the following command to clone the repository:
   ```bash
   git clone https://github.com/rishisy/HRMS.git
   cd HRMS
   ```

-  **Set Up Environment Variables**
   Create a `.env` file in the root directory of the project and configure your database connection settings:
   ```plaintext
   DB_STRING="postgres://YourUserName:YourPassword@YourHostname:5432/YourDatabaseName"
   PORT=8080
   ```

- **Run Database Migrations**
  Ensure that you have the necessary migrations set up in your GORM models. Run the migrations to create the required tables in your database.
  ```bash
  go run migrations/migrate.go
  ```


- **Build the Application**
  Compile the application by running:
   ```bash
      go build -o HRMS
   ``` 

-  **Run the Application**
   Start the application using the following command:
   ```bash
   ./HRMS
   ```

## Docker

### 1. Clone the Repository

Clone this repository to your local machine and Load Env Variables:

```bash
git clone https://github.com/rishisy/HRMS
cd HRMS
```

### 2. Build the Docker Image

To build the Docker image for the application, run the following command in the root directory of the project (where your Dockerfile is located):

```bash
docker build -t go-hrms-docker . 
```

### 3. Run the Docker Container

Once the image is built, you can run the Docker container and map the container's port to your host machine. Use the following command:

```bash
docker image ls
docker run -p 8080:8080 go-hrms-docker
```

In this command:
- `-p 8080:8080` maps port 8080 on your host to port 8080 in the container.

### 4. Access the API Endpoints

After running the container, you can access the API endpoints from your laptop. Open a web browser or use a tool like `curl` or Postman to make requests to:

```
http://localhost:8080/patient/abcde
```

## Additional Notes

- Ensure that your firewall settings allow traffic on port 8080.
- If you encounter any issues, verify that the Docker daemon is running and that your Docker installation is up to date.

## Conclusion

You have successfully built and run your Dockerized Go application! For further modifications or enhancements, feel free to explore the code and make changes as needed.

Citations:
[1] https://img.icons8.com/clouds/100/doctors-bag.png