#!/bin/sh


# Create user
printf -- "- Creating user:\n"
mongo --username $MONGO_INITDB_ROOT_USERNAME \
	--password $MONGO_INITDB_ROOT_PASSWORD \
	--authenticationDatabase $AUTHDB $DB_NAME << EOF
db.createUser({ user: "$DB_USERNAME", pwd: "$DB_PASSWORD", roles: [ 'readWrite' ]})
EOF

# Import data
printf -- "- Importing data:\n"
mongoimport \
	--username $MONGO_INITDB_ROOT_USERNAME \
	--password $MONGO_INITDB_ROOT_PASSWORD \
	--authenticationDatabase $AUTHDB \
	--db $DB_NAME \
	--collection $COLLECTION \
	--mode upsert \
	--file /docker-entrypoint-initdb.d/01-products.json
