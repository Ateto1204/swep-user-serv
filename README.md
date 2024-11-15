# Swep user service

- url path: `/user`

#### Properties
- ID: email
    - user id aka email
- Name: string
    - user name
- Chats: `[]Chat.ID`
    - record the id of chat rooms
    - `Chat.ID`: string
- Friends: `[]User.ID`
    - record the user's friends through the id of user's friends
    - `User.ID`: string
- CreatAt: string
    - record the created time of the user

#### API
- `user-01` : SaveUser(userID, name: string)
    - url path: `/api/user-add`
    - method: `POST`
    - request body : 
        ```json
        {
            "id": "demo@gmail.com",
            "name": "demo-user"
        }
        ```
    - response body : 
        ```json
        {
            "id": "demo@gmail.com",
            "name": "demo-user",
            "chats": "[]",
            "friends": "[]",
            "create_at": "2024-11-10T04:59:11.461707232Z"
        }
        ```
- `user-02` : GetUser(userID: string)
    - url path: `/api/user-get`
    - method: `POST`
    - request body :
        ```json
        {
            "id": "demo@gmail.com"
        }
        ```
    - response body : 
        ```json
        {
            "id": "demo@gmail.com",
            "name": "demo-user",
            "chats": [],
            "friends": [],
            "create_at": "2024-11-03T04:32:36.886643Z"
        }
        ```
- `user-03` : AddFriend(userID, friendID: string)
    - url path: `/api/friend-add`
    - method: `PATCH`
    - request body : 
        ```json
        {
            "user_id": "demo@gmail.com",
            "friend_id": "friend@gmail.com"
        }
        ```
    - response body : 
        ```json
        {
            "id": "demo@gmail.com",
            "name": "demo",
            "chats": [],
            "friends": [
                "friend@gmail.com"
            ],
            "create_at": "2024-11-03T04:32:36.886643Z"
        }
        ```
- `user-04` : RemoveFriend(userID, friendID: string)
    - url path: `/api/friend-remove`
    - method: `PATCH`
    - request body :
        ```json
        {
            "user_id": "demo@gmail.com",
            "friend_id": "friend@gmail.com"
        }
        ```
    - response body : 
        ```json
        {
            "id": "demo@gmail.com",
            "name": "demo",
            "chats": [],
            "friends": [],
            "create_at": "2024-11-03T04:32:36.886643Z"
        }
        ```
- `user-05` : AddNewChat(userID, chatID)
    - url path: `/api/chat-add`
    - method: `PATCH`
    - request body : 
        ```json
        {
            "user_id": "demo@gmail.com",
            "chat_id": "demo-chat-id"
        }
        ```
- `user-06` : RemoveChat(userID, chatID)
    - url path: `/api/chat-remove`
    - method: `PATCH`
    - request body : 
        ```json
        {
            "user_id": "demo@gmail.com",
            "chat_id": "demo-chat-id"
        }
        ```
- `user-07` : DeleteUser(userID)
    - url path: `/api/user-del`
    - method: `DELETE`
    - request body : 
        ```json
        {
            "id": "demo@gmail.com"
        }
        ```