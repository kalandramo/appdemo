include "./base.thrift"
namespace go user

typedef i64 int64

struct User {
    1: int64 user_id
    2: string username
    3: string avatar
}

struct CreateUserRequest {
    1: string username (vt.min_size = "5")
    2: string password (vt.min_size = "8")
}

struct CreateUserResponse {
    255: base.BaseResp base_resp
}



service UserService {
    CreateUserResponse CreateUser(1: CreateUserRequest req)
}