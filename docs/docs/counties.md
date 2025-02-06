---
sidebar_position: 5
---

# Counties

This endpoint allows you to retrieve counties in Uganda.

## Get All Counties

Fetches a list of all counties in Uganda.

### Endpoint

```
GET https://api.opendataug.org/v1/counties
```

### Headers

| Header    | Required | Description                     |
| --------- | -------- | ------------------------------- |
| x-api-key | Yes      | Your API key for authentication |

### Example Request

```bash
curl -X GET \
  'https://api.opendataug.org/v1/counties' \
  -H 'x-api-key: your_api_key_here'
```

### Response

```json
{
  "status": "success",
  "data": [
    {
      "id": "1",
      "name": "Nakawa",
      "district_id": "1",
      "district_name": "Kampala"
    },
    {
      "id": "2",
      "name": "Kawempe",
      "district_id": "1",
      "district_name": "Kampala"
    }
  ]
}
```

### Response Fields

| Field                | Type   | Description                                      |
| -------------------- | ------ | ------------------------------------------------ |
| status               | string | The status of the request ("success" or "error") |
| data                 | array  | Array of county objects                          |
| data[].id            | string | Unique identifier for the county                 |
| data[].name          | string | Name of the county                               |
| data[].district_id   | string | ID of the district this county belongs to        |
| data[].district_name | string | Name of the district this county belongs to      |

### Error Responses

| Status Code | Description                             |
| ----------- | --------------------------------------- |
| 401         | Invalid or missing API key              |
| 429         | Too many requests - Rate limit exceeded |
| 500         | Internal server error                   |

### Rate Limiting

This endpoint is subject to rate limiting. Please refer to our [rate limiting documentation](./authentication.md#rate-limiting) for more details.

### Notes

- Counties are administrative divisions within districts
- The API response is paginated with 100 counties per page
- Counties are sorted alphabetically by name

## Get Counties by District

Fetches all counties within a specific district.

### Endpoint

```
GET https://api.opendataug.org/v1/districts/{district_id}/counties
```

### Parameters

| Parameter   | Type   | Required | Description        |
| ----------- | ------ | -------- | ------------------ |
| district_id | string | Yes      | ID of the district |

### Example Request

```bash
curl -X GET \
  'https://api.opendataug.org/v1/districts/1/counties' \
  -H 'x-api-key: your_api_key_here'
```

The response format is the same as the Get All Counties endpoint.
