---
slug: /
sidebar_position: 1
---

# Administrative Units API

Access comprehensive data about Uganda's districts, villages, and administrative units through a simple API.

## Overview

The **Administrative Units API** provides developers with programmatic access to detailed information about Uganda's administrative divisions and geographical data. Whether you're building applications for local governance, analyzing demographic patterns, or developing location-based services, our API delivers accurate and up-to-date data about Uganda's geography.

## Key Features

Access reliable data about:

- Districts
- Counties
- Sub-counties
- Parishes, and
- Villages

## Quick Start

1. **Sign up** - Create an account to get your API key [https://app.opendataug.org/](https://app.opendataug.org/)
2. **Authentication** - Add your API key to request headers
3. **Make requests** - Start querying our endpoints
4. **Handle responses** - Process returned JSON data

### Example Request

```bash
# Fetch all districts
curl -X GET https://api.opendataug.com/v1/districts \
  -H "x-api-key: Bearer YOUR_API_KEY"
```

## Next Steps

Ready to dive in? Check out our:

- [Authentication Guide](authentication.md)
