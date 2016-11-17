# Parrot
WIP - Project localization system built with Go (backend) and Angular 2 (frontend).

TODO:

Backend:
- Add cache (redis?)
- Handle HEAD locale revisions (how would this work with project keys update while using old HEAD?)
- Pass project user role on GET /projects/:id/users
- Add support for client access token and role
- Separate auth issuing service from authenticator, conform to oauth2?
- Refactor main to cli and make configurable, migrate command, serve command etc...
- Make store non-destructive. Add snapshots?
- Add tests once features have been settled.

Frontend:
- Add styles.
- Add tests.
- Introduce typings.
