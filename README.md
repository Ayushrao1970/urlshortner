# Go URL Shortener

A simple URL shortener built with Go. It takes long URLs and gives you a short one.

# How to run

You’ll need two terminals

*Terminal 1 -start the server


command - go run main.go

*Terminal 2 – For sending the curl request
command-{
curl -X POST http://localhost:3000/shorten\
  -H "Content-Type: application/json"\
  -d '{"url":"https://example.com"}' (You can put any link you want to shorten inside the url field)
}

