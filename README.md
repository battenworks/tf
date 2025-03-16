# tf

Wrapper for the Terraform CLI (as of v2.0.0-beta5, it targets the OpenTofu CLI).
Provides some opinionated commands to help with Terraform CLI use.
All other commands are passed directly to the Terraform CLI.
It only works when run in a directory that contains a **backend.tf**.
If no backend.tf is found, the tool exits.

# NOTES

As of v2.0.0-beta5, this tool targets OpenTofu because that's where my focus lies.
I'll update the documentation when v2 is ready.
A lot has changed.

## Use tf clean

To clean and initialize your Terraform working directory.

```
tf clean
```

- Removes the **.terraform/** directory and **.terraform.lock.hcl** file
- Runs **terraform init**

## Use tf qplan (removed in v2 beta)

To run a Terraform plan in quiet mode, which removes drift output for clarity.

```
tf qplan
```

- Runs **terraform plan** and filters the output

Your output will resemble the following.

```
Note: Objects have changed outside of Terraform

---- 12 lines hidden ----

No changes. Your infrastructure matches the configuration.

Your configuration already matches the changes detected above. If you'd like to update the Terraform state to match, create and apply a refresh-only plan.
```

## Use tf off

For rapid development of Terraform config.
This command modifies config files but **backend.tf** and **providers.tf**, so the next **apply** will tear down the resources.
Used in conjunction with **tf on**.

```
tf off
```

- Adds the **.off** extension to config files in the current directory (with the exception of **backend.tf** and **providers.tf**).

## Use tf on

For rapid development of Terraform config.
This command modifies config files so the next **apply** will stand up the resources.
Used in conjunction with **tf off**.

```
tf on
```

- Removes the **.off** extension from config files in the current directory.

## Use tf replan (added in v2)

Because sometimes I don't want to do a clean, then a plan.

```
tf replan
```

- Runs `init -upgrade`, then `plan`.

## Use tf test (added in v2)

Out of the box, `test` requires some extra steps I don't want to do.

```
tf test
```

- Runs `fmt -recursive`, then `init -upgrade`, then `test`.
