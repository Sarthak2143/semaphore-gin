# semaphore-gin

[![forthebadge made-with-go](http://ForTheBadge.com/images/badges/made-with-go.svg)](https://go.dev/)

Microservice to display list of articles and deatils.

**Stack used:**

- Gin (Golang Web Framework)
- Semaphore (CI/CD)

This application is a simple article manager. It will be able to show articles in HTML, JSON and XML as needed.

## Endpoints

- `/`: `GET` Home page of app, displaying all articles and links to them.

- `/articles`
    - `/view/:id`: `GET` View articles by there unique ID, can render in HTML, JSON and XML. View [examples](#examples) for more info.

## Examples

Run the application by running

```bash
go run .
```
in project directory and visit `localhost:8080`.

**Using `curl`**

- Get all articles, visiting home page.

1. HTML

```bash
curl http://localhost:8080/
```

Output:

```html
<html>
...
...
<a href="/article/view/1">
    <h2> Article I </h2>
</a>
<p> Artcle I Content </p>
<a href="/article/view/2">
    <h2> Article II </h2>
</a>
<p> Artcle II Content </p>
...
...
</html>
```

2. JSON

```bash
curl -X GET -H "Accept: application/json" http://localhost:8080
```
Ouput:

```json
[
    {
        "id": 1,
        "title": "Article I",
        "content": "Artcle I Content"
    },
    {
        "id": 2,
        "title": "Article II",
        "content": "Artcle II Content"
    }
]
```

3. XML

```bash
curl -X GET -H "Accept: application/xml" http://localhost:8080
```
Output:

```xml
<article><ID>1</ID><Title>Article I</Title><Content>Artcle I Content</Content></article><article><ID>2</ID><Title>Article II</Title><Content>Artcle II Content</Content></article>...
```

- Get a particular article using its ID.

1. HTML

```bash
curl http://localhost:8080/article/view/1
```

Output:

```html
<html>
...
<title> Article I </title>
...
...
<h1> Article I </h1>
<p> Article I Content </p>
...
...
</html>
```

2. JSON

```bash
curl -X GET -H "Accept: application/json" http://localhost:8080/article/view/1
```

> Note: Work is in progress

3. XML

```bash
curl -X GET -H "Accept: application/xml" http://localhost:8080/article/view/1
```

> Note: Work is in progress

---

## Installation

## Prerequisites

- **An installation of Go 1.16 or later**. For installation instructions, see [Installing Go.](https://go.dev/doc/install)

- **A tool to edit your code**. Any text editor you have will work fine.

- **A command terminal**. Go works well using any terminal on Linux and Mac, and on PowerShell or cmd in Windows.

- **The curl tool**. On Linux and Mac, this should already be installed. On Windows, it’s included on Windows 10 Insider build 17063 and later. For earlier Windows versions, you might need to install it. For more, see [Tar and Curl Come to Windows.](https://docs.microsoft.com/en-us/virtualization/community/team-blog/2017/20171219-tar-and-curl-come-to-windows)

### Setup

```bash
git clone https://github.com/Sarthak2143/semaphore-gin
cd semaphore-gin/
go get .
```

---
