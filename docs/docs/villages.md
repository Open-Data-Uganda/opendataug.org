---
sidebar_position: 8
---

# Villages

This endpoint allows you to retrieve villages in Uganda.

## Get All Villages

Fetches a list of all villages in Uganda.

### Endpoint

```
GET https://api.opendataug.org/v1/villages
```

### Headers

| Header    | Required | Description                     |
| --------- | -------- | ------------------------------- |
| x-api-key | Yes      | Your API key for authentication |

### Example Request

```bash
curl -X GET \
  'https://api.opendataug.org/v1/villages' \
  -H 'x-api-key: your_api_key_here'
```

### Response

```json
{
  "status": "success",
  "data": [
    {
      "id": "1",
      "name": "Kiwafu",
      "parish_id": "1",
      "parish_name": "Bukoto",
      "subcounty_id": "1",
      "subcounty_name": "Nakawa"
    }
  ]
}
```

### Response Fields

| Field                 | Type   | Description                                      |
| --------------------- | ------ | ------------------------------------------------ |
| status                | string | The status of the request ("success" or "error") |
| data                  | array  | Array of village objects                         |
| data[].id             | string | Unique identifier for the village                |
| data[].name           | string | Name of the village                              |
| data[].parish_id      | string | ID of the parish this village belongs to         |
| data[].parish_name    | string | Name of the parish this village belongs to       |
| data[].subcounty_id   | string | ID of the subcounty this village belongs to      |
| data[].subcounty_name | string | Name of the subcounty this village belongs to    |

### Error Responses

| Status Code | Description                             |
| ----------- | --------------------------------------- |
| 401         | Invalid or missing API key              |
| 429         | Too many requests - Rate limit exceeded |
| 500         | Internal server error                   |

### Rate Limiting

This endpoint is subject to rate limiting. Please refer to our [rate limiting documentation](./authentication.md#rate-limiting) for more details.

### Notes

- Villages are the smallest administrative units in Uganda
- The API response is paginated with 100 villages per page
- Villages are sorted alphabetically by name

## Get Villages by Parish

Fetches all villages within a specific parish.

### Endpoint

```
GET https://api.opendataug.org/v1/parishes/{parish_id}/villages
```

### Parameters

| Parameter | Type   | Required | Description      |
| --------- | ------ | -------- | ---------------- |
| parish_id | string | Yes      | ID of the parish |

### Example Request

```bash
curl -X GET \
  'https://api.opendataug.org/v1/parishes/1/villages' \
  -H 'x-api-key: your_api_key_here'
```

The response format is the same as the Get All Villages endpoint.
