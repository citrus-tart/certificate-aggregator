# Certificate aggregator
This microservice will pick up all events on start, create the current state of certificates in memory and expose them via an API.

## TODO
* Store state in a persistent store (DB or Redis)
* Add tests