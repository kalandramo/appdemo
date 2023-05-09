namespace go base

typedef i64 int64

struct BaseResp {
    1: int64 status_code
    2: string status_message
    3: int64 service_time
}


