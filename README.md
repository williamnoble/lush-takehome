# URL Shortener

Take Home Task for **Lush**. Took approx 5hr.

### Setup

1. Run docker-compose to access Postgres instance.
2. Apply `up.sql`
```shell
$ docker-compose up -d
```

### Run the API

```shell
$ go run cmd/api/main.go
```

### Test the API

I use Goland and run the tests from within `handlers_test.go`. Alternatively:

```shell
$ go test ./...

$ curl -i http://localhost:8000/YH-nYjDnR
```

### Run the CLI

```shell
$ go run cli/cmd/main.go generate http://www.bbc.co.uk
```

### Limitations & Improvements
- Move tests to TDD. Create Interface for db abstraction.
- Isolate shorten fn so I can use !reflect.DeepEqual rather than asserting individual fields
- Postgres may not be the best choice of db vs e.g. redis.
- Wrapping the end json made it harder to test.

### Attribution
The file `serve.go` file in addition to the ReadJSON helper function are inspired by 'Alex Edwards', granted they are edited/simplified. Otherwise, all work is my own.

<br/>
William Noble 13th October 2021.