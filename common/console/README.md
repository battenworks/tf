# console
Shared console functions that can be leveraged by the other tools in this repo

## Usage
Create a `go.mod` file in the new tool's root directory. Contents of the file as follows:
```
module <tool name>

require console v0.0.0

replace console => ../../common/console

```
where `../../common/console` is the relative path from your tool's `go.mod` file to this directory
