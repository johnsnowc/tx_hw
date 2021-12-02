namespace go api

struct AddRequest {
	1: i64 id
	2: i64 room_id
    3: i64 item_id
    4: double payment
    5: string create_time
    6: i64 user_id
}

struct AddResponse {
    1: string message
	2: i64 id
}

struct DeleteByIDRequest {
    1: i64 id
}

struct DeleteByIDResponse {
    1: string message
    2: i64 id
}

struct PaymentsByRoomIDRequest {
    1: i64 room_id
}

struct PaymentsByRoomIDResponse {
    1: double sum
}

service Transaction {
    AddResponse add(1: AddRequest req)
    DeleteByIDResponse deleteByID(1: DeleteByIDRequest req)
    PaymentsByRoomIDResponse paymentsByRoomID(1: PaymentsByRoomIDRequest req)
}