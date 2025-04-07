# Open Data Uganda

A modern web application for accessing and visualizing Uganda's open data, built with React, TypeScript, and Go.

## Project Structure

The project consists of two main components:

### Frontend (`/app`)
- Built with React, TypeScript, and Vite
- Uses Tailwind CSS for styling
- Modern UI components and responsive design

### Backend (`/backend`)
- Built with Go
- RESTful API architecture
- PostgreSQL database integration
- Authentication and authorization middleware

## Prerequisites

- Node.js (version specified in `.nvmrc`)
- Go (latest stable version)
- Docker and Docker Compose (for containerized deployment)
- PostgreSQL

## Getting Started

### Frontend Setup

1. Navigate to the app directory:
```bash
cd app
```

2. Install dependencies:
```bash
npm install
```

3. Create a `.env` file based on the example:
```bash
cp .env.example .env
```

4. Start the development server:
```bash
npm run dev
```

### Backend Setup

1. Navigate to the backend directory:
```bash
cd backend
```

2. Install Go dependencies:
```bash
go mod download
```

3. Create a `.env` file based on the example:
```bash
cp .env.example .env
```

4. Start the development server:
```bash
go run main.go
```

## Development

- Frontend runs on `http://localhost:5173` by default
- Backend API runs on `http://localhost:8080` by default
- Air is used for hot reloading in the backend

## Deployment

The project includes Docker configurations for both frontend and backend:

- Frontend: Uses Nginx to serve the built React application
- Backend: Uses a multi-stage build to create a minimal Go binary

To build and run with Docker:

```bash
# Build and run frontend
cd app
docker build -t ug-data-frontend .
docker run -p 80:80 ug-data-frontend

# Build and run backend
cd backend
docker build -t ug-data-backend .
docker run -p 8080:8080 ug-data-backend
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the terms specified in the LICENSE.md file.

## Contact

For any queries or support, please open an issue in the repository. 