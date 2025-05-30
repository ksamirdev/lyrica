## Lyrics Generator

The Lyrics Generator is a simple Go application designed to effortlessly create time-synced lyrics.

![App Preview](/assets/app.png)
_Music: [Please Please Please by Sabrina Carpenter](https://open.spotify.com/track/5N3hjp1WNayUPZrA8kJmJP)_

### Tech Stack

- **Frontend:** HTML, Alphine.JS, Tailwind CSS
- **Backend:** Go ([go chi](https://github.com/go-chi/chi))

### Features

- **Simple Interface:** Easy-to-use application for generating synchronized lyrics.
- **Efficient:** Quickly creates time-synced lyrics for various uses.
- **Export to LRC:** Export generated lyrics as LRC files.
- **Dashboard:** Interactive dashboard available on specified port (configured via `.env` file).

### Installation

You can download the binaries from the [GitHub Releases](https://github.com/ksamirdev/lyrica/releases) page.

### Usage

1. **Download the Binary:** Obtain the appropriate binary for your operating system from the Releases page.
2. **Run the Application:**
   ```bash
   ./lyrica-[os-arch]
   ```
3. **Access the Dashboard:**
   - Navigate to `http://localhost:<port>` in your browser (port specified in `.env` file).
   - Import a track to the dashboard.
   - Play the track and insert lyrics along with current timestamps.

### Contributing

Contributions are welcome! If you find any issues or have suggestions, please open an issue or a pull request on the [GitHub repository](https://github.com/ksamirdev/lyrica).

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
