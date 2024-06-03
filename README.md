# API Handy_sports
## Description
Api made for the handy_sports project, which is a project that aims to help the manager and the user to manager their teams and their sports events. The API is made with Go and the database is made with SurrealDB.

## Installation
Check if you have Go installed in your machine, if not, you can download it [here](https://golang.org/dl/).
Or with homebrew in MacOS:
```bash
brew install go
```

After that, you can clone the repository and run the following command in the root of the project:
```bash
go run main.go
```

## Usage
The API has the following endpoints:
- GET /{collection}
- GET /{collection}/{id}
- POST /{collection}
- PUT /{collection}/{id}
- DELETE /{collection}/{id}

The collections are:
- users
- players
- teams
- sports
- events
- event-types
- positions
- match-history
- bills
