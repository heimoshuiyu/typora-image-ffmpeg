# typora-image-ffmpeg

`typora-image-ffmpeg` is a simple photo compress tool, with `typora` and `ffmpeg`.

Basically, it receive source photo, using `ffmpeg` to compress into `jpg` format, and <strong style="color: red;">DELETE</strong> source photo.

(Lossy) compress can significant reduce the size of photo, at the same time maintaining a high quality that is indistinguishable to the naked eye.

## Install

### Install from source code

1. Clone this repository `git clone https://github.com/heimoshuiyu/typora-image-ffmpeg`
2. Change directory into this repository `cd typora-image-ffmpeg`
3. Compile `make` or `go build`
4. Install into system path `make install` or `sudo mv typora-image-ffmpeg /usr/bin/`

### Configure typora

Open `typora` - `File` - `Perferences` - `Image` - `When Insert` choose `Upload Image`

In `Image Uploader` choose `Custom command`，and type in `typora-image-ffmpeg`

Then you can click `Test uploader` for checking.

## Contributions welcome

Only few lines of simple source code, welcome any new idea!