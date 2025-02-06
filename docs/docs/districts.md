---
sidebar_position: 4
---

# Districts

This endpoint allows you to retrieve a list of all districts in Uganda.

## Get All Districts

Fetches a list of all districts in Uganda.

### Endpoint

````

### Headers

| Header | Required | Description |
|--------|----------|-------------|
| x-api-key | Yes | Your API key for authentication |

### Example Request

```bash
curl -X GET \
  'https://api.opendataug.org/v1/districts' \
  -H 'x-api-key: your_api_key_here'
````

### Response

```json
{
  "status": "success",
  "data": [
    {
      "id": "1",
      "name": "Kampala",
      "region": "Central"
    },
    {
      "id": "2",
      "name": "Wakiso",
      "region": "Central"
    }
  ]
}
```

### Response Fields

| Field         | Type   | Description                                      |
| ------------- | ------ | ------------------------------------------------ |
| status        | string | The status of the request ("success" or "error") |
| data          | array  | Array of district objects                        |
| data[].id     | string | Unique identifier for the district               |
| data[].name   | string | Name of the district                             |
| data[].region | string | Region where the district is located             |

### Error Responses

| Status Code | Description                             |
| ----------- | --------------------------------------- |
| 401         | Invalid or missing API key              |
| 429         | Too many requests - Rate limit exceeded |
| 500         | Internal server error                   |

### Rate Limiting

This endpoint is subject to rate limiting. Please refer to our [rate limiting documentation](./authentication.md#rate-limiting) for more details.

### Notes

- The list of districts is regularly updated to reflect administrative changes
- Districts are sorted alphabetically by name
- The API response is paginated with 100 districts per page
