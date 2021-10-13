# URL Shortener

Take Home Task for **Lush**.
This took me approximately 5 hours to complete. There are several things I could improve but I thought this was a good limit for how long I should reasonably take for a take home.

### Setup

```shell
$ docker-compose up -d

### Apply the migration to the database
```

### Run the API

```shell
$ go run cmd/api/main.go
```

### Test the API

I use Goland and run the tests from within `handlers_test.go`. Alternatively:

```shell
$ go test ./...
```

### Run the CLI

```shell
$ go run cli/cmd/main.go generate http://www.bbc.co.uk
```

### Limitations & Improvements

- I didn't want to spend more time writing tests. Because of how I designed the handler it was
  difficult to write smaller unit tests, with more time I would've abstracted the database interactions to an Interface
  and broken the Shorten function to a separate function. A limitation of the testing is that I cannot predict the
  result of the shorten function, whereas if I split out the function I would have more granular control and could use
  !reflect.DeepEqual to compare the full structs as opposed to certain fields.
- The tests can be easily extended to include TDD []struct{}{} as we have more than one row in initial migration.
- I haven't done much with the CLI as I was running out of time. It's the simplest implementation possible.
- I wrote the env variable for ease in the program but left the facility to retrieve from env, also included .env as example as opposed to adding to .gitignore.
- Postgres is probably not the best choice of database, Redis is better, however I have more familiarity with Postgres
- Difficulties; working out how to test my handler properly, because I wrapped the Link struct it made it harder to test and had to unwrap to a separate struct. Before now, I didn't know how to test the Location of a redirect or indeed how to redirect to begin with. I spent over half my (self-allocated) time trying to work out how to test certain things. 

### Attribution
I used no outside reference for working out how to complete this task. However, the code within `serve.go` is at least in part attributed to AlexEdwards, as are my WriteJSON and ReadJSON functions albeit I've modified (simplified) them all heavily.

<br/>
<br/>
William Noble 13th October 2021.