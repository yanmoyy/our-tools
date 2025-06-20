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

### 3. title

- search ppt download url by title.

- this command target specific site.

  - `https://cwy0675.tistory.com/entry/`

- to search ppt download url by title, you type input with newline
- that is, you can type multiple lines.
- and if you type nothing, it will search all of ppt download url about your titles.

- to quit title state, type **exit** command.

  - example

    ```
    ppt-crawler > title
    title > 나의 힘이 되신 여호와여

    Processing 나의 힘이 되신 여호와여
    title >
    title > exit
    ppt-crawler >
    ```

### 4. num

- search ppt download url by number.

- this command target specific site.

  - `https://lifestoryteller.tistory.com/`

- to search ppt download url by number, you type input with newline
- that is, you can type multiple lines.
- and if you type nothing, it will search all of ppt download url about your numbers.

- to quit num state, type **exit** command.

  - example

    ```
    ppt-crawler > num
    num > 150

    Processing 150
    num >

    num > exit
    ppt-crawler >
    ```

### 5. download <path?>

- download all of ppt download url that you have searched.

- if you type path argument, it will download ppt files that you have searched to the path.
- otherwise, it will download ppt files to the root directory.

### 6. exit

- exit ppt-crawling cli.
