FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64\
    GM65_PORT_NAME=/dev/serial/by-id/usb-USBKey_Chip_USBKey_Module_202730041341-if00
#RUN apk update && apk upgrade && \
#    apk add --no-cache bash git openssh
# Move to working directory /build
WORKDIR /gm65

# Copy and download dependency using go mod
#COPY go.mod .
#COPY go.sum .
#RUN go mod download

# Copy the build image into the container
# need to build like this:
# $ CGO_ENABLED=0 go build -o gm65server
ADD gm65server .

# Build the application
#RUN go build -o server .

# Export necessary port
EXPOSE 8070

# Command to run when starting the container
CMD ["/gm65/gm65server"]