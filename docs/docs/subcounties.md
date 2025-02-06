---
sidebar_position: 6
---

# Subcounties

This endpoint allows you to retrieve subcounties in Uganda.

## Get All Subcounties

Fetches a list of all subcounties in Uganda.

### Endpoint

```
GET https://api.opendataug.org/v1/subcounties
```

### Headers

| Header    | Required | Description                     |
| --------- | -------- | ------------------------------- |
| x-api-key | Yes      | Your API key for authentication |

### Example Request

```bash
curl -X GET \
  'https://api.opendataug.org/v1/subcounties' \
  -H 'x-api-key: your_api_key_here'
```

### Response

```json
{
  "status": "success",
  "data": [
    {
      "id": "1",
      "name": "Central Division",
      "county_id": "1",
      "county_name": "Nakawa",
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
| data                 | array  | Array of subcounty objects                       |
| data[].id            | string | Unique identifier for the subcounty              |
| data[].name          | string | Name of the subcounty                            |
| data[].county_id     | string | ID of the county this subcounty belongs to       |
| data[].county_name   | string | Name of the county this subcounty belongs to     |
| data[].district_id   | string | ID of the district this subcounty belongs to     |
| data[].district_name | string | Name of the district this subcounty belongs to   |

### Error Responses

| Status Code | Description                             |
| ----------- | --------------------------------------- |
| 401         | Invalid or missing API key              |
| 429         | Too many requests - Rate limit exceeded |
| 500         | Internal server error                   |

### Rate Limiting

This endpoint is subject to rate limiting. Please refer to our [rate limiting documentation](./authentication.md#rate-limiting) for more details.

### Notes

- Subcounties are administrative divisions within counties
- The API response is paginated with 100 subcounties per page
- Subcounties are sorted alphabetically by name

## Get Subcounties by County

Fetches all subcounties within a specific county.

### Endpoint

```
GET https://api.opendataug.org/v1/counties/{county_id}/subcounties
```

### Parameters

| Parameter | Type   | Required | Description      |
| --------- | ------ | -------- | ---------------- |
| county_id | string | Yes      | ID of the county |

### Example Request

```bash
curl -X GET \
  'https://api.opendataug.org/v1/counties/1/subcounties' \
  -H 'x-api-key: your_api_key_here'
```

The response format is the same as the Get All Subcounties endpoint.
