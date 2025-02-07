---
sidebar_position: 2
---

# Authentication

Learn how to authenticate your requests to our Uganda Districts API. All API endpoints require authentication using an API key.

## Authentication Methods

### API Keys

To access the API, you'll need to include your API key in the `Authorization` header of each request. API keys provide a simple and secure way to authenticate your requests.

#### Getting an API Key

To obtain an API key:

1. Create an account on our platform
2. Navigate to your account dashboard
3. Go to the "API Keys" section
4. Click "Generate New API Key"
5. Save your API key securely - it will only be shown once

#### Using Your API Key

Include your API key in all API requests using the `x-api-key` header with an API Key:

```bash
curl -H "x-api-key: your_api_key_here" \
  https://api.opendataug.com/v1/counties
```

#### API Key Format

API keys follow this format: `opu_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`

- `uk`: Identifies our Uganda Districts API
- Random string: Unique identifier for your API key

## Security Best Practices

To ensure secure usage of our API, follow these best practices:

- **Protect Your API Keys**: Never share your API keys or commit them to version control
- **Use Environment Variables**: Store API keys as environment variables instead of hardcoding them
- **Separate Keys for Different Environments**: Use different API keys for development and production
- **Regular Rotation**: Rotate your API keys periodically for enhanced security

### Code Examples

Here's how to properly use your API key in different programming languages:

```javascript
const apiKey = process.env.UGANDA_API_KEY;
fetch("https://api.example.com/v1/districts", {
  headers: {
    "x-api-key": apiKey,
  },
});
```

**Python**

````python
import os
import requests

api_key = os.getenv('UGANDA_API_KEY')
headers = {'x-api-key': api_key}
response = requests.get('https://api.example.com/v1/districts', headers=headers)


**Go**
```go
import (
    "net/http"
    "os"
)

apiKey := os.Getenv("UGANDA_API_KEY")
req, := http.NewRequest("GET", "https://api.example.com/v1/districts", nil)
req.Header.Add("x-api-key", apiKey)
````

**PHP**

```php
<?php
$api_key = getenv('UGANDA_API_KEY');

$curl = curl_init();
curl_setopt_array($curl, [
    CURLOPT_URL => 'https://api.example.com/v1/districts',
    CURLOPT_RETURNTRANSFER => true,
    CURLOPT_HTTPHEADER => [
        'x-api-key: ' . $api_key
    ]
]);

$response = curl_exec($curl);
curl_close($curl);
```

## Error Responses

When authentication fails, you'll receive one of these responses:

```json
// 401 Unauthorized - Missing or invalid API key
{
"error": {
"code": "unauthorized",
"message": "No API key provided or invalid API key"
}
}


// 403 Forbidden - API key lacks required permissions
{
"error": {
"code": "forbidden",
"message": "API key doesn't have permission to access this resource"
}
}
```

## Rate Limiting

API requests are rate-limited based on your API key:

- Free tier: 1000 requests per hour

---

# Troubleshooting

If you're experiencing authentication issues:

1. Verify your API key is valid and active
2. Check the Authorization header format
3. Ensure you're not exceeding rate limits
4. Confirm your API key has the necessary permissions

For additional help, contact our support team at support@example.com or visit our [GitHub repository](https://github.com/example/uganda-api) for more examples.
