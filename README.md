## API-DOCUMENTATION-SERVICE 
#### Version: 1
- ```API-BASE-URL: http://0.0.0.0:8093/api/v1```

-------

### List of available endpoints:

#### docs
- `POST /docs`
- `GET /docs`

#### Error response format:
 - `status`: `4xx`,`5xx`
 - ```json 
   {
       "code": 0,
       "message": "...",
       "data": null
   }
   ```
#### POST /docs
- Request Body:
    - `title: {
           type: String,
           required: Required
        }`
    - `version: {
           type: Integer,
           required: Optional
        }`
     - `method: {
           type: String,
           required: Required
        }`
     - `base_url: {
           type: String,
           required: Required
        }`
     - `path: {
           type: String,
           required: Required
        }`
     - `params: {
           type: Array,
           required: Optional
        }`
        - `key: {
              type: String
           }`
     - `body: {
           type: Array,
           required: Optional
        }`
        - `key: {
              type: String
           }`
        - `type: {
              type: String
           }`
     - `result: {
           type: Object,
           required: Required
        }`
        - `http_code: {
              type: Integer
           }`
        - `type: {
              type: String
           }`
                        
- Example request:
  - ```json
    {
       "title": "book",
       "version": 1,
       "method": "POST",
       "base_url": "http://localhost/api/v1",
       "path": "/books",
       "params": null,
       "headers": [
           {"key": "api-key"},
           {"key": "Accept-Language"},
           {"key": "token"},
           {"key": "Content-Type"}
        ],
       "body": [
           {
              "key": "title",
              "type": "string"
           },
           {
              "key": "sub title",
              "type": "string"
           },
           {
              "key": "author",
              "type": "string"
           }
        ],
       "result": {
           "http_code": "201",
           "type": "json"
        }
    }
    ```
       
- Response:
  - `status`: `201`
  - ```json
     {
        "code": 1,
        "message": "Docs successfully registered",
        "data" : {
            "title": "book",
            "version": 1,
            "method": "POST",
            "base_url": "http://localhost/api/v1",
            "path": "/books",
            "params": null,
            "headers": [
                {"key": "api-key"},
                {"key": "Accept-Language"},
                {"key": "token"},
                {"key": "Content-Type"}
             ],
            "body": [
                {
                   "key": "title",
                   "type": "string"
                },
                {
                   "key": "sub title",
                   "type": "string"
                },
                {
                   "key": "author",
                   "type": "string"
                }
             ],
            "result": {
                "http_code": "201",
                "type": "json"
             }
        }
     }
    ```
  - ```json
       {
           "code": 2,
           "message": "docs already registered",
           "data": null
       }
    ```
    
#### GET /docs
- Query Params :
  - `title: {
         type: String,
         required: Required
      }`
     
- Request:
  - `/docs/{title}`
- Example:
  - `/docs/book`
  
- Response:
  - `status`: `200`
  - ```json
       {
           "code": 1,
           "message": "book",
           "data": [
               {
                   "title": "book",
                   "version": 1,
                   "method": "GET",
                   "base_url": "http://localhost/api/v1",
                   "path": "/books",
                   "params": [
                       {"key": "title"},
                       {"key": "page"},
                       {"key": "limit"}
                    ],
                   "headers": [
                       {"key": "api-key"},
                       {"key": "Accept-Language"},
                       {"key": "token"},
                       {"key": "Content-Type"}
                    ],
                   "body": null,
                   "result": {
                       "http_code": "200",
                       "type": "json"
                    }
               },
               {
                   "title": "book",
                   "version": 1,
                   "method": "POST",
                   "base_url": "http://localhost/api/v1",
                   "path": "/books",
                   "params": null,
                   "headers": [
                       {"key": "api-key"},
                       {"key": "Accept-Language"},
                       {"key": "token"},
                       {"key": "Content-Type"}
                    ],
                   "body": [
                       {
                          "key": "title",
                          "type": "string"
                       },
                       {
                          "key": "sub title",
                          "type": "string"
                       },
                       {
                          "key": "author",
                          "type": "string"
                       }
                    ],
                   "result": {
                       "http_code": "201",
                       "type": "json"
                    }
               }
           ]
       }
    ```
  - ```json
         {
             "code": 3,
             "message": "Docs not found",
             "data": null
         }
    ```