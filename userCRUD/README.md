# gRPC CRUD

Basic gRPC server written in go for CRUD operation

## Installation

Project requires having Go 1.16 or above and postgreSQL. Once you clone(or go get) you need to configure the following:

### For Server

1. Change the [DotEnv](https://github.com/KushagraMehta/gRPC-Tutorial/blob/main/userCRUD/.env) file according to your envirnment.
2. Install Task runner `go get -u github.com/go-task/task/v3/cmd/task`
3. Install dependence `go mod download`
4. `task migration` for initializing database.
5. `task server` for starting server(at DEFAULT PORT 50051).

### For Client

1. Install grpcUi dependence `go get github.com/fullstorydev/grpcui/...`
2. grpcUi dependence `go install github.com/fullstorydev/grpcui/cmd/grpcui`
3. `task client` for starting client(at DEEFAULT PORT 8090)

> You can also use [Client.go](https://github.com/KushagraMehta/gRPC-Tutorial/blob/main/userCRUD/cmd/client/client.go) to connect with server

### Or

### By Docker

```shell
 docker-compose up --build
```

Will start server(at port 50051) and client(at port 80)

> Sometime client does not start because server wait for database container to start so it decline connection with client, so try to restart the docker

---

## API end-points

[proto file](https://github.com/KushagraMehta/gRPC-Tutorial/blob/main/userCRUD/pkg/protobuf/user/user.proto)

1. `RegisterUser(*RegisterUserRequest) (*UserID, error)` (Unary End-Point)

   Will take [RegisterUserRequest](https://github.com/KushagraMehta/gRPC-Tutorial/blob/main/userCRUD/pkg/protobuf/user/user.proto#L6-L12) which shoud have proper user data and send [userid](https://github.com/KushagraMehta/gRPC-Tutorial/blob/main/userCRUD/pkg/protobuf/user/user.proto#L13-L15) regester in database.

2. `GetUserDetail(*UserID) (*UserDetailResponse, error)` (Unary End-Point)

   Will take [UserId](https://github.com/KushagraMehta/gRPC-Tutorial/blob/main/userCRUD/pkg/protobuf/user/user.proto#L13-L15) and Send [UserDetailResponse](https://github.com/KushagraMehta/gRPC-Tutorial/blob/main/userCRUD/pkg/protobuf/user/user.proto#L17-L24) consist of User Data if user is present or it will raise [error](https://github.com/KushagraMehta/gRPC-Tutorial/blob/b41251a0719f2086364a7e1f2fd6251a8c4b79c6/userCRUD/pkg/model/user.go#L54-L59) if user does not exist.

3. `GetUsersList(*UsersListRequest) (*UsersListResponse, error)` (Unary End-Point)

   Will take an array of userids as [UsersListRequest](https://github.com/KushagraMehta/gRPC-Tutorial/blob/b41251a0719f2086364a7e1f2fd6251a8c4b79c6/userCRUD/pkg/protobuf/user/user.proto#L26-L28) And return an Array of [UsersListResponse](https://github.com/KushagraMehta/gRPC-Tutorial/blob/b41251a0719f2086364a7e1f2fd6251a8c4b79c6/userCRUD/pkg/protobuf/user/user.proto#L17-L24).

   > As Go append array in power of 2 so this End-point can waste resource for better efficiency use StreamUsersList

4. `StreamUsersList(UserService_StreamUsersListServer) error` (Bi-Directional Streaming End-Point)

   Will take stream of [userIDs](https://github.com/KushagraMehta/gRPC-Tutorial/blob/b41251a0719f2086364a7e1f2fd6251a8c4b79c6/userCRUD/pkg/protobuf/user/user.proto#L13-L15) and send Stream of [user details](https://github.com/KushagraMehta/gRPC-Tutorial/blob/b41251a0719f2086364a7e1f2fd6251a8c4b79c6/userCRUD/pkg/protobuf/user/user.proto#L17-L24).
