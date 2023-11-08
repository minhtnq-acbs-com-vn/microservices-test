# microservices-test

# TESTING

- go test ./... -coverprofile=coverage.out

# NOTE

- Price service is on port 10000
- Book service is on port 11000
- Helper service is on port 12000
- This application doesn't have multiple layers like controller, service, repository, etc. because it's easier to review
  this way, I think.

# LIMITATION

- Don't have a fallback, circuit breaker, retry, etc. because I don't know how to implement it
- Can't exclude proto files from coverage
- Booking Service and Helper Service don't have high test coverage because I don't know how to replicate bad cases like
  database error