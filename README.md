# Online Job Portal

Welcome to the JP an online job portal! This web application allows job seekers and employers to interact with each other. 
Job seekers can register, create and manage their profiles, and apply for jobs. Employers can post job listings and manage their profiles. 
This project uses Go and Gin for the backend, and MongoDB as the database.

## Features

- **User Authentication:** Register and log in as an employee or employer.
- **Profile Management:** Employees can manage their profiles, including academic background, work experience, skills, and more. Employers can manage their profiles and job postings.
- **Job Listings:** Browse and apply for jobs posted by employers.
- **Post Jobs:** Employers can create and manage job postings.
- **Contact Form:** Users can send inquiries or feedback through the contact form.

## Project Structure

- **`main.go`**: The entry point of the application, sets up routes and middleware.
- **`handlers/employee/`**: Contains handler functions for employee-related operations such as profile management, academics, experience, etc.
- **`handlers/employer/`**: Contains handler functions for employer-related operations such as posting jobs and managing employer profiles.
- **`handlers/landing_pages/`**: Contains handler functions for landing pages and the contact form.
- **`handlers/register/`**: Handles user registration and login.
- **`templates/`**: Directory containing HTML templates for various pages (login, signup, profiles, job postings, etc.).

## Setup

### 1. Clone the Repository
``` bash
git clone https://github.com/4hm3d57/jp.git
cd online-job-portal

#installing dependencies
go mod tidy

# running the applicatoin
go run main.go
```

### 2. Setup MongoDB
- Ensure MongoDB is installed and running on your local machine. You can download and install MongoDB from the official MongoDB website.
- By default, MongoDB runs on mongodb://localhost:27017. If you are using a different configuration or a cloud MongoDB service, update your connection settings accordingly in your code.

### 3. Access the Application:
- Open your web browser and go to http://localhost:9000 to see the running application.


## Final Notes
- Thank you for checking out the Online Job Portal project! We hope you find it useful. If you encounter any issues or have suggestions for improvements, please feel free to reach out or contribute to the project.

Happy coding!







