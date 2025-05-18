# better-spotify-wrapped

A simple CLI tool that can parse extended lifetime listening data from Spotify [which can be downloaded here](https://www.spotify.com/us/account/privacy/). The data is parsed from raw listening history data to JSON files containing aggregated listening data, which can then be queried for interesting insights into your listening.

##### Disclaimer: This tool was my first ever Go project. Do not expect it to be perfect (or to work for that matter).

---

## ✨ Features

- Parse extended lifetime listening data from Spotify
- Aggregate listening data into JSON files
- Query the JSON files for interesting insights

---

## 📦 Installation

1. Make sure you have [Go](https://golang.org/doc/install) installed.
2. Clone the repository.
3. Build into an executable.
4. Run the executable to generate the data folder.
5. Place Spotify lifetime listening files in the data folder.
6. Run the parse command to generate aggregated data.
7. Run commands for interesting insights.

---

## 🧠 Usage

```bash
wrapped [command] [options]
```

🔧 Available Commands

- `parse` - Parse raw Spotify listening data to a JSON file
- `favorite <artist|track>` - Show information about your favorite artist or track
- `top <artist|track> <amount>` - Show the top tracks or artists by listening time
- `hate <artist|track>` - Show information about your least favorite artist or track by listening time
- `count <artist|track>` - Show the number of different artists or tracks you have listened to
- `help` - Show help message

---

📄 License
MIT © Birk Nordbo
