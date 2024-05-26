# Fit

A platform that tracks user health entries & fitness data. User health entries support text, images and videos. Fitness data is obtained from any configured health & wellness app.

# TODO

[X] -> Setup a database connection. A database should be created in case it does not exist.

[ ] -> Parse exported data from **configured** health & wellness apps. Any parsed data should not be re-parsed.

[ ] -> Process & store user-input health entry.

[ ] -> Setup a HTTP server to access a GUI of the application. The server should allow for timeline view of the stored entries.

[ ] -> Setup Qdrant to embed the stored unstructured data.

[ ] -> Process user prompt in Qdrant to obtain entries matching the prompt.

[ ] -> Process user prompt in LLM to provide answers on health entries in the provided timerange.

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

1. **Samsung Health** - After data has been exported from the samsung application, the unzipped export should be places under `/data/parsers/samsung` directory. The unzipped directory contains csv files & a json directory. The files & contents parsed may change over time. Currently, the CSV files whose names contain the substrings **_activity_day_summary_**, **_calories_burned_details_**, **_report_**, **_step_daily_trend_**, **_pedometer_day_summary_** are parsed. Only some of the data is parsed.

### /pkg/server/

This directory contains operations specific to the web application server.
