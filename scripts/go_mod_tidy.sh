#!/bin/bash

# Define all services
SERVICES=("user-service" "transfer-service") # Add more services if needed

# Loop through each service and remove the package
for SERVICE in "${SERVICES[@]}"; do
  (cd "./services/$SERVICE" && go mod tidy)
done

echo "✅ Package updated!"
