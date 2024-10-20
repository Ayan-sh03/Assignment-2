## Realtime Weather Rollup and Aggregation Application

This is a Realtime Weather Rollup and Aggregation Application built with Golang for the backend, PostgreSQL as the database, and React.js for the frontend.

### Prerequisites

Before getting started, ensure you have the following installed:

- **Golang** (version 1.18 or above)
- **PostgreSQL** (Ensure the service is running)
- **Node.js** (version 14.x or above)
- **npm** or **yarn**
- **Make** (for building the backend)

### Getting Started

#### 1. Setting up the Backend

1. Navigate to the backend directory:

   ```bash
   cd Realtime-backend
   ```

2. Create a .env file by copying the example:

   ```bash
   cp .env.example .env
   ```

3. Update the .env file with your PostgreSQL connection details and required Config variables.

4. Build the backend:

```bash
make
```

5. To run the backend server:

```bash
 make run
```

this will start the backend server on port 8080.

6. Setting up the Frontend

7. Navigate to the frontend directory:

8. Install the required dependencies:

```bash
npm install
```

9. Start the frontend development server:

```bash
npm run dev
```

This will start the frontend development server on port 5173.

### Usage

Once both the backend and frontend are running:

Access the frontend at http://localhost:5173.
The backend API will be accessible at http://localhost:8080.
