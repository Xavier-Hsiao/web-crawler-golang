# web-crawler-golang
A simple web crawler built in Golang.

## Prerequisites
Go 1.20 or higher installed on your machine.

## Run the program

### Clone the repository

```bash
git clone https://github.com/Xavier-Hsiao/web-crawler-golang.git
cd web-crawler-golang
```

### Build the CLI
```bash
go build web-crawler-golang
```

### Run the CLI
```bash
go run . [baseURL] [maxConcurrency] [maxPages]
```

## Ideas for extending the project
- Make a script run on a timer and deployer it to the server. Have it email you every so often with a report
- Save the report as a CSV spreadsheet rather than printing it to the console
- Use a graphics library to create an image that shows the links between the pages as a graph visualization
- Count external links, as well as internal links, and add them to the report
