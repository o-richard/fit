# Fit

A platform that tracks user health entries & fitness data. User health entries support text, images and videos. Fitness data is obtained from any configured health & wellness app.

# TODO

[X] -> Setup a database connection. A database should be created in case it does not exist.

[ ] -> Parse exported data from **configured** health & wellness apps. Any parsed data should not be re-parsed.

[ ] -> Process & store user-input health entry.

[ ] -> Setup a HTTP server to access a GUI of the application. The server should allow for timeline view of the stored entries.

[ ] -> Setup Qdrant to embed the stored unstructured data.

[ ] -> Process user prompt in Qdrant to obtain entries matching the prompt.

[ ] -> Process user prompt in LLM to provide answers on health entries in the provided timerange. Additional context may be provided together with the prompt.

# Project Setup

## Prerequisites

### Usage

### Development

- [go](https://go.dev/doc/install). It should be installed in your system. The version installed should be compatible with version 1.22.2

## Instructions

### Usage

### Development

Run the application.

```
go run .
```

# Project structure

### /assets/images/

This directory stores user uploaded images.

### /assets/static/

This directory contains the assets used in the web application i.e Javascript, CSS, images etc

### /pkg/db/

This directory contains database specific operations.

### /pkg/parser/

This directory contains operations for parsing files from configured health & wellness apps and storing the input in the database.

Any new parser should implement `Parser` interface which obtains records to insert to the database. The parser is responsible for ensuring that entries are not re-parsed & the process is concurrent-safe.

##### Samsung Health (samsung.go)

After data has been exported from the samsung health application, the unzipped export should be places under `/data/parsers/samsung` directory. The unzipped directory contains csv files & a json directory. The files & contents parsed may change over time.

In the CSV files, the first row is assumed to contain 'useless' metadata. The timestamps are in the layout `2006-01-02 15:04:05.000`.

Currently, the CSV files whose names contain the following substrings are parsed. Only some of the data is parsed.

1. **.report.**
2. **.tracker.pedometer_day_summary.** - step_count, update_time, create_time, distance, calorie. all the fields are assumed to be provided and no defaults are assigned.

Samsung Health does not record the timezone information together with the datetime. Timezone is obtained from one of the files containing the substrings (**.badge.**, **.exercise.**, **.report.**, **.pedometer_step_count.**) in the file name. The header should be **_time_offset_** in **.badge.**, **_timezone_** in **.report.**, **_com.samsung.health.exercise.time_offset_** in **.exercise.**, **_com.samsung.health.step_count.time_offset_** in **.pedometer_step_count.**. At least one of the files should be present. **ASSUMPTION:** the first timezone found in any of the listed file corresponds to all the datetimes present in the files.

The Samsung Health version used is **_Version 6.26.6.001_**

### /pkg/server/

This directory contains operations specific to the web application server.
