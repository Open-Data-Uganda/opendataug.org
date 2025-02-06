---
sidebar_position: 3
---

# Regions

This endpoint allows you to retrieve all regions in Uganda.

## Get All Regions

Fetches a list of all regions in Uganda.

### Endpoint

```
GET https://api.opendataug.org/v1/regions
```

### Headers

| Header    | Required | Description                     |
| --------- | -------- | ------------------------------- |
| x-api-key | Yes      | Your API key for authentication |

### Example Request

```bash
curl -X GET \
  'https://api.opendataug.org/v1/regions' \
  -H 'x-api-key: your_api_key_here'
```

### Response

```json
{
  "status": "success",
  "data": [
    {
      "id": "1",
      "name": "Central",
      "districts_count": 24
    },
    {
      "id": "2",
      "name": "Eastern",
      "districts_count": 32
    }
  ]
}
```

### Response Fields

| Field                  | Type   | Description                                      |
| ---------------------- | ------ | ------------------------------------------------ |
| status                 | string | The status of the request ("success" or "error") |
| data                   | array  | Array of region objects                          |
| data[].id              | string | Unique identifier for the region                 |
| data[].name            | string | Name of the region                               |
| data[].districts_count | number | Number of districts in the region                |

### Error Responses

| Status Code | Description                             |
| ----------- | --------------------------------------- |
| 401         | Invalid or missing API key              |
| 429         | Too many requests - Rate limit exceeded |
| 500         | Internal server error                   |

### Rate Limiting

This endpoint is subject to rate limiting. Please refer to our [rate limiting documentation](./authentication.md#rate-limiting) for more details.

### Notes

- Uganda is divided into four main regions: Central, Eastern, Northern, and Western
- Each region contains multiple districts
