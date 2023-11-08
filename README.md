# microservices-test

# TESTING

- go test ./... -coverprofile=coverage.out

# NOTE

- This application doesn't have multiple layers like controller, service, repository, etc. because it's easier to review
  this way.

# LIMITATION

- Can't properly return error message
- Can't exclude proto files from coverage