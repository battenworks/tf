# dc
This tool wraps docker-compose to provide additional functionality

## Build dc
From the `app/dc` directory, execute the following:

```go build -o <some place on your path>/dc```

## Use dc up
Execute the following:

```dc up```

* Runs the command ```docker-compose up -d```
* Runs the command ```docker-compose stop healer```
* Runs the command ```docker-compose stop listener```
* Runs the command ```docker-compose start healer```
* Runs the command ```docker-compose start listener```

## Use dc fix
Execute the following:

```dc fix```

* Runs the command ```docker-compose stop healer```
* Runs the command ```docker-compose stop listener```
* Runs the command ```docker-compose start healer```
* Runs the command ```docker-compose start listener```

## Use dc down
Execute the following:

```dc down```

* Runs the command ```docker-compose down```
