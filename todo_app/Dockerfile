FROM golang:1.19.1-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /ToDoBot

EXPOSE 5000
CMD [ "/ToDoBot" ]