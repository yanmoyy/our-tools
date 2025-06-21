# PPT CRAWLING

## Purpose

- I have to download some PPT files of praise for our church every week.
- So I want to do this easily and fastly a PPT file by crawling from Google.

  - just type some query like **number** or **title** .
  - and want to search multiple PPT files.

## How to setup

1. Create a 'env' file in the root directory of this project.

   - path : ./.env

2. Set the following environment variables in env file.

```
GOOGLE_API_KEY=your_google_api_key
CSE_API_KEY=your_cse_api_key
```

3. Run the command.

   - 3-1. just **go run** command

     - `go run ./cmd/ppt-crawling/...`

   - 3-2. build the binary file and run it.

     - `go build -o ppt-crawling ./cmd/ppt-crawling`
     - `./ppt-crawling`

## Commands

### 1. list

- display all of ppt download urls that you have searched.

### 2. help

- show all of command's information.

### 3. search <start? | num> / s <start?| num>

- search ppt download url by query.

- start argument is optional.

  - start argument is number and it will search 10 ppt from start number

    - it is not searched exactly 10 ppt. because it will happen error.
    - but it will show how many ppt have been searched.

  - if you type start argument, it will search ppt download url from start number.
  - other wise, it will search ppt from default start number(1).

  - it is not more than 90. (because of google api limit)

- to search ppt download url by query, you type input with newline
- that is, you can type multiple lines.
- and if you type nothing, it will search all of ppt download url about your queries.

- to quit title state, type **exit** command.

  - example

    ```
    ppt-crawler > search
    search > 페이커

    ===================================
    Starting Searching... (페이커)

    Found 0 urls

    no url found in 페이커
    ===================================
    search >
    search > exit
    ppt-crawler >
    ```

- **note** : if you type title with korea language, it will happen error if you doesn't delete the space.

### 4. download <path?>

- download all of ppt download url that you have searched.

- if you type path argument, it will download ppt files that you have searched to the path.
- otherwise, it will download ppt files to the root directory.

### 5. exit

- exit ppt-crawling cli.
