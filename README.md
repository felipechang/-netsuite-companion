**NetSuite Companion**
=====================

A command-line tool for managing NetSuite development projects and scripts.

**Getting Started**
-------------------

To use this tool, simply run the command `nsc` in your terminal. You will be presented with a list of available
commands.

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

* `--inference=your instructions`: Run the generated file through OpenAI ChatGPT4.

**License**
---------

This software is licensed under the MIT License.