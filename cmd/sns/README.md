# SNS Message Sender

A command-line tool to send SNS (Social Network Service) messages from your
terminal. I wanted to build this because I don't want to leave my terminal.

### Features

- Interactive REPL (Read-Eval-Print Loop) for sending SNS messages
- Switch between modes (currently only `kakao` is available)
- Built-in help and exit commands (different for each mode)
- Available modes:
  - `KakaoTalk` (login, send, list friends)

### How to Use

Before using the tool, you need to set `API keys` for each SNS platform to
`.env` file in the root directory.

```bash
KAKAO_API_KEY=your_api_key
```

Then, you can run the tool with the following command:

```bash
go run ./cmd/sns
```

#### Available commands (Default mode)

- `help` — Show all commands and usage
- `exit` — Exit the tool
- `mode kakao` — Switch to KakaoTalk mode

#### KakaoTalk Mode

- After switching to `kakao` mode, you can use KakaoTalk-specific commands (type
  `help` to see them).
- `login` — Login to KakaoTalk
- `send` — Send a message to a friend or yourself
- `list friends` — List your friends

## Notes

- Sadly, `Kakao Talk` API does not support getting all friends from a user. It
  only allows users that have used this APP's API and agreed to share their
  friends.

- More SNS platforms may be added in the future.

### Example

mode: `KAKAO`

```bash
❯ go run ./cmd/sns
=========== SNS-Sender ==========
Sending your SNS message on terminal.
SNS > mode
Current mode: default
Usage: mode [mode]
Available modes:
  - kakao (Kakao Talk)
Set mode to default
SNS > mode kakao
Set mode to kakao
KAKAO > ls
Error: not logged in
KAKAO > login
Login to Kakao Talk...
Please open the following URL in your browser:
https://kauth.kakao.com/oauth/authorize?client_id=(your_api_key)&redirect_uri=...
Login to Kakao Talk successfully!
KAKAO > ls
Friends:
0. me
1. (your team member nickname)
KAKAO > send 0 message
Message sent to me
KAKAO > exit
Exiting...
❯
```
