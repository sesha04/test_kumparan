# test_kumparan

A simple Web service to create and fetch articles

The article tables is stored in mysql with these columns :
  1.  id INT
  2.  author STRING
  3.  title STRING
  4.  body  Text
  5.  created_at  DATETIME
  6.  updated_at  DATETIME
  
There is 2 endpoint in this service
  1. [POST]  `/articles` to create articles
  2. [GET]  `/articles`to get list of articles, sorted by the newest first. this endpoint has two optional query parameters :
    - query: to search by title or word in body
    - auhor : to search by author's name
    
This code is create by following the structure on `https://github.com/bxcodec/go-clean-arch`.
get articles response is cached using `https://github.com/patrickmn/go-cache`
