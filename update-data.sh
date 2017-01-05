#!/bin/bash

# Election Database Installation Script
# ========================================
# PostgreSQL 9.1, 9.3 on Linux

# Setup Instructions
# ------------------
#
# Create a PostgreSQL user
# ref: Create a PostgreSQL user.md
#
# Open a terminal, then run the following commands:
#
# vi ~/.pgpass
#
# Add the following 2 lines to .pgpass (Without the first # on each line)
#    #hostname:port:database:username:password
#    localhost:5432:election:member:ru123*Member
#
# Save and close .pgpass
#
# chmod 600 ~/.pgpass
#
# Close and re-open the terminal
#
# chmod +x update-data.sh
#
# cd bin
#
# ./update-data.sh
#


# Variables
# ------------------------------------------------------------------------------
server=localhost
database=election
port=5432
username=member


# Drop and Create the DB as postgres (use sudo)
# -------------------------------------------------------------------------------
psql -h localhost -p 5432 -U member -w -d election < muslims-19.sql > logm19 2>&1
psql -h localhost -p 5432 -U member -w -d election < muslims-20.sql > logm20 2>&1
psql -h localhost -p 5432 -U member -w -d election < muslims-21.sql > logm21 2>&1
psql -h localhost -p 5432 -U member -w -d election < others.sql > logo 2>&1
psql -h localhost -p 5432 -U member -w -d election < others-19.sql > logo19 2>&1
psql -h localhost -p 5432 -U member -w -d election < others-20.sql > logo20 2>&1
psql -h localhost -p 5432 -U member -w -d election < others-21.sql > logo21 2>&1
