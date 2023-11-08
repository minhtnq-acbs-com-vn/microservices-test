# microservices-test

# HOW TO RUN

- docker compose build 
- docker compose up
- (while inside root project directory)

# TESTING

- go test ./... -coverprofile=coverage.out

# NOTE

- Price service is on port 10000
- Book service is on port 11000
- Helper service is on port 12000
- All proto files are in the proto folder which is located in the root project directory
- This application doesn't have multiple layers like controller, service, repository, etc. because it's easier to review
  this way, I think.

# LIMITATION

- Don't have a fallback, circuit breaker, retry, etc. because I don't know how to implement it yet
- Can't exclude proto files from test coverage, for more accuracy on test coverage; I recommend running individual test
  coverage for each service
- Booking Service and Helper Service don't have high test coverage because I don't know how to replicate bad cases like
  database, buffer unmarshal + marshal error
- Bad system design decision - Should have implement dependency injection from the start to reuse db connection and
  better management
- Could have used more logging to trace and config file rather than hardcode
- Can only test between services by manually running the application