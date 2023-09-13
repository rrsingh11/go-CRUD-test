#!/bin/bash

# Drop the MongoDB database
mongosh contact_book --eval "db.dropDatabase()"

go run .