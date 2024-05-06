FROM oven/bun:latest as base

WORKDIR /app/frontend

COPY package.json bun.lockb ./

COPY . .

RUN bun install

EXPOSE 3000

CMD ["bun", "dev"]