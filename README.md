# better-spotify-wrapped

A simple CLI tool that can parse extended lifetime listening data from Spotify [which can be downloaded here](https://www.spotify.com/us/account/privacy/). The data is parsed from raw listening history data to a JSON file containing aggregated listening data, which can then be queried for interesting insights into your listening.

##### Disclaimer: This tool was my first ever Go project. Do not expect it to be perfect (or to work for that matter).

---

## ✨ Features

- Parse extended lifetime listening data from Spotify
- Aggregate listening data into a JSON file
- Query the JSON file for interesting insights

---

## 📦 Installation

1. Make sure you have [Go](https://golang.org/doc/install) installed.
2. Clone the repository.
3. build & run (TODO write this when the project is more than a hello world)

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

---

📄 License
MIT © Birk Nordbo
