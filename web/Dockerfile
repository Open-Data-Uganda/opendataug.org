FROM node:20-alpine AS build

WORKDIR /app

COPY package*.json ./

ENV NODE_OPTIONS="--max_old_space_size=2048"

ENV NODE_ENV=production

RUN npm ci --only=production

COPY . .

RUN npm run build

FROM nginx:alpine

COPY --from=build /app/dist /usr/share/nginx/html
COPY ./nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 3000

CMD ["nginx", "-g", "daemon off;"] 