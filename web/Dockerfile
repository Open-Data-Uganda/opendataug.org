FROM node:20-alpine AS build

RUN apk add --no-cache libc6-compat

WORKDIR /app

COPY package.json package-lock.json ./

COPY . .

RUN npm install --omit=dev --production=true

FROM node:20-alpine AS builder

WORKDIR /app

COPY --from=build /app/node_modules ./node_modules

COPY . .

RUN NODE_ENV=production npm run build && npm prune --production


FROM node:20-alpine AS runner

WORKDIR /app

ENV NODE_ENV production

RUN addgroup --system --gid 1001 nodejs && adduser --system --uid 1001 nextjs

COPY --from=builder /app/next.config.js ./

COPY --from=builder --chown=nextjs:nodejs /app/.next/standalone ./

COPY --from=builder --chown=nextjs:nodejs /app/.next/static ./.next/static

USER nextjs

EXPOSE 3000

ENV PORT 3000

ENV NEXT_TELEMETRY_DISABLED 1

CMD ["node", "server.js"]