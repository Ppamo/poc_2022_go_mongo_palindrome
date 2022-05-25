#!/bin/sh


# Create user
printf -- "- Creating user:\n"
mongo --username $MONGO_INITDB_ROOT_USERNAME \
	--password $MONGO_INITDB_ROOT_PASSWORD \
	--authenticationDatabase $AUTHDB << EOF
db.createUser({ user: "$DB_USERNAME", pwd: "$DB_PASSWORD", roles: [ { role: 'readWrite', db: "$DB" } ]})
EOF

# Import data
printf -- "- Creating user:\n"
mongoimport \
	--username $MONGO_INITDB_ROOT_USERNAME \
	--password $MONGO_INITDB_ROOT_PASSWORD \
	--authenticationDatabase $AUTHDB \
	--db $DB \
	--collection $COLLECTION \
	--mode upsert \
	--file /docker-entrypoint-initdb.d/01-products.json
