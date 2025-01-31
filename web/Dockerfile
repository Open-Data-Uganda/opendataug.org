# Build stage
FROM node:20-alpine AS builder

WORKDIR /app

# Copy package files first to cache dependencies
COPY package*.json ./

# Install dependencies with clean install
RUN npm ci

# Copy the rest of the source code
COPY . .

# Build the application
RUN npm run build

# Production stage
FROM node:20-alpine

WORKDIR /app

# Copy package files for production dependencies
COPY package*.json ./

# Install only production dependencies
RUN npm ci --production

# Copy built assets from builder stage
COPY --from=builder /app/build ./build
COPY --from=builder /app/static ./static

EXPOSE 3000

CMD ["node", "build"] 