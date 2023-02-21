# Requirements
yt-dlp: https://github.com/yt-dlp/yt-dlp  
ffmpeg: https://ffmpeg.org/

# Download audio only, from youtube
./dl filename "description" "youtube-link"

# Example

```bash
./dl hah "hah ***" "https://www.youtube.com/watch?v=YaG5SAw1n0c"
[youtube] Extracting URL: https://www.youtube.com/watch?v=YaG5SAw1n0c
[youtube] YaG5SAw1n0c: Downloading webpage
[youtube] YaG5SAw1n0c: Downloading android player API JSON
[info] YaG5SAw1n0c: Downloading 1 format(s): 251
[download] Destination: hah.webm
[download] 100% of   53.80KiB in 00:00:00 at 3.62MiB/s
[ExtractAudio] Destination: hah.mp3
Deleting original file hah.webm (pass -k to keep)
----
Add this line to the main function:
add_track("hah", "hah.dca", "hah ***")
----
```
