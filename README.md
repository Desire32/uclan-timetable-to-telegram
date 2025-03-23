<p align="center">
  <img src="https://github.com/user-attachments/assets/e1aa2fc0-8a16-47be-ba1f-6e49f32f6a70" alt="Frame 1" />
</p>


This project does not collect any of your data, all code is available to view to certify. This project is just an implementation and not a finished solution and was done only in research purposes.

## Getting started

1. Clone repo
   
```bash
   git clone https://github.com/Desire32/uclan-timetable-telegram.git
   cd uclan-timetable-telegram
```

2. Create .env file
   
```bash
   touch internal/config/.env
```
   
To get started, an .env file with the following data must be imported into the project:

## Mandatory:

University account:
- Email ( university mail for login)
  
```bash
EMAIL="your_uclan_email"
```
- Password ( university password)
```bash
PASSWORD="your_uclan_password"
```
  
Timetable:
- Password ( password to log in to Timetable, mail must remain the same)
  
```bash
TIMETABLE_PASS="your_timetable_pass"
```

Telegram:
- Bot Token ( prepare a token for your Telegram bot to link it to the system)
```bash
BOT_TOKEN="your_telegram_bot_token"
```

## Optional:

Browser:
- Preferred browser ( chromedp will use google chrome by default, but you can also specify your own)
```bash
BROWSER_PATH="your_browser"
```

Mongo:
- MongoDB Url (to send fetched data to your cluster in MongoDB, use the format to work with go)
```bash
MONGODB_URI="your_mongodb_url_go"
```

## Screenshots
Initial launch:

![Screenshot 2025-02-22 at 15 44 01](https://github.com/user-attachments/assets/713a9985-11b8-4009-8a45-ad62588bd7d2)

Bot interaction:

To interact with the bot, use "/menu" command.

![Screenshot 2025-02-22 at 15 54 51 1](https://github.com/user-attachments/assets/72b02227-761c-4206-b146-9c868c2d043a)


## License

The project is protected by MIT License.
