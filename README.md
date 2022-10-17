# SOiCal

An iCal bridge for the Czech edutech provider Škola OnLine

## Stability

This project is still being worked on.
It works with my account, but I can't guarantee that it will work with yours, as the API was reverse-engineered.
Most of the values are also hard-coded, so it will for example only show events ranging from a week ago to a month in advance.

## Drama

So it looks like they blacklisted the IP of my server. This is weird, because I've only been doing three requests every hour or so, compared to one request every 5 minutes I did from my Raspberry Pi at home (which I ran for months and didn't get blocked). I will try to get in contact with them through my IT admin and will update this when I have more info.

## Building

First set up a Go environment

```sh
./build.sh
```

## Running

Use `go run ./cmd/server` or the Docker container.

## Adding the calendar

Use a calendar app to add the url `{basePath}/calendar/v1` where `{basePath}` is the path of your deployment.
Use basic HTTP authentication with your Škola OnLine credentials (like you'd use in the app).
