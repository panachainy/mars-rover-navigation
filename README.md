# Mars rover navigation

[![codecov](https://codecov.io/gh/panachainy/mars-rover-navigation/graph/badge.svg?token=n9ftX5F3nA)](https://codecov.io/gh/panachainy/mars-rover-navigation)

## Project Description

- Project structure explanation.

  ```txt
  .
  ├── go.mod
  ├── go.sum
  ├── Makefile
  ├── README.md
  ├── src
  │   ├── console  // handles console input, output and format.
  │   │   ├── console.go
  │   │   ├── consoleImpl_test.go // I decide to places unit-test file follow golang strategy.
  │   │   ├── consoleImpl.go
  │   │   ├── wire_gen.go
  │   │   └── wire.go
  │   ├── main.go  // first place that go is run (in normally I place it at `/cmd/http/main.go`, `/cmd/consumer/main.go`)
  │   ├── model
  │   │   └── share_model.go // share model that use in this application.
  │   └── modules
  │       ├── environment // handle Grid, Boundary & Obstacles
  │       │   ├── environment_impl_test.go
  │       │   ├── environment_impl.go
  │       │   └── environment.go
  │       ├── game // main logic `NavigateRover` & control the game with rover, environment.
  │       │   ├── game_impl.go
  │       │   └── game.go
  │       └── rover // handle Rover movement, direction and commands
  │           ├── rover_impl_test.go
  │           ├── rover_impl.go
  │           └── rover.go
  └── tests
      └── integration // integration test should be places isolate from source code, I decide to place here.
          └── main_test.go
  ```

- programming language is Golang

## Installation Instructions

- install dependencies & clean by `make tidy`

## Usage Instructions

- For development use `make dev` (auto reload)
- For run use `make start`

## Testing Instructions

- `make t` for run all unit tests.
- `make it` for run integration tests.
- `make test.all` for run both of unit and integration tests.
- `make test.report` for generate test coverage report.

## Additional Notes
