# go-crawler

WIP
Components

- manager: All in one, command line launcher
- api: A HTTP API server
- fetcher: A fetch web site worker
- scheduler: A task schedule controller

## a crawler task spec

```json
{
  "id": "d9c5d709-610a-4712-83ec-7e33bfb09105",  // task id 
  "version": 0.1,
  "created_time": 0,
  "modified_time": 0,
  "deleted_time": 0,
  "url": "", // crawler
  "path": "",
  "query": "",
  "raw_html": "",
  "lastest_crawled_time": 0,
  "crawled_times": [],
  "crontab": "",
  "cooling_time": "",
  "selectors": [
    {
      "css_select": "",
      "xpath": "",
      "result": ""
    },
    {
      "css_select": "",
      "xpath": "",
      "result": ""
    }
  ]
}
```
