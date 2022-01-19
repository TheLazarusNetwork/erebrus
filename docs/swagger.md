


# Erebrus
The api is used to create,update,delete,retrive clients and
get client configiration and email them. It also provides
server operations like get, update, status and server configuration.

This should demonstrate all the possible comment annotations
that are available to turn go code into a fully compliant swagger 2.0 spec
  

## Informations

### Version

1.0.0

### License

[MIT](http://opensource.org/licenses/MIT)

### Contact

sachindra sachindra.spidey@gmail.com 

## Content negotiation

### URI Schemes
  * http
  * https

### Consumes
  * application/json
  * application/x-protobuf

### Produces
  * application/octet-stream
  * application/json
  * application/x-protobuf

## All endpoints

###  client

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /api/v1.0/client/{id}/config | [config client](#config-client) | Get client configuration |
| POST | /api/v1.0/client | [create client](#create-client) | Create client |
| DELETE | /api/v1.0/client/{id} | [delete client](#delete-client) | Delete client |
| GET | /api/v1.0/client/{id}/email | [email client](#email-client) | Email client Configuration |
| GET | /api/v1.0/client/{id} | [read client](#read-client) | Read client |
| PATCH | /api/v1.0/client/{id} | [update client](#update-client) | Update client |
  


###  serverops

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /api/v1.0/server/config | [config server](#config-server) |  |
| GET | /api/v1.0/server | [read server](#read-server) | Read Server |
| GET | /api/v1.0/server/status | [status server](#status-server) | Get Server status |
| PATCH | /api/v1.0/server | [update server](#update-server) | Update Server |
  


## Paths

### <span id="config-client"></span> Get client configuration (*configClient*)

```
GET /api/v1.0/client/{id}/config
```

Return client configuration file in byte format based on the given uuid.

#### Produces
  * application/json
  * application/octet-stream

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | The Identifier of the Client |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#config-client-200) | OK |  |  | [schema](#config-client-200-schema) |
| [400](#config-client-400) | Bad Request |  |  | [schema](#config-client-400-schema) |
| [401](#config-client-401) | Unauthorized |  |  | [schema](#config-client-401-schema) |
| [502](#config-client-502) | Bad Gateway |  |  | [schema](#config-client-502-schema) |

#### Responses


##### <span id="config-client-200"></span> 200
Status: OK

###### <span id="config-client-200-schema"></span> Schema
   
  

[]uint8 (formatted integer)

##### <span id="config-client-400"></span> 400
Status: Bad Request

###### <span id="config-client-400-schema"></span> Schema
   
  

[ConfigClientBadRequestBody](#config-client-bad-request-body)

##### <span id="config-client-401"></span> 401
Status: Unauthorized

###### <span id="config-client-401-schema"></span> Schema
   
  

[ConfigClientUnauthorizedBody](#config-client-unauthorized-body)

##### <span id="config-client-502"></span> 502
Status: Bad Gateway

###### <span id="config-client-502-schema"></span> Schema
   
  

[ConfigClientBadGatewayBody](#config-client-bad-gateway-body)

###### Inlined models

**<span id="config-client-bad-gateway-body"></span> ConfigClientBadGatewayBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `502` |
| Sucess | boolean| `bool` |  | |  | `false` |



**<span id="config-client-bad-request-body"></span> ConfigClientBadRequestBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `400` |
| Sucess | boolean| `bool` |  | |  | `false` |



**<span id="config-client-unauthorized-body"></span> ConfigClientUnauthorizedBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `401` |
| Sucess | boolean| `bool` |  | |  | `false` |



### <span id="config-server"></span> config server (*configServer*)

```
GET /api/v1.0/server/config
```

Get Server Configuration
Retrieves the server configuration details.

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#config-server-200) | OK |  |  | [schema](#config-server-200-schema) |
| [400](#config-server-400) | Bad Request |  |  | [schema](#config-server-400-schema) |
| [401](#config-server-401) | Unauthorized |  |  | [schema](#config-server-401-schema) |
| [502](#config-server-502) | Bad Gateway |  |  | [schema](#config-server-502-schema) |

#### Responses


##### <span id="config-server-200"></span> 200
Status: OK

###### <span id="config-server-200-schema"></span> Schema

##### <span id="config-server-400"></span> 400
Status: Bad Request

###### <span id="config-server-400-schema"></span> Schema
   
  

[ConfigServerBadRequestBody](#config-server-bad-request-body)

##### <span id="config-server-401"></span> 401
Status: Unauthorized

###### <span id="config-server-401-schema"></span> Schema
   
  

[ConfigServerUnauthorizedBody](#config-server-unauthorized-body)

##### <span id="config-server-502"></span> 502
Status: Bad Gateway

###### <span id="config-server-502-schema"></span> Schema
   
  

[ConfigServerBadGatewayBody](#config-server-bad-gateway-body)

###### Inlined models

**<span id="config-server-bad-gateway-body"></span> ConfigServerBadGatewayBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `502` |
| Sucess | boolean| `bool` |  | |  | `false` |



**<span id="config-server-bad-request-body"></span> ConfigServerBadRequestBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `400` |
| Sucess | boolean| `bool` |  | |  | `false` |



**<span id="config-server-unauthorized-body"></span> ConfigServerUnauthorizedBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `401` |
| Sucess | boolean| `bool` |  | |  | `false` |



### <span id="create-client"></span> Create client (*createClient*)

```
POST /api/v1.0/client
```

Create client based on the given client model.

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| client | `body` | [ClientReq](#client-req) | `models.ClientReq` | |  | | Requestbody  used for create and update client operations. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-client-201) | Created |  |  | [schema](#create-client-201-schema) |
| [400](#create-client-400) | Bad Request |  |  | [schema](#create-client-400-schema) |
| [401](#create-client-401) | Unauthorized |  |  | [schema](#create-client-401-schema) |
| [502](#create-client-502) | Bad Gateway |  |  | [schema](#create-client-502-schema) |

#### Responses


##### <span id="create-client-201"></span> 201
Status: Created

###### <span id="create-client-201-schema"></span> Schema
   
  

[CreateClientCreatedBody](#create-client-created-body)

##### <span id="create-client-400"></span> 400
Status: Bad Request

###### <span id="create-client-400-schema"></span> Schema
   
  

[CreateClientBadRequestBody](#create-client-bad-request-body)

##### <span id="create-client-401"></span> 401
Status: Unauthorized

###### <span id="create-client-401-schema"></span> Schema
   
  

[CreateClientUnauthorizedBody](#create-client-unauthorized-body)

##### <span id="create-client-502"></span> 502
Status: Bad Gateway

###### <span id="create-client-502-schema"></span> Schema
   
  

[CreateClientBadGatewayBody](#create-client-bad-gateway-body)

###### Inlined models

**<span id="create-client-bad-gateway-body"></span> CreateClientBadGatewayBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `502` |
| Sucess | boolean| `bool` |  | |  | `false` |



**<span id="create-client-bad-request-body"></span> CreateClientBadRequestBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `400` |
| Sucess | boolean| `bool` |  | |  | `false` |



**<span id="create-client-created-body"></span> CreateClientCreatedBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Message | string| `string` |  | |  | `sucess message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `201` |
| Sucess | boolean| `bool` |  | |  | `true` |
| client | [Client](#client)| `models.Client` |  | |  |  |



**<span id="create-client-unauthorized-body"></span> CreateClientUnauthorizedBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `401` |
| Sucess | boolean| `bool` |  | |  | `false` |



### <span id="delete-client"></span> Delete client (*deleteClient*)

```
DELETE /api/v1.0/client/{id}
```

Delete client based on the given uuid.

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | The Identifier of the Client |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-client-200) | OK |  |  | [schema](#delete-client-200-schema) |
| [400](#delete-client-400) | Bad Request |  |  | [schema](#delete-client-400-schema) |
| [401](#delete-client-401) | Unauthorized |  |  | [schema](#delete-client-401-schema) |
| [502](#delete-client-502) | Bad Gateway |  |  | [schema](#delete-client-502-schema) |

#### Responses


##### <span id="delete-client-200"></span> 200
Status: OK

###### <span id="delete-client-200-schema"></span> Schema
   
  

[DeleteClientOKBody](#delete-client-o-k-body)

##### <span id="delete-client-400"></span> 400
Status: Bad Request

###### <span id="delete-client-400-schema"></span> Schema
   
  

[DeleteClientBadRequestBody](#delete-client-bad-request-body)

##### <span id="delete-client-401"></span> 401
Status: Unauthorized

###### <span id="delete-client-401-schema"></span> Schema
   
  

[DeleteClientUnauthorizedBody](#delete-client-unauthorized-body)

##### <span id="delete-client-502"></span> 502
Status: Bad Gateway

###### <span id="delete-client-502-schema"></span> Schema
   
  

[DeleteClientBadGatewayBody](#delete-client-bad-gateway-body)

###### Inlined models

**<span id="delete-client-bad-gateway-body"></span> DeleteClientBadGatewayBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `502` |
| Sucess | boolean| `bool` |  | |  | `false` |



**<span id="delete-client-bad-request-body"></span> DeleteClientBadRequestBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `400` |
| Sucess | boolean| `bool` |  | |  | `false` |



**<span id="delete-client-o-k-body"></span> DeleteClientOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Message | string| `string` |  | |  | `sucess message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `200` |
| Sucess | boolean| `bool` |  | |  | `true` |



**<span id="delete-client-unauthorized-body"></span> DeleteClientUnauthorizedBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `401` |
| Sucess | boolean| `bool` |  | |  | `false` |



### <span id="email-client"></span> Email client Configuration (*emailClient*)

```
GET /api/v1.0/client/{id}/email
```

Email the configuration file of the client to the email associated with client.

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | The Identifier of the Client |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#email-client-200) | OK |  |  | [schema](#email-client-200-schema) |
| [400](#email-client-400) | Bad Request |  |  | [schema](#email-client-400-schema) |
| [401](#email-client-401) | Unauthorized |  |  | [schema](#email-client-401-schema) |
| [502](#email-client-502) | Bad Gateway |  |  | [schema](#email-client-502-schema) |

#### Responses


##### <span id="email-client-200"></span> 200
Status: OK

###### <span id="email-client-200-schema"></span> Schema
   
  

[EmailClientOKBody](#email-client-o-k-body)

##### <span id="email-client-400"></span> 400
Status: Bad Request

###### <span id="email-client-400-schema"></span> Schema
   
  

[EmailClientBadRequestBody](#email-client-bad-request-body)

##### <span id="email-client-401"></span> 401
Status: Unauthorized

###### <span id="email-client-401-schema"></span> Schema
   
  

[EmailClientUnauthorizedBody](#email-client-unauthorized-body)

##### <span id="email-client-502"></span> 502
Status: Bad Gateway

###### <span id="email-client-502-schema"></span> Schema
   
  

[EmailClientBadGatewayBody](#email-client-bad-gateway-body)

###### Inlined models

**<span id="email-client-bad-gateway-body"></span> EmailClientBadGatewayBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `502` |
| Sucess | boolean| `bool` |  | |  | `false` |



**<span id="email-client-bad-request-body"></span> EmailClientBadRequestBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `400` |
| Sucess | boolean| `bool` |  | |  | `false` |



**<span id="email-client-o-k-body"></span> EmailClientOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Message | string| `string` |  | |  | `sucess message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `200` |
| Sucess | boolean| `bool` |  | |  | `true` |



**<span id="email-client-unauthorized-body"></span> EmailClientUnauthorizedBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `401` |
| Sucess | boolean| `bool` |  | |  | `false` |



### <span id="read-client"></span> Read client (*readClient*)

```
GET /api/v1.0/client/{id}
```

Return client based on the given uuid.

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | The Identifier of the Client |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#read-client-200) | OK |  |  | [schema](#read-client-200-schema) |
| [400](#read-client-400) | Bad Request |  |  | [schema](#read-client-400-schema) |
| [401](#read-client-401) | Unauthorized |  |  | [schema](#read-client-401-schema) |
| [502](#read-client-502) | Bad Gateway |  |  | [schema](#read-client-502-schema) |

#### Responses


##### <span id="read-client-200"></span> 200
Status: OK

###### <span id="read-client-200-schema"></span> Schema
   
  

[ReadClientOKBody](#read-client-o-k-body)

##### <span id="read-client-400"></span> 400
Status: Bad Request

###### <span id="read-client-400-schema"></span> Schema
   
  

[ReadClientBadRequestBody](#read-client-bad-request-body)

##### <span id="read-client-401"></span> 401
Status: Unauthorized

###### <span id="read-client-401-schema"></span> Schema
   
  

[ReadClientUnauthorizedBody](#read-client-unauthorized-body)

##### <span id="read-client-502"></span> 502
Status: Bad Gateway

###### <span id="read-client-502-schema"></span> Schema
   
  

[ReadClientBadGatewayBody](#read-client-bad-gateway-body)

###### Inlined models

**<span id="read-client-bad-gateway-body"></span> ReadClientBadGatewayBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `502` |
| Sucess | boolean| `bool` |  | |  | `false` |



**<span id="read-client-bad-request-body"></span> ReadClientBadRequestBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `400` |
| Sucess | boolean| `bool` |  | |  | `false` |



**<span id="read-client-o-k-body"></span> ReadClientOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Message | string| `string` |  | |  | `sucess message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `201` |
| Sucess | boolean| `bool` |  | |  | `true` |
| client | [Client](#client)| `models.Client` |  | |  |  |



**<span id="read-client-unauthorized-body"></span> ReadClientUnauthorizedBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `401` |
| Sucess | boolean| `bool` |  | |  | `false` |



### <span id="read-server"></span> Read Server (*readServer*)

```
GET /api/v1.0/server
```

Retrieves the server details.

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#read-server-200) | OK |  |  | [schema](#read-server-200-schema) |
| [400](#read-server-400) | Bad Request |  |  | [schema](#read-server-400-schema) |
| [401](#read-server-401) | Unauthorized |  |  | [schema](#read-server-401-schema) |
| [502](#read-server-502) | Bad Gateway |  |  | [schema](#read-server-502-schema) |

#### Responses


##### <span id="read-server-200"></span> 200
Status: OK

###### <span id="read-server-200-schema"></span> Schema
   
  

[ReadServerOKBody](#read-server-o-k-body)

##### <span id="read-server-400"></span> 400
Status: Bad Request

###### <span id="read-server-400-schema"></span> Schema
   
  

[ReadServerBadRequestBody](#read-server-bad-request-body)

##### <span id="read-server-401"></span> 401
Status: Unauthorized

###### <span id="read-server-401-schema"></span> Schema
   
  

[ReadServerUnauthorizedBody](#read-server-unauthorized-body)

##### <span id="read-server-502"></span> 502
Status: Bad Gateway

###### <span id="read-server-502-schema"></span> Schema
   
  

[ReadServerBadGatewayBody](#read-server-bad-gateway-body)

###### Inlined models

**<span id="read-server-bad-gateway-body"></span> ReadServerBadGatewayBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `502` |
| Sucess | boolean| `bool` |  | |  | `false` |



**<span id="read-server-bad-request-body"></span> ReadServerBadRequestBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `400` |
| Sucess | boolean| `bool` |  | |  | `false` |



**<span id="read-server-o-k-body"></span> ReadServerOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Message | string| `string` |  | |  | `sucess message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `201` |
| Sucess | boolean| `bool` |  | |  | `true` |
| server | [Server](#server)| `models.Server` |  | |  |  |



**<span id="read-server-unauthorized-body"></span> ReadServerUnauthorizedBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `401` |
| Sucess | boolean| `bool` |  | |  | `false` |



### <span id="status-server"></span> Get Server status (*statusServer*)

```
GET /api/v1.0/server/status
```

Retrieves the server  status details.

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#status-server-200) | OK |  |  | [schema](#status-server-200-schema) |
| [400](#status-server-400) | Bad Request |  |  | [schema](#status-server-400-schema) |
| [401](#status-server-401) | Unauthorized |  |  | [schema](#status-server-401-schema) |
| [502](#status-server-502) | Bad Gateway |  |  | [schema](#status-server-502-schema) |

#### Responses


##### <span id="status-server-200"></span> 200
Status: OK

###### <span id="status-server-200-schema"></span> Schema
   
  

[Status](#status)

##### <span id="status-server-400"></span> 400
Status: Bad Request

###### <span id="status-server-400-schema"></span> Schema
   
  

[StatusServerBadRequestBody](#status-server-bad-request-body)

##### <span id="status-server-401"></span> 401
Status: Unauthorized

###### <span id="status-server-401-schema"></span> Schema
   
  

[StatusServerUnauthorizedBody](#status-server-unauthorized-body)

##### <span id="status-server-502"></span> 502
Status: Bad Gateway

###### <span id="status-server-502-schema"></span> Schema
   
  

[StatusServerBadGatewayBody](#status-server-bad-gateway-body)

###### Inlined models

**<span id="status-server-bad-gateway-body"></span> StatusServerBadGatewayBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `502` |
| Sucess | boolean| `bool` |  | |  | `false` |



**<span id="status-server-bad-request-body"></span> StatusServerBadRequestBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `400` |
| Sucess | boolean| `bool` |  | |  | `false` |



**<span id="status-server-unauthorized-body"></span> StatusServerUnauthorizedBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `401` |
| Sucess | boolean| `bool` |  | |  | `false` |



### <span id="update-client"></span> Update client (*updateClient*)

```
PATCH /api/v1.0/client/{id}
```

Update client based on the given uuid and client model.

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| id | `path` | string | `string` |  | ✓ |  | The Identifier of the Client |
| client | `body` | [ClientUpdateReq](#client-update-req) | `models.ClientUpdateReq` | |  | | Requestbody  used for create and update client operations. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-client-200) | OK |  |  | [schema](#update-client-200-schema) |
| [400](#update-client-400) | Bad Request |  |  | [schema](#update-client-400-schema) |
| [401](#update-client-401) | Unauthorized |  |  | [schema](#update-client-401-schema) |
| [502](#update-client-502) | Bad Gateway |  |  | [schema](#update-client-502-schema) |

#### Responses


##### <span id="update-client-200"></span> 200
Status: OK

###### <span id="update-client-200-schema"></span> Schema
   
  

[UpdateClientOKBody](#update-client-o-k-body)

##### <span id="update-client-400"></span> 400
Status: Bad Request

###### <span id="update-client-400-schema"></span> Schema
   
  

[UpdateClientBadRequestBody](#update-client-bad-request-body)

##### <span id="update-client-401"></span> 401
Status: Unauthorized

###### <span id="update-client-401-schema"></span> Schema
   
  

[UpdateClientUnauthorizedBody](#update-client-unauthorized-body)

##### <span id="update-client-502"></span> 502
Status: Bad Gateway

###### <span id="update-client-502-schema"></span> Schema
   
  

[UpdateClientBadGatewayBody](#update-client-bad-gateway-body)

###### Inlined models

**<span id="update-client-bad-gateway-body"></span> UpdateClientBadGatewayBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `502` |
| Sucess | boolean| `bool` |  | |  | `false` |



**<span id="update-client-bad-request-body"></span> UpdateClientBadRequestBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `400` |
| Sucess | boolean| `bool` |  | |  | `false` |



**<span id="update-client-o-k-body"></span> UpdateClientOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Message | string| `string` |  | |  | `sucess message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `201` |
| Sucess | boolean| `bool` |  | |  | `true` |
| client | [Client](#client)| `models.Client` |  | |  |  |



**<span id="update-client-unauthorized-body"></span> UpdateClientUnauthorizedBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `401` |
| Sucess | boolean| `bool` |  | |  | `false` |



### <span id="update-server"></span> Update Server (*updateServer*)

```
PATCH /api/v1.0/server
```

Update the server with given details.

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| server | `body` | [Server](#server) | `models.Server` | |  | | Requestbody  used for update server operations. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-server-200) | OK |  |  | [schema](#update-server-200-schema) |
| [400](#update-server-400) | Bad Request |  |  | [schema](#update-server-400-schema) |
| [401](#update-server-401) | Unauthorized |  |  | [schema](#update-server-401-schema) |
| [502](#update-server-502) | Bad Gateway |  |  | [schema](#update-server-502-schema) |

#### Responses


##### <span id="update-server-200"></span> 200
Status: OK

###### <span id="update-server-200-schema"></span> Schema
   
  

[UpdateServerOKBody](#update-server-o-k-body)

##### <span id="update-server-400"></span> 400
Status: Bad Request

###### <span id="update-server-400-schema"></span> Schema
   
  

[UpdateServerBadRequestBody](#update-server-bad-request-body)

##### <span id="update-server-401"></span> 401
Status: Unauthorized

###### <span id="update-server-401-schema"></span> Schema
   
  

[UpdateServerUnauthorizedBody](#update-server-unauthorized-body)

##### <span id="update-server-502"></span> 502
Status: Bad Gateway

###### <span id="update-server-502-schema"></span> Schema
   
  

[UpdateServerBadGatewayBody](#update-server-bad-gateway-body)

###### Inlined models

**<span id="update-server-bad-gateway-body"></span> UpdateServerBadGatewayBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `502` |
| Sucess | boolean| `bool` |  | |  | `false` |



**<span id="update-server-bad-request-body"></span> UpdateServerBadRequestBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `400` |
| Sucess | boolean| `bool` |  | |  | `false` |



**<span id="update-server-o-k-body"></span> UpdateServerOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Message | string| `string` |  | |  | `sucess message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `201` |
| Sucess | boolean| `bool` |  | |  | `true` |
| server | [Server](#server)| `models.Server` |  | |  |  |



**<span id="update-server-unauthorized-body"></span> UpdateServerUnauthorizedBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  | `error message` |
| Status | int64 (formatted integer)| `int64` |  | |  | `401` |
| Sucess | boolean| `bool` |  | |  | `false` |



## Models

### <span id="client"></span> Client


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Address | []string| `[]string` |  | | Address range client must will assigned | `["10.0.0.2/32"]` |
| AllowedIPs | []string| `[]string` |  | | IP addresses allowed to connect | `["0.0.0.0/0","::/0"]` |
| Created | int64 (formatted integer)| `int64` |  | | Time the client is created | `1642409076544` |
| CreatedBy | string| `string` |  | | Denoting person creates the client | `jonsnow@mail.com` |
| Email | string| `string` |  | | Email that the client device belongs | `jonsnow@mail.com` |
| Enable | boolean| `bool` |  | | Status signal for client | `true` |
| IgnorePersistentKeepalive | boolean| `bool` |  | |  | `true` |
| Name | string| `string` |  | | Name of the client | `jon snow` |
| PresharedKey | string| `string` |  | | Preshared key for the client | `twDZk0lehYtst3Zclb+SRniVfoHnug9N6gjxuaipcvc=` |
| PrivateKey | string| `string` |  | | Private key for the client | `KFOyCoR9Eq+LpqT9VzJCilXYmFwhMFw7UDkdRRxoWVg=` |
| PublicKey | string| `string` |  | | Public key for the client | `YeT/lG9L4AeYOHNrkohnmXfljx3/JgThulskllayxi4=` |
| Tags | []string| `[]string` |  | | Tags for client device | `["laptop","PC"]` |
| UUID | string| `string` |  | | Client identifier | `6c8ff96f-ce8a-4c64-a76d-07e9af0b75ab` |
| Updated | int64 (formatted integer)| `int64` |  | | Time the client is last updated | `1642409076544` |
| UpdatedBy | string| `string` |  | | Denoting person updates the client | `jonsnow@mail.com` |



### <span id="client-req"></span> ClientReq


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Address | []string| `[]string` | ✓ | | Address range client must will assigned | `["10.0.0.0/24"]` |
| AllowedIPs | []string| `[]string` | ✓ | | IP addresses allowed to connect | `["0.0.0.0/0","::/0"]` |
| CreatedBy | string| `string` | ✓ | | Denoting person creates the client | `jonsnow@mail.com` |
| Email | string| `string` | ✓ | | Email that the client device belongs | `jonsnow@mail.com` |
| Enable | boolean| `bool` | ✓ | | Status signal for client | `true` |
| Name | string| `string` | ✓ | |  | `jon snow` |
| Tags | []string| `[]string` | ✓ | | Tags for client device | `["laptop","PC"]` |
| UpdatedBy | string| `string` | ✓ | | Denoting person updates the client | `jonsnow@mail.com` |



### <span id="client-update-req"></span> ClientUpdateReq


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Address | []string| `[]string` | ✓ | | IP addresses allowed to connect | `["10.0.0.2/32"]` |
| AllowedIPs | []string| `[]string` | ✓ | | IP addresses allowed to connect | `["0.0.0.0/0","::/0"]` |
| Created | int64 (formatted integer)| `int64` |  | | Time the client is created | `1642409076544` |
| CreatedBy | string| `string` |  | | Denoting person creates the client | `jonsnow@mail.com` |
| Email | string| `string` | ✓ | | Email that the client device belongs | `jonsnow@mail.com` |
| Enable | boolean| `bool` | ✓ | | Status signal for client | `true` |
| IgnorePersistentKeepalive | boolean| `bool` |  | |  | `true` |
| Name | string| `string` | ✓ | | Name of the client | `jon snow` |
| PresharedKey | string| `string` |  | | Preshared key for the client | `twDZk0lehYtst3Zclb+SRniVfoHnug9N6gjxuaipcvc=` |
| PrivateKey | string| `string` |  | | Private key for the client | `KFOyCoR9Eq+LpqT9VzJCilXYmFwhMFw7UDkdRRxoWVg=` |
| PublicKey | string| `string` |  | | Public key for the client | `YeT/lG9L4AeYOHNrkohnmXfljx3/JgThulskllayxi4=` |
| Tags | []string| `[]string` | ✓ | | Tags for client device | `["laptop","PC"]` |
| UUID | string| `string` | ✓ | | Client identifier | `6c8ff96f-ce8a-4c64-a76d-07e9af0b75ab` |
| Updated | int64 (formatted integer)| `int64` |  | | Time the client is last updated | `1642409076544` |
| UpdatedBy | string| `string` | ✓ | | Denoting person updates the client | `jonsnow@mail.com` |



### <span id="server"></span> Server


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Address | []string| `[]string` |  | | Server address | `["10.0.0.1/24"]` |
| AllowedIPs | []string| `[]string` |  | | IP addresses allowed to connect | `["0.0.0.0/0","::/0"]` |
| Created | int64 (formatted integer)| `int64` |  | | Time when server is created | `26103870` |
| DNS | []string| `[]string` |  | | DNS of the VPN server | `["1.1.1.1"]` |
| Endpoint | string| `string` |  | | Endpoint of the server | `region.example.com` |
| ListenPort | int64 (formatted integer)| `int64` |  | | Port the server listens | `51280` |
| Mtu | int64 (formatted integer)| `int64` |  | |  |  |
| PersistentKeepalive | int64 (formatted integer)| `int64` |  | | Persistent keep alive for server | `16` |
| PostDown | string| `string` |  | | Post down command | `iptables -D FORWARD -i %i -j ACCEPT; iptables -D FORWARD -o %i -j ACCEPT; iptables -t nat -D POSTROUTING -o eth0 -j MASQUERADE` |
| PostUp | string| `string` |  | | Post up command | `iptables -A FORWARD -i %i -j ACCEPT; iptables -A FORWARD -o %i -j ACCEPT; iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE` |
| PreDown | string| `string` |  | | Pre down command | `echo WireGuard PreDown` |
| PreUp | string| `string` |  | | Pre up command | `echo WireGuard PreUp` |
| PrivateKey | string| `string` |  | | Private key for the server | `UFWsgb/Ax5B8zZGx0YtHBAuQVRrOHrxKz2zS2p1LuUE=` |
| PublicKey | string| `string` |  | | Public key for the server | `T5ZMOnik3YuaRhZgAhcxXrmn2+C0B7qFaqnCypMMcks=` |
| Updated | int64 (formatted integer)| `int64` |  | | Time when server is created | `26103870` |
| UpdatedBy | string| `string` |  | | Updater email address | `admin@mail.com` |



### <span id="status"></span> Status


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Domain | string| `string` |  | | Domain which server is running | `vpn.example.com` |
| GRPCPort | string| `string` |  | | Port which gRPC service is running | `5000` |
| Hostname | string| `string` |  | | Server Hostname | `ubuntu` |
| HttpPort | string| `string` |  | | Port which HTTP service is running | `4000` |
| PrivateIP | string| `string` |  | | Private IP of server host | `10.0.1.5` |
| PublicIP | string| `string` |  | | Server's public IP | `14.10.35.65` |
| Region | string| `string` |  | | Region where server running | `India/Banglore` |
| VPNPort | string| `string` |  | | VPN port | `5128` |
| Version | string| `string` |  | | Server version | `1.0` |


