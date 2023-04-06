# versed
A tool for string replacement in automation.

## Rationale
In many situations when working with pipelines there is not always a good way to handle versioning. A common way to get around this is to introduce `sed` in a pipeline to replace certain parts.
However, if you would like the versions to be stored as Infrastructure as Code (IaC) and be able to update these through a pipeline, then it's more convenient to store the data structured.

A use case which triggered this project was the poor version handling in Terraform. Terrafile implementations all relied on pulling down entire repositories beforehand which is very clumsy if working with a monorepo or multiple modules in a repository.

## How it works

Versed is a lightweight approach to versioning, which basically replaces variables in text files in a given folder and puts the output in a different folder. The current version supports 2 parts as it's common with a name and version, although this is likely to become more generic.

Versed uses a `versed.yml` file to configure the variables to replace.

A versed.yml could look like this:
```
target: data
output: output
sources:
  testsource1: 
    source: source1
    version: v1
  testsource2:
    source: source2
    version: v2
```

In order for versed to update values, the target files would need to be tagged like the following, using the `&(versed.SOURCE_NAME)` format. 

So, if a target file looks like this:
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

## Action

### Usage

The recommended approach is that the `versed.yml` is put in the root of where you want to apply your versions. If you are in a GitOps environment with multiple stages and regions, you might have something like this: 
```
    steps:
    - name: Set versions
      uses: fehlhabers/versed@v2
      with: 
        dir: 'environment/westeurope/prod'

    - name: Apply terraform
    ...
    ...
```


## Stand-alone

Apart from being used in a pipeline, versed is packaged as a binary which can be used stand alone in local development as well. 

### Installation

#### Linux
Download the binary from the release pages and move it to a folder which is in your path. For example: `~/.local/bin` would work for Ubuntu.

#### Others
Multiple OS builds will come, but until then, the project can be built locally. In order to build, Golang is needed.
Clone the repository and run `go build .` in order to build an executable.

### Usage

Versed is simply run in the same folder as where you have your `versed.yml` or you specify a different config file.
```
versed -f test/versed.yml
versed
```
