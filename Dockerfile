#FROM golang:1.16-alpine as BUILDER
#WORKDIR /app
#COPY . .
#RUN go build -o main main.go


FROM alpine:latest
WORKDIR /app
#COPY --from=BUILDER /app/main .
COPY main .
COPY app.env .

EXPOSE 8080
CMD ["/app/main"]
