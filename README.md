# tf
Wrapper for the Terraform CLI.
Provides some opinionated commands to help with Terraform CLI use.
All other commands are passed directly to the Terraform CLI.
It only works when run in a directory that contains a **backend.tf**.
If no backend.tf is found, the tool exits.

## Use tf clean
To clean and initialize your Terraform working directory.
```
tf clean
```
- Removes the **.terraform/** directory and **.terraform.lock.hcl** file
- Runs **terraform init**

## Use tf qplan
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
This command modifies all config files but **backend.tf**, so the next **apply** will tear down the resources.
Used in conjunction with **tf on**.
```
tf off
```
- Adds the **.off** extension to all config files in the current directory (with the exception of **backend.tf**).

## Use tf on
For rapid development of Terraform config.
This command modifies all config files so the next **apply** will stand up the resources.
Used in conjunction with **tf off**.
```
tf on
```
- Removes the **.off** extension from all config files in the current directory.
