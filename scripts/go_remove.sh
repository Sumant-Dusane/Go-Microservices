#!/bin/bash

# Define all services
SERVICES=("user-service" "transfer-service") # Add more services if needed

# Package to remove (pass it as an argument)
PACKAGE="$1"

# Check if a package is provided
if [ -z "$PACKAGE" ]; then
  echo "Usage: ./remove.sh <package-name> \nPlease provide package name"
  exit 1
fi

# Loop through each service and remove the package
for SERVICE in "${SERVICES[@]}"; do
  echo "Removing $PACKAGE from $SERVICE..."
  (cd "$SERVICE" && go get -u "$PACKAGE" && go mod tidy)
done

echo "✅ Package removal complete!"
