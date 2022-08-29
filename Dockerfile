FROM cosmtrek/air
WORKDIR /app
COPY . /app
RUN go get ./...