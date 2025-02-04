FROM node:20-alpine AS builder

WORKDIR /app

COPY package*.json ./

RUN npm ci

COPY . .

RUN npm run build

FROM node:20-alpine

WORKDIR /app

COPY package*.json ./

RUN npm ci --production

COPY --from=builder /app/build ./build

COPY --from=builder /app/static ./static

EXPOSE 3000

CMD ["node", "build"] 