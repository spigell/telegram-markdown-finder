# Synopsys

Simple Telegram bot for learning purposes. Return part of markdown file by given text in header. 


# Configuration

 - TelegramBotToken - token for your bot (you can find it by chatting with BotFather).
 - PastePath - dictionary. Places where md files exists. For example

 ```
 "PastePath": { 
    "paste":  "./paste.md",
    "passwords":  "http://127.0.0.1:80/passwords.md"
 }
 ```

Refs
https://godoc.org/github.com/go-telegram-bot-api/telegram-bot-api
https://ashirobokov.wordpress.com/2016/10/04/dnd-spells-telegram-bot-1/
