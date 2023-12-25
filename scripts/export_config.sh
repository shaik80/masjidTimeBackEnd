#!/bin/bash

if [ $# -eq 0 ]; then
    echo "Usage: $0 <config_file>"
    exit 1
fi

config_file=$1

# Parse config.yaml and export variables
export APP_NAME=$(yq eval '.app.name' "$config_file")
export APP_ENV=$(yq eval '.app.environment' "$config_file")
export APP_DEBUG=$(yq eval '.app.debug' "$config_file")

export SERVER_ADDRESS=$(yq eval '.server.address' "$config_file")
export SERVER_PORT=$(yq eval '.server.port' "$config_file")

export DB_DRIVER=$(yq eval '.database.driver' "$config_file")
export DB_HOST=$(yq eval '.database.host' "$config_file")
export DB_PORT=$(yq eval '.database.port' "$config_file")
export DB_USERNAME=$(yq eval '.database.username' "$config_file")
export DB_PASSWORD=$(yq eval '.database.password' "$config_file")
export DB_NAME=$(yq eval '.database.dbname' "$config_file")
export DB_MAX_IDLE_CONNS=$(yq eval '.database.maxIdleConns' "$config_file")
export DB_MAX_OPEN_CONNS=$(yq eval '.database.maxOpenConns' "$config_file")

export LOG_LEVEL=$(yq eval '.logging.level' "$config_file")
export LOG_FILE_ENABLED=$(yq eval '.logging.file.enabled' "$config_file")
export LOG_FILENAME=$(yq eval '.logging.file.filename' "$config_file")

# Print exported variables for verification
echo "Exported variables:"
env | grep '^APP_\|^SERVER_\|^DB_\|^LOG_'
