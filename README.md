# microservices-test

# TESTING

- go test ./... -coverprofile=coverage.out

# NOTE

- Price service is on port 10000
- Book service is on port 11000
- This application doesn't have multiple layers like controller, service, repository, etc. because it's easier to review
  this way.

# LIMITATION

- Can't properly return error message
- Can't exclude proto files from coverage
- Booking Service don't have high test coverage because I don't know how to replicate bad cases like database error