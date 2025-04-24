# DogsAPI

The idea is to build an API for a web app that handles dog accommodation/lodging reservations, as that is my family current area of work.

#### Goal (for now)
The final goal is to have an API with:

- Auth
- relationship logic between users and dogs entities
- CI
- CD
- observability
    - Logging
    - Metrics
    - Dashboards
- Caching
- API behind reverse proxy load balancer
- High test coverage (focusing on integration tests)

### Stack and Tools (for now)
-  go
-  Gin framework (for now, i might change)
-  PostgreSQL
-  [migrate](https://github.com/golang-migrate/migrate) for DB migrations
- Github Actions for CI and CD

