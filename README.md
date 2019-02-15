# travis-play
[![Build Status](https://travis-ci.com/dvictor/travis-play.svg?token=qZUpiMaAyepwqousKgbJ&branch=master)](https://travis-ci.com/dvictor/travis-play)


Simple Travis CI example

In docker-compose.yml we create two containers: a PostgreSQL DB and the application.
 
The go application waits for the DB to start and makes a simple query.

See .travis.yml

To run the test locally: `integration-test.sh`
