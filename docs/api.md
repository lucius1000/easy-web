# API Documentation

## Health Check
- `GET /health`
- Response: `{"code":0,"message":"success","data":"OK"}`

## User APIs
### 1. Create User
- **Method**: POST
- **URL**: `/api/v1/users`
- **Body**: 
```json
{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
}
```

### 2. List Users
- **Method**: GET
- **URL**: `/api/v1/users?page=1&page_size=10`

### 3. Get User by ID
- **Method**: GET
- **URL**: `/api/v1/users/:id`

### 4. Update User
- **Method**: PUT
- **URL**: `/api/v1/users/:id`
- **Body**:
```json
{
    "name": "John Doe Updated",
    "email": "john.updated@example.com"
}
```

### 5. Delete User
- **Method**: DELETE
- **URL**: `/api/v1/users/:id`
