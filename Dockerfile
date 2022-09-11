FROM golang

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /InterviewQuestions_API

EXPOSE 8080

CMD [ "/InterviewQuestions_API" ]