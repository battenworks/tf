# tf
This tool wraps Terraform to provide additional functionality.

## Build tf
From the `app/tf` directory, execute the following.

```
go build -o <some place on your path>/tf
```

## Use tf clean
To clean and prepare your Terraform working directory for further commands.
```
tf clean
```
- Removes the `.terraform/` directory and `.terraform.lock.hcl` file
- Initializes Terraform (if there is no `backend.tf` in the working directory, the program will exit)
- Selects the `default` workspace

You can supply your own workspace with an optional flag.
```
tf clean -workspace=my-workspace
```

## Use tf plan
To run a Terraform plan, optionally altering the output for clarity.
```
tf plan
```
- Runs `terraform plan` and outputs the results

You can supply the `-hide-drift` flag to suppress Terraform's verbose refresh-step output.
```
tf plan -hide-drift
```
Your output will resemble the following.
```
Note: Objects have changed outside of Terraform

---- 12 lines hidden ----

No changes. Your infrastructure matches the configuration.

Your configuration already matches the changes detected above. If you'd like to update the Terraform state to match, create and apply a refresh-only plan.
```

## Use tf off
For rapid development of Terraform config. This command modifies all config files but `backend.tf`, so the next `apply` will tear down the resources. Used in conjunction with `tf on`.
```
tf off
```
- Adds the `.off` extension to all config files in the current directory (with the exception of `backend.tf`).

## Use tf on
For rapid development of Terraform config. This command modifies all config files so the next `apply` will stand up the resources. Used in conjunction with `tf off`.
```
tf on
```
- Removes the `.off` extension from all config files in the current directory.