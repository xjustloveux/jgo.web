[![JGo Web](https://jgo.dev/assets/images/logo_300.svg)](https://jgo.dev/)

[![Build Status](https://github.com/xjustloveux/jgo.web/actions/workflows/cicd.yml/badge.svg)](https://github.com/xjustloveux/jgo/actions/workflows/cicd.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/xjustloveux/jgo.web)](https://goreportcard.com/report/github.com/xjustloveux/jgo.web)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/xjustloveux/jgo.web/blob/master/LICENSE)

---

* [Overview](#Overview)
* [Middlewares](#Middlewares)
* [Build](#Build)
* [Database](#Database)

# Overview

---

JGo web is a project to introduce and demo JGo.

JGo Web：[https://jgo.dev](https://jgo.dev)

JGo Project：[https://github.com/xjustloveux/jgo](https://github.com/xjustloveux/jgo)

# Middlewares

---

JGo web is built with [gin](https://github.com/gin-gonic/gin) and [angular](https://angular.io/).

# Build

---

### Step1

After you copy this project, you need to change `config/config.json` `jEnv` value to `dev` or remove `jEnv`.

If you keep `jEnv` then when you run project will load `*-dev.json`.

### Step2

You have to build a database.

For database details, please refer to [Database](#Database).

### Step3

You need to change `sql.json`(or `sql-dev.json`) file connection setting to your setting.

If you want to change log path you can change `log.json`(or `log-dev.json`)

### Step4

Install angular package.

E.g. open cmd and cd to angular folder path

```shell
npm install
```

### Step5

Run back-end go project.

```shell
go run main.go
```

### Step6

Run front-end angular project.

```shell
npm run start
```

# Database

### JOB_LOG

| Column  | Required | Type     | Size | PK    |
|---------|----------|----------|------|-------|
| NAME    | true     | nvarchar | 10   | false |
| ID      | true     | nvarchar | 10   | false |
| CT_DATE | true     | datetime |      | false |

### PKG_CONTENT

| Column      | Required | Type     | Size | PK    |
|-------------|----------|----------|------|-------|
| SEQ         | true     | int      |      | true  |
| PRO_ID      | true     | nvarchar | 10   | false |
| TITLE       | false    | nvarchar | 50   | false |
| DESCRIPTION | false    | nvarchar | 1000 | false |
| CODE        | false    | nvarchar | 1000 | false |
| NOTE        | false    | nvarchar | 500  | false |
| SORT        | false    | int      |      | false |

### SYS_PRO

| Column        | Required | Type     | Size | PK    |
|---------------|----------|----------|------|-------|
| PRO_ID        | true     | nvarchar | 10   | true  |
| PRO_URL       | false    | nvarchar | 50   | false |
| PRO_NAME      | false    | nvarchar | 50   | false |
| PARENT_PRO_ID | false    | nvarchar | 10   | false |
| SORT          | false    | int      |      | false |

### VER_HS

| Column | Required | Type     | Size | PK    |
|--------|----------|----------|------|-------|
| VER    | true     | int      |      | true  |
| NAME   | false    | nvarchar | 50   | false |

### VER_UPDATE

| Column  | Required | Type     | Size | PK    |
|---------|----------|----------|------|-------|
| SEQ     | true     | int      |      | true  |
| VER_S   | true     | int      |      | false |
| VER_E   | false    | int      |      | false |
| CONTENT | false    | nvarchar | 1000 | false |