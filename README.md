# Go API Clean architecture
This repository contains an implementation of a Go application following Clean Architecture principles. The application is a simple API that demonstrates the use of Clean Architecture in Go.

## Directory structure
```
├── cmd
│   └── api
│       └── main.go
└── internal
    ├── configs
    │   └── configs.go
    ├── domain                # Application core layer
    │   ├── entities
    │   │   ├── base.go
    │   │   └── product.go
    │   ├── models
    │   │   └── product.go
    │   └── services
    │       └── product.go
    ├── infrastructure        # Infrastructure layer
    │   ├── db
    │   │   └── db.go
    │   └── repository
    │       └── product.go
    └── presentation          # Presentation layer
        ├── handlers
        │   ├── handlers.go
        │   ├── product
        │   │   ├── handlers.go
        │   │   └── routes.go
        │   └── user
        │       ├── handlers.go
        │       └── routes.go
        └── http
            └── http.go
```

I tend to follow the [Go standard project layout](https://github.com/golang-standards/project-layout), and try to implement the core idea of Clean architecture. The goal is to make the application's core business logic independent of any particular framework, database, or delivery mechanism.

![Clean Architecture: horizontal layer view](https://learn.microsoft.com/en-us/dotnet/architecture/modern-web-apps-azure/media/image5-8.png)
[Image source](https://learn.microsoft.com/en-us/dotnet/architecture/modern-web-apps-azure/common-web-application-architectures#clean-architecture)

The application's concerns are separated into three main layers: the Domain layer, the Infrastructure layer, and the Presentation layer.
### Domain
The Domain layer contains the core business logic of the application. It defines the data structures and operations that represent the problem domain, as well as interfaces and services that define the contract between the domain and the other layers. This layer should be independent of any particular framework or delivery mechanism.

### Infrastructure
The Infrastructure layer contains the code that implements the application's technical requirements, such as persistence, logging, and communication with external systems. In this repository, the Infrastructure layer mainly contains concrete implementations of interfaces defined in the Domain layer and also third-party libraries used to implement the technical requirements of the application.

### Presentation
The Presentation layer is responsible for handling user input and displaying the data to the user. In this repository, The key responsibilities of the Presentation layer are to handle user input, call the appropriate use cases in the Domain layer, and translate the results into a format that can be displayed to the user.

By separating the application's concerns into these three layers, and making sure that each layer depends only on the layers below it, the application becomes more testable and maintainable. The dependencies between the layers should be based on abstractions and contracts (interfaces) rather than on concrete implementations, this is what makes it easy to test and change components without affecting the whole application.


## References
[The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

[Common web application architectures](https://learn.microsoft.com/en-us/dotnet/architecture/modern-web-apps-azure/common-web-application-architectures#clean-architecture)

[Goapp](https://github.com/bnkamalesh/goapp)

## Todo
- [ ] General response format
- [ ] Logging
- [ ] Error handling
- [ ] Adding Authentication
