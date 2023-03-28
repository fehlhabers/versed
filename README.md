# versed
A tool for string replacement in automation.

## Rationale
In many situations when working with pipelines there is not always a good way to handle versioning. A common way to get around this is to introduce `sed` in a pipeline to replace certain parts.
However, if you would like the versions to be stored as Infrastructure as Code (IaC) and be able to update these through a pipeline, then it's more convenient to store the data structured.

A use case which triggered this project was the poor version handling in Terraform. Terrafile implementations all relied on pulling down entire repositories beforehand which is very clumsy if working with a monorepo or multiple modules in a repository.

## How it works

Versed is a lightweight approach to versioning, which basically replaces variables in text files in a given folder and puts the output in a different folder. The current version supports 2 parts as it's common with a name and version, although this is likely to become more generic.

## Installation

### Linux
Download the binary from the release pages and move it to a folder which is in your path. For example: `~/.local/bin` would work for Ubuntu.

### Others
Multiple OS builds will come, but until then, the project can be built locally. In order to build, Golang is needed.
Clone the repository and run `go build .` in order to build an executable.

## Usage

Versed requires a config file which defines what variable which should be replaced and by what concatenated value.
Defaults to use `versed.yml` in the same folder.
```
versed -f test/versed.yml
versed
```

A versed.yml could look like this:
```
target: test/data
output: test/output
sources:
  testsource1: 
    source: source1
    version: v1
  testsource2:
    source: source2
    version: v2
```

In order for versed to update values, the target files would need to be tagged like the following, using the `&(versed.SOURCE_NAME)` format. 
```
[
  "&(versed.testsource1)",
  "&(versed.testsource2)"
]
```
When running versed with the above config and target file, it would produce:
```
[
  "source1v1",
  "source2v2"
]
```