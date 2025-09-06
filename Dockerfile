FROM golang:1.24 AS build-stage

WORKDIR /app

# First, copy mod and sum files to directory,
COPY go.mod go.sum ./

# Afterwards, download the dependencies from the go.mod file
RUN go mod download

# Copy the main.go file which starts the project, and the api folder
COPY . .

RUN go build -o digital-scoreboard .



# This is the run stage now, pulling from gcr
FROM gcr.io/distroless/base-debian12 AS runtime

# Create a new directory for the run time image
WORKDIR /app

# Copy the binary create during the build stage into the run time image
COPY --from=build-stage /app/digital-scoreboard .

# expose port 8080 to run the image on
EXPOSE 8080

# Finally, runthe go binary!
CMD [ "./digital-scoreboard" ]