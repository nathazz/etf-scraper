
# 📊 ETF Scraper API

A web scraper and REST API for extracting real-time data on **ETFs (Exchange-Traded Funds)** from [TrackingDifferences.com](https://www.trackingdifferences.com).  
This API provides endpoints to:

- ✅ Fetch individual ETF details  
- 📦 Retrieve information on multiple ETFs  
- 📄 Generate PDF reports  
- 🔄 Compare multiple ETFs

---

##  Endpoints:

#### `GET /` 

Health Check
Confirms the API is running.

---

#### `GET /etf/:isin`

Fetch a specific ETF by ISIN

**Path Parameter:**
- `isin` – The ISIN code of the ETF (e.g. `LU1737652310`)


#### `GET /etf/LU1737652310`

Fetch detailed information about a specific ETF.

---

#### `POST /etf`

Fetch details for up to 10 ETFs at once.

**Request Body:**

```json
{
  "isins": [
    "LU1737652310",
    "IE00B4K48X80",
    "IE00B1YZSC51",
    "LU0446734104"
  ]
}
```

---

#### `POST /generate-pdf`

Generate a PDF report of up to 10 ETFs

**Request Body:**

```json
{
  "isins": [
    "LU1737652310",
    "IE00B4K48X80",
    "IE00B1YZSC51",
    "LU0446734104"
  ]
}
```
