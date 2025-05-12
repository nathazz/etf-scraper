# 📊 ETF Scraper API

A web scraper and REST API for extracting real-time data on **ETFs (Exchange-Traded Funds)** from [TrackingDifferences.com](https://www.trackingdifferences.com).
This API provides endpoints to:

- Fetch individual ETF details
- Retrieve information on multiple ETFs
- Generate PDF reports
- Compare multiple ETFs

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

**response:**

```json
[
    {
        "isin": "LU1737652310",
        "title": "Amundi Index MSCI Europe UCITS ETF DR (D)",
        "replication": "physical",
        "earnings": "distributing",
        "total_expense_ratio": "0.12 %",
        "tracking_difference": "-0.24 %",
        "fund_size": "307 Mio EUR",
        "description": "\"Amundi Index MSCI Europe UCITS ETF DR..."
    }
 ]
```

---


#### both post routes receive this body:

**example**:
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

#### `POST /etf `

This endpoint allows you to fetch details of up to 10 ETFs at once, by providing a list of ISINs.

**response:**

```json
[
    {
        "isin": "LU0446734104",
        "title": "UBS ETF (LU) MSCI Europe UCITS ETF (EUR) A-dis",
        "replication": "physical",
        "earnings": "distributing",
        "total_expense_ratio": "0.10 %",
        "tracking_difference": "-0.11 %",
        "fund_size": "309 Mio EUR",
        "description": "UBS ETF (LU) MSCI Europe UCITS ETF (EUR)..."
    },
    {
        "isin": "LU1737652310",
        "title": "Amundi Index MSCI Europe UCITS ETF DR (D)",
        "replication": "physical",
        "earnings": "distributing",
        "total_expense_ratio": "0.12 %",
        "tracking_difference": "-0.24 %",
        "fund_size": "307 Mio EUR",
        "description": "Amundi Index MSCI Europe UCITS ETF DR (D)..."
    }
]
    "and more..."

```

---

#### `POST /generate-pdf`

This endpoint generates a PDF report containing information for up to 10 ETFs. The PDF includes details such as:

- ISIN
- ETF Name
- Tracking Difference
- Total Expense Ratio (TER)
- Fund Size
- ETF Description

Note: the PDF also includes the comparison between them

**example:**  [etf in pdf](https://drive.google.com/file/d/1Q97UBHKY3V7aUvGFVrvgNXkcH4vFl-9c/view?usp=sharing)

---

#### `POST /compare-etff`

Generate multiple comparisons of ETFs ( limit of 10 ETFs)

**response:**

```json
{
  "tracking_difference": [
    {
      "title": "Amundi Index MSCI Europe UCITS ETF DR (D)",
      "value": -0.24
    },
    {
      "title": "UBS ETF (LU) MSCI Pacific (ex Japan) UCITS ETF (USD) A-dis",
      "value": 0.27
    }
  ],
  "total_expense_ratio": [
    {
      "title": "Amundi Index MSCI Europe UCITS ETF DR (D)",
      "value": 0.12
    },
    {
      "title": "UBS ETF (LU) MSCI Pacific (ex Japan) UCITS ETF (USD) A-dis",
      "value": 0.14
    }
  ],
  "fund_size": [
    {
      "title": "Amundi Index MSCI Europe UCITS ETF DR (D)",
      "value": 0.307
    },
    {
      "title": "UBS ETF (LU) MSCI Pacific (ex Japan) UCITS ETF (USD) A-dis",
      "value": 0.061
    }
  ]
}
```
