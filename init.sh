#!/bin/sh

# Wait for the database to be ready
echo "Waiting for Postgres..."
while ! nc -z db 5432; do
  sleep 1
done
echo "Postgres is up - executing command"

# Run migrations
./main migrate

# Run the application
./main
