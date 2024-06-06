# Fit

A platform that tracks user health entries & fitness data. User health entries support text and images. Fitness data is obtained from any configured health & wellness app.

The functionality of the platform happens on device unless otherwise.

# TODO

[X] -> Setup a database connection. A database should be created in case it does not exist.

[X] -> Parse exported data from **configured** health & wellness apps. Any parsed data should not be re-parsed.

[X] -> Process & store user-input health entry.

[ ] -> Setup a HTTP server to access a GUI of the application. The server should allow for timeline view of the stored entries.

[ ] -> Setup Qdrant to embed the stored unstructured data.

[ ] -> Process user prompt in Qdrant to obtain entries matching the prompt. History of prompts in the current execution window may be shown.

[ ] -> Process user prompt in LLM to provide answers on health entries in the provided timerange. Additional context may be provided together with the prompt.

# Project Setup

## Prerequisites

### Usage

- The application binary of the target platform.
- The `data/parsers/[parser type]` directory in the same directory as the application binary.
- The `assets/media` & `assets/static` directories in the same directory as the application binary.

### Development

- [go](https://go.dev/doc/install). It should be installed in your system. The version installed should be compatible with version 1.22.2
- [tailwind cli](https://github.com/tailwindlabs/tailwindcss/releases/). It should be installed in the tools directory as an executable named "tailwindcss". It outputs CSS stylesheets based on used tailwind classes. The version to download is 3.4.3.
- [templ](https://templ.guide/quick-start/installation). It should be installed in your system. It generates template components. The version installed should be compatible with version v0.2.707.
- [air](https://github.com/cosmtrek/air). It should be installed in your system. It performs live reloading for the Go application. The version installed should be compatible with version 1.52.1.

## Instructions

### Usage

- Append `--help` to the executable build path for usage instructions.

```
Usage of [executable build path] [command] [options]

Available commands:
parse
  -type string
        Name of the parser. Choices are samsung.
runserver
  -port string
        Start the HTTP server at the provided port. (default "8000")
```

### Development

Generates CSS styles based on the tailwind setup

```bash
make styles
```

Generate templ components.
Modify the SHELL variable in the Makefile if need be i.e templ should be discovered within the system path.

```bash
make generate
```

Run the application.

```
go run ./... [command] [options]
```

Live reload of the application (Only the HTTP Server).
Modify the SHELL variable in the Makefile if need be i.e templ should be discovered within the system path.

```bash
air
```

For other Go commands, refer to the offical Golang documentation.

For other templ commands, refer to the [offical guide](https://templ.guide/).

For other air commands, refer to the [offical guide](https://github.com/cosmtrek/air).

# Project structure

### /assets/dev/

This directory development-specific configurations eg Tailwind setup.

### /assets/media/

This directory stores user uploaded images.

### /assets/static/

This directory contains the assets used in the web application i.e Javascript, CSS, images etc

### /pkg/db/

This directory contains database specific operations.

### /pkg/parser/

This directory contains operations for parsing files from configured health & wellness apps and storing the input in the database.

Any new parser should implement `Parser` interface which obtains records to insert to the database. The parser is responsible for ensuring that entries are not re-parsed & the process is concurrent-safe. The parser should be registered in `ParseFitnessAppRecords` function.

##### Samsung Health (samsung.go)

After data has been exported from the samsung health application, the unzipped export should be places under `/data/parsers/samsung` directory. The unzipped directory contains csv files & a json directory. The files & contents parsed may change over time.

Regarding the json directory, the subdirectories in each directory have a single character name that corresponds to first letter of its files.

In the CSV files, the first row is assumed to contain 'useless' metadata. The timestamps are in the layout `2006-01-02 15:04:05.000`.

Currently, the CSV files whose names contain the following substrings are parsed. Only some of the data is parsed.

1. **.tracker.pedometer_day_summary.** - step_count, update_time, create_time, distance (in metres), calorie (in kcal). all the fields are assumed to be provided and no defaults are assigned.

**.report.** is may or may to be parsed in the future. currently it is not parsed since there is not a clear distinction on whether the values are 'correct' or not i.e the integers `2147483647` & `3.4028235e38` are used where sensible values are expected. furthermore, each rows in the CSV file correspond to period (at least a week long) & the retrieval of the year depends on certain hacks since it is not provided with the period.

Samsung Health does not record the timezone information together with the datetime. Timezone is obtained from one of the files containing the substrings (**.badge.**, **.exercise.**, **.report.**, **.pedometer_step_count.**) in the file name. The header should be **_time_offset_** in **.badge.**, **_timezone_** in **.report.**, **_com.samsung.health.exercise.time_offset_** in **.exercise.**, **_com.samsung.health.step_count.time_offset_** in **.pedometer_step_count.**. At least one of the files should be present. **ASSUMPTION:** the first timezone found in any of the listed file corresponds to all the datetimes present in the files.

The Samsung Health version used is **_Version 6.26.6.001_**

### /pkg/server/

This directory contains operations specific to the web application server.

### /testdata

This directory contains sample data used in the project.

### /testdata/parsers/samsung

This directory contains some of the files that are part of the unzipped export from Samsung Health. Unnecessary columns have been dropped from the CSV files.

### /tools

This directory contains external tools used in the project.

### db.sqlite3

This sqlite database contains the test data used while building the application, it should be deleted when one plans to use the application with their own data.

# UI/UX

Currently, the GUI is focussed on medium & large screen sizes.

# Credits

- [Icons Pack](https://www.iconpacks.net/)
