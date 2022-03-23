# Planned API architecture

## Auth Endpoints

- GET /api 
    - POST /auth/login - _Login user_ - {"identity": "foo","password": "bar"}
    - POST /auth/create - _Create user_ - {"username": "foo","password": "bar"}

## Application Endpoints

- GET /api
    - GET /trk (Only for logged in users)
        - GET   /whoami - _Return username_

        - POST  /submit - _Save a new timestamp_
        - POST  /delete/:uuid - _Delete a timestamp by UUID_
        - POST  /delete/:id - _Delete a timestamp by UUID_

        - GET   /list/all - _List all your Timestamps_
        - GET   /list/date/:date - _List all Timestamps for specific day_
        - GET   /list/week/:week - _List all Timestamps for specific week_
        
        - GET   /summary/date/:date - _Get summary for a date_
        - GET   /summary/week/:week - _Get summary for a week_
## Test Endpoints
- GET /ping - _Returns a "Pong" message_

## Standardized API responses
```JSON
{
    "status": "<error|success>", 
    "message": "<generic error or success message>", 
    "data": <Object>
}
```



