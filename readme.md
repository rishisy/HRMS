
![Logo](https://img.icons8.com/clouds/100/doctors-bag.png)

# Hospital Management System

The Hospital Management System is a web application built using Go (Golang) with the Gin framework for routing and GORM as the ORM library for Postgre Database interactions. This system manages doctors and patients, allowing for the addition, updating, searching, and retrieval of records.




## Tech Stack

**Server:** Golang , Gin , GORM

**Database:** PostgresSQL

**Testing:** Postman , Insomnia



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
├── server
│   ├── dbConfig.go
│   └── uuid.go
├── doctors
│   ├── doctors.go
│   ├── handler.go
│   └── routes.go
├── patients
│   ├── handler.go
│   ├── patients.go
│   └── routes.go
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
- **Database**: MySQL or PostgreSQL

### Steps to Install

-  **Clone the Repository**
   Open your terminal and run the following command to clone the repository:
   ```bash
   git clone https://github.com/yourusername/HRMS.git
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



### Notes:
- Remember to replace `yourusername` and other placeholder values in the `.env` file with your actual database credentials.
- Adjust any commands based on your specific setup or additional requirements.

