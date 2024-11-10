**NetSuite Companion**
=====================

A command-line tool for managing NetSuite development projects and scripts.

**Getting Started**
-------------------

To use this tool, you'll need to build the binary first.

Run the following command in your terminal:

```
go build -o ~/go/bin/nsc main.go
```

to produce the `nsc` binary in go binaries directory. Once built, you can run the command `nsc` in your terminal. You
will be presented with a list of available commands.

**Commands**
------------

### Initialization

* `init`: Initializes the global and working directory settings. Use the `--force` flag to re-initialize the settings.
  You can also set the `OPENAI_API_KEY` environment variable to enable the inference service.

### File Management

* `add`: Creates new files for your NetSuite development projects.
  + `project`: Creates a new project file.
  + `bundle`: Creates a new bundle script file.
  + `client`: Creates a new client script file.
  + `formclient`: Creates a new form client script file.
  + `mapreduce`: Creates a new map reduce script file.
  + `massupdate`: Creates a new mass update script file.
  + `portlet`: Creates a new portlet script file.
  + `restlet`: Creates a new restlet script file.
  + `scheduled`: Creates a new scheduled script file.
  + `suitelet`: Creates a new suitelet script file.
  + `userevent`: Creates a new user event script file.
  + `workflowaction`: Creates a new workflow action script file.
    + `module`: Creates a new module file.
  + `type`: Creates a new TypeScript type file.

### Additional Features

* `--inference=your instructions`: Run the generated file through OpenAI ChatGPT.

**License**
---------

This software is licensed under the MIT License.

**Contributing**
--------------

We welcome contributions to the NetSuite Companion! Please see our contributing guide for more information on how to get
involved.

**Issues**
---------

If you encounter any issues while using the NetSuite Companion, please report them to our issue tracker. We'll do our
best to help you resolve the issue as quickly as possible.

I added a step to build the binary using `go build` before running the `nsc` command.