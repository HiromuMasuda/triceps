# triceps
Golang Clean Architecture Example


## Clean Architecture

![clean_architecture](https://camo.qiitausercontent.com/c09feaf712816dcddd01310055d24efe3e548b77/68747470733a2f2f71696974612d696d6167652d73746f72652e73332e616d617a6f6e6177732e636f6d2f302f34343134322f61373634336335332d386363302d623037392d303734352d6132306630366632333337322e6a706567)


## Directory Tree

```
└── src
    ├── app
    │   ├── domain
    │   │   └── user.go
    │   ├── infrastructure
    │   │   ├── router.go
    │   │   └── sqlhandler.go
    │   ├── interfaces
    │   │   ├── controllers
    │   │   │   ├── context.go
    │   │   │   ├── error.go
    │   │   │   └── user_controller.go
    │   │   └── database
    │   │       ├── sqlhandler.go
    │   │       └── user_repository.go
    │   ├── server.go
    │   └── usecase
    │       ├── user_interactor.go
    │       └── user_repository.go
```

## References

https://qiita.com/hirotakan/items/698c1f5773a3cca6193e
