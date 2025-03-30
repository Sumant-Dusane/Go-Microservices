
#!/bin/bash
# Define all services
SERVICES=("user-service" "transfer-service")

# Package to install (pass it as an argument)
PACKAGE="$1"

# Check if a package is provided
if [ -z "$PACKAGE" ]; then
  echo "Usage: ./go_get.sh <package-name> \nPlease provide a package name"
  exit 1
fi

# Loop through each service and install the package
for SERVICE in "${SERVICES[@]}"; do
  echo "Installing $PACKAGE in $SERVICE..."
  (cd "./services/$SERVICE" && go get "$PACKAGE" && go mod tidy)
done

echo "✅ Installation complete!"
