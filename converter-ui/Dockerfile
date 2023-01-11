FROM node:16-alpine AS builder

COPY . /app
WORKDIR /app

RUN yarn install \
    && yarn build

# Path: converter-ui/Dockerfile
FROM nginx:1.23-alpine

COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80