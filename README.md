

# Snippetbox: A Web App for Sharing Text Snippets

**Snippetbox** is a web application where users can paste, save, and share text snippets—similar to services like Pastebin or GitHub’s Gists.

### Key Features

1. **Simple Setup and Configuration**  
   Set up the project, manage configurations, and control the flow of requests. Learn the basics of web app development, routing, and handling HTTP headers.
  
2. **Database-Driven Responses**  
   Store and retrieve user data using MySQL, which is integrated into the app to provide persistence and efficient querying.

3. **Dynamic HTML Templates**  
   Create, cache, and display dynamic HTML pages, with robust error handling and custom template functions.

4. **Middleware for Security and Logging**  
   Implement middleware to add security headers, request logging, and panic recovery, enhancing application stability.

5. **Advanced Routing and Form Handling**  
   Utilize clean URLs, process forms, validate user input, and manage session states for a seamless experience.

6. **User Authentication and Security**  
   Includes signup, login, logout, user roles, CSRF protection, and HTTPS configuration.

7. **End-to-End Testing**  
   Comprehensive testing features, including unit, integration, and end-to-end testing for HTTP handlers, forms, and middleware.

### Database-Driven Responses

For Snippetbox to store and retrieve user-entered data, we use MySQL as our database. This allows us to save snippets and access them dynamically during runtime. You’ll learn about setting up MySQL, database connection pooling, and crafting efficient SQL queries.

---

### Application Endpoints

The following routes provide all interactions users need within Snippetbox:

| HTTP Method | Route                | Description                           |
|-------------|-----------------------|---------------------------------------|
| `GET`       | `/static/*filepath`   | Serves static files (CSS, JS, etc.).  |
| `GET`       | `/`                   | Homepage with a list of snippets.     |
| `GET`       | `/snippet/view/:id`   | View a specific snippet.              |
| `GET`       | `/user/signup`        | User signup form.                     |
| `POST`      | `/user/signup`        | Submits the signup form.              |
| `GET`       | `/user/login`         | User login form.                      |
| `POST`      | `/user/login`         | Submits the login form.               |
| `GET`       | `/snippet/create`     | Snippet creation form (auth required).|
| `POST`      | `/snippet/create`     | Submits the new snippet (auth required). |
| `POST`      | `/user/logout`        | Logs the user out.                    |
| `GET`       | `/ping`               | Health check route for testing.       |

---

### Getting Started

1. **Install Dependencies**  
   Clone the repository, set up the database, and configure your environment settings.

2. **Run the App**  
   Start the app using the main application file, which will load all configured settings and routes.

3. **Testing**  
   Use the `/ping` endpoint for health checks and explore available testing utilities to ensure everything works as expected.
