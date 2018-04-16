#!/bin/bash

echo "Creating test database"

mongo --eval "use test"

mongo --eval "db.roles.insert({'name':'master'})"

echo "test database created"