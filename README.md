# Simple jwt auth service
## Commands
1. `docker compose up` - run project
## Endpoints
1. `POST: /api/auth/get-tokens` - issue pair refresh(via cookie) and access(via response body) tokens 
   1. Input params 
   ``` 
   {
     "guid": string
   }
   ```
   2. Response 
   ```
    "accessToken": string
   ```
   Cookie: `refreshToken=string`

2. `POST: /api/auth/refresh` - refresh tokens pair 
    1. Input params
   ``` 
   {
    "accessToken": string 
    "refreshToken": string 
   }
   ```
    2. Response
   ```
    "accessToken": string
   ```
   Cookie: `refreshToken=string`
