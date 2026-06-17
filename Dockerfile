FROM golang:1.26.3

WORKDIR /app

COPY . .

RUN go mod tidy

WORKDIR /app/database/postgresqldatabase/example

RUN go build -o example .

EXPOSE 8080

CMD ["/app/database/postgresqldatabase/example/example"]


#FROM golang:1.26

#WORKDIR /app

#COPY . .

#RUN go mod tidy

#WORKDIR /app/authentications

# RUN go build -o authapp .

#EXPOSE 8080

#CMD ["/app/authentications/authapp"]


#FROM	Base image
#WORKDIR	Set working folder
#COPY	Copy files
#RUN	Execute build commands
#EXPOSE	Declare port
#CMD	Start app