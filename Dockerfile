FROM golang:latest

# Create the directory where the application will reside
RUN mkdir /app

# Copy the application files (needed for production)
ADD conf /app/conf
ADD controller /app/controller
ADD model /app/model
ADD util /app/util

# Set the working directory to the app directory
WORKDIR /app

# Expose the application on port 9000.
# This should be the same as in the app.conf file
EXPOSE 9000

# Set the entry point of the container to the application executable
ENTRYPOINT /app/gotax
