docker build -t hwr-timetable-bot:latest .
docker run --rm --env-file .env -it hwr-timetable-bot