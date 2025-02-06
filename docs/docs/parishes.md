---
sidebar_position: 7
---

# Parishes

This endpoint allows you to retrieve parishes in Uganda.

## Get All Parishes

Fetches a list of all parishes in Uganda.

### Endpoint

```
GET https://api.opendataug.org/v1/parishes
```

### Headers

| Header    | Required | Description                     |
| --------- | -------- | ------------------------------- |
| x-api-key | Yes      | Your API key for authentication |

### Example Request

```bash
curl -X GET \
  'https://api.opendataug.org/v1/parishes' \
  -H 'x-api-key: your_api_key_here'
```

### Response

```json
{
  "status": "success",
  "data": [
    {
      "id": "1",
      "name": "Bukoto",
      "subcounty_id": "1",
      "subcounty_name": "Nakawa",
      "county_id": "1",
      "county_name": "Nakawa"
    }
  ]
}
```

### Response Fields

| Field                 | Type   | Description                                      |
| --------------------- | ------ | ------------------------------------------------ |
| status                | string | The status of the request ("success" or "error") |
| data                  | array  | Array of parish objects                          |
| data[].id             | string | Unique identifier for the parish                 |
| data[].name           | string | Name of the parish                               |
| data[].subcounty_id   | string | ID of the subcounty this parish belongs to       |
| data[].subcounty_name | string | Name of the subcounty this parish belongs to     |
| data[].county_id      | string | ID of the county this parish belongs to          |
| data[].county_name    | string | Name of the county this parish belongs to        |

### Error Responses

| Status Code | Description                             |
| ----------- | --------------------------------------- |
| 401         | Invalid or missing API key              |
| 429         | Too many requests - Rate limit exceeded |
| 500         | Internal server error                   |

### Rate Limiting

This endpoint is subject to rate limiting. Please refer to our [rate limiting documentation](./authentication.md#rate-limiting) for more details.

### Notes

- Parishes are administrative divisions within subcounties
- The API response is paginated with 100 parishes per page
- Parishes are sorted alphabetically by name

## Get Parishes by Subcounty

Fetches all parishes within a specific subcounty.

### Endpoint

```
GET https://api.opendataug.org/v1/subcounties/{subcounty_id}/parishes
```

### Parameters

| Parameter    | Type   | Required | Description         |
| ------------ | ------ | -------- | ------------------- |
| subcounty_id | string | Yes      | ID of the subcounty |

### Example Request

```bash
curl -X GET \
  'https://api.opendataug.org/v1/subcounties/1/parishes' \
  -H 'x-api-key: your_api_key_here'
```

The response format is the same as the Get All Parishes endpoint.
