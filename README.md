# BuildGram

BuildGram is a full-stack application featuring a Go-based backend (using Gin) and a React-based frontend (using Vite).

## Prerequisites

- [Docker](https://www.docker.com/) and Docker Compose (Required for the easiest one-click setup)
- *Optional (for manual setup without Docker)*:
  - [Go](https://golang.org/dl/) (1.23 or later)
  - [Node.js](https://nodejs.org/) (18.x or later)
  - [PostgreSQL](https://www.postgresql.org/)

## Build and Run Instructions (Docker)

The entire application stack (Database, Backend, and Frontend) is fully containerized. You can build and start the app without needing to configure databases or install local runtimes!

1. Open your terminal at the root of the project (`BuildGram/`).
2. **Build the Docker images:**
   ```bash
   docker-compose build
   ```
3. **Run the application:**
   ```bash
   docker-compose up
   ```
   *(Tip: You can use `docker-compose up -d` to run it in the background)*

**That's it!** Docker will start all the necessary services for you.

- The **Frontend** will be available at `http://localhost:5173`
- The **Backend API** will be running at `http://localhost:8080`
- The **PostgreSQL Database** runs internally but is exposed on port `5432` if you need to connect to it directly (credentials: user=`buildgram_user`, pass=`buildgram_password`, db=`buildgram`).

> To stop the application, press `Ctrl+C` in your terminal or run `docker-compose down` from another terminal window.

---

## Manual Setup (Without Docker)

If you prefer to run the application manually on your host machine, follow these steps:

### 1. Database
1. Create a PostgreSQL database named `buildgram`.
2. Update the database credentials in the `backend/.env` file.

### 2. Backend
1. Navigate to the backend directory: `cd backend`
2. Download modules: `go mod download`
3. Start the server: `go run cmd/api/main.go`

### 3. Frontend
1. Navigate to the frontend directory: `cd frontend`
2. Install dependencies: `npm install`
3. Start the development server: `npm run dev`
