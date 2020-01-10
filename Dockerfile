# Gobuilder
FROM golang:alpine as gobuilder
RUN apk add --no-cache git make

RUN mkdir -p /go/src/habits
WORKDIR /go/src/habits

COPY Makefile .
COPY go.mod   .
COPY go.sum   .
COPY cmd      ./cmd
COPY database ./database
COPY internal ./internal

RUN make install
RUN make build-api-production

# Jsbuilder
FROM node:lts-alpine as jsbuilder
RUN apk add --no-cache make
WORKDIR /app

COPY package*.json ./

ENV NODE_ENV=production
RUN npm install --no-optional

COPY babel.config.js .
COPY vue.config.js   .
COPY Makefile        .
COPY web/src         ./web/src
COPY web/public      ./web/public

RUN make build-js-production

# Final image
FROM alpine
ARG COMMIT
WORKDIR /app

COPY --from=gobuilder /go/src/habits/bin/*    .
COPY --from=gobuilder /go/src/habits/database ./database
COPY --from=jsbuilder /app/web/dist           ./web/dist

RUN echo ${COMMIT} > ./web/dist/version

EXPOSE 80
EXPOSE 443
ENV HABITS_ENVIRONMENT=production
CMD ["./habits", "serve"]
