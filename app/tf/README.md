# tf
This tool wraps Terraform to provide additional functionality

## Build tf
From the `app/tf` directory, execute the following:

```go build -o <some place on your path>/tf```

## Use tf clean
From the desired Terraform scope directory, execute the following:

```tf clean```

* Removes the `.terraform/` directory and `.terraform.lock.hcl` file
* Initializes Terraform (if there is no `backend.tf` in the working directory, the program will exit)
* Selects the `default` workspace

You can supply your own workspace with an optional flag:

```tf clean -workspace=my-workspace```
