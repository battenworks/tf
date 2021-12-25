# command
Shared command functions that can be leveraged by the other tools in this repo

## Usage
Create a `go.mod` file in the new tool's root directory. Contents of the file as follows:
```
module <tool name>

require command v0.0.0

replace command => ../../common/command

```
where `../../common/command` is the relative path from your tool's `go.mod` file to this directory
