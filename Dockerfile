ARG COMMIT="N/A"

FROM golang:alpine AS build-back
WORKDIR /app
ARG COMMIT
COPY . .
RUN go build -o wg-gen-web -ldflags="-X 'github.com/jasoryeh/wg-gen-web/version.Version=${COMMIT}'" github.com/jasoryeh/wg-gen-web

FROM node:10-alpine AS build-front
WORKDIR /app
COPY ui/package*.json ./
RUN npm install
COPY ui/ ./
RUN npm run build

FROM alpine
WORKDIR /app
COPY --from=build-back /app/wg-gen-web .
COPY --from=build-front /app/dist ./ui/dist
COPY .env .
RUN chmod +x ./wg-gen-web
RUN apk add --no-cache ca-certificates
EXPOSE 8080

CMD ["/app/wg-gen-web"]