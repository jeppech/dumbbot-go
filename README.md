# Requirements
yt-dlp: https://github.com/yt-dlp/yt-dlp  
ffmpeg: https://ffmpeg.org/

# Download audio only, from youtube
./dl \<filename\> "\<youtube-link\>"

# Convert downloaded file to DCA
./conv \<filename\> "\<description\>"

# Example

```bash
./dl hah "https://www.youtube.com/watch?v=YaG5SAw1n0c"
[youtube] Extracting URL: https://www.youtube.com/watch?v=YaG5SAw1n0c
[youtube] YaG5SAw1n0c: Downloading webpage
[youtube] YaG5SAw1n0c: Downloading android player API JSON
[info] YaG5SAw1n0c: Downloading 1 format(s): 251
[download] Destination: hah.webm
[download] 100% of   53.80KiB in 00:00:00 at 501.79KiB/s
[ExtractAudio] Destination: hah.mp3
Deleting original file hah.webm (pass -k to keep)

./conv hah "hah *****"
Input #0, mp3, from 'hah.mp3':
  Metadata:
    encoder         : Lavf59.27.100
  Duration: 00:00:03.19, start: 0.023021, bitrate: 232 kb/s
  Stream #0:0: Audio: mp3, 48000 Hz, stereo, fltp, 232 kb/s
    Metadata:
      encoder         : Lavc59.37
Stream mapping:
  Stream #0:0 -> #0:0 (mp3 (mp3float) -> pcm_s16le (native))
Press [q] to stop, [?] for help
Output #0, s16le, to 'pipe:1':
  Metadata:
    encoder         : Lavf59.27.100
  Stream #0:0: Audio: pcm_s16le, 48000 Hz, stereo, s16, 1536 kb/s
    Metadata:
      encoder         : Lavc59.37.100 pcm_s16le
size=     590kB time=00:00:03.14 bitrate=1536.0kbits/s speed= 139x
video:0kB audio:590kB subtitle:0kB other streams:0kB global headers:0kB muxing overhead: 0.000000%

Add this line to the main function:
add_track("hah", "hah.dca", "hah *****")
```
