# Media Pipeline

How OpenClaw processes inbound and outbound media: images, audio, video, and documents.

```mermaid
flowchart TD
    InboundMedia["Inbound Media\n(from channel API)"]
    MediaStore["Media Store\nsrc/media/store.ts"]
    MimeSniff["MIME Sniffing\nsrc/media/sniff-mime-from-base64.ts"]
    ImageOps["Image Operations\nsrc/media/image-ops.ts"]
    AudioProcess["Audio Processing\nsrc/media/audio.ts"]
    FFmpegExec["FFmpeg Executor\nsrc/media/ffmpeg-exec.ts"]
    FFmpegLimits["FFmpeg Limits\nsrc/media/ffmpeg-limits.ts"]
    PDFExtract["PDF Text Extraction\nsrc/media/pdf-extract.ts"]
    Base64["Base64 Encode/Decode\nsrc/media/base64.ts"]
    PNGEncode["PNG Encoder\nsrc/media/png-encode.ts"]
    TempFiles["Temp File Manager\nsrc/media/temp-files.ts"]
    MediaServer["Media HTTP Server\nsrc/media/server.ts"]
    OutboundAttach["Outbound Attachment\nsrc/media/outbound-attachment.ts"]
    InputFiles["Input File Resolution\nsrc/media/input-files.ts"]
    InboundPolicy["Inbound Path Policy\nsrc/media/inbound-path-policy.ts"]

    InboundMedia --> MimeSniff
    MimeSniff --> MediaStore
    MediaStore --> ImageOps
    MediaStore --> AudioProcess
    MediaStore --> PDFExtract
    ImageOps --> PNGEncode
    AudioProcess --> FFmpegExec
    FFmpegExec --> FFmpegLimits
    FFmpegExec --> TempFiles
    PDFExtract --> TempFiles
    Base64 --> MediaStore
    TempFiles --> MediaServer
    InboundPolicy --> InputFiles
    InputFiles --> Base64
    MediaStore --> OutboundAttach
    MediaServer --> OutboundAttach
```

## Image Processing Flow

```mermaid
sequenceDiagram
    participant Channel
    participant Store as media/store.ts
    participant ImageOps as media/image-ops.ts
    participant PNG as media/png-encode.ts
    participant Agent as Agent

    Channel->>Store: store inbound image bytes
    Store->>ImageOps: resize / convert if needed
    ImageOps->>PNG: encode to PNG
    PNG-->>Store: PNG bytes
    Store-->>Agent: ImageContent { type, data, mediaType }
```

## Audio Processing Flow

```mermaid
sequenceDiagram
    participant Channel
    participant AudioTS as media/audio.ts
    participant FFmpeg as media/ffmpeg-exec.ts
    participant TempDir as media/temp-files.ts
    participant Tags as media/audio-tags.ts
    participant Agent

    Channel->>AudioTS: inbound voice/audio file
    AudioTS->>TempDir: create temp dir
    AudioTS->>FFmpeg: transcode to compatible format
    FFmpeg->>FFmpegLimits: check file size / duration limits
    FFmpeg-->>AudioTS: transcoded file path
    AudioTS->>Tags: read audio tags (metadata)
    Tags-->>AudioTS: title, duration, codec info
    AudioTS-->>Agent: audio content ready
```

## Media Types and MIME Support

```mermaid
graph LR
    MediaTypes["Supported Media Types\nsrc/media/constants.ts"]
    MediaTypes --> Images["Images\nJPEG, PNG, GIF, WEBP"]
    MediaTypes --> Audio["Audio\nMP3, OGG, OGG/Opus, M4A, WAV"]
    MediaTypes --> Video["Video\nMP4 (via ffmpeg)"]
    MediaTypes --> Docs["Documents\nPDF (text extraction)"]
    MediaTypes --> Files["Files\ngeneric attachment passthrough"]
```

## Key Files

| File | Role |
|------|------|
| `src/media/store.ts` | Central media object store |
| `src/media/audio.ts` | Audio download, transcode, prepare |
| `src/media/ffmpeg-exec.ts` | FFmpeg subprocess wrapper |
| `src/media/ffmpeg-limits.ts` | Size/duration limits enforcement |
| `src/media/audio-tags.ts` | Read ID3/audio metadata tags |
| `src/media/image-ops.ts` | Image resize, crop, convert |
| `src/media/png-encode.ts` | Raw-to-PNG encoding |
| `src/media/pdf-extract.ts` | Extract text content from PDFs |
| `src/media/base64.ts` | Base64 encode/decode helpers |
| `src/media/sniff-mime-from-base64.ts` | Detect MIME type from base64 header |
| `src/media/temp-files.ts` | Temp file lifecycle management |
| `src/media/server.ts` | Serve media files over HTTP (for agent access) |
| `src/media/outbound-attachment.ts` | Prepare attachment for channel send |
| `src/media/input-files.ts` | Resolve file paths from user input |
| `src/media/inbound-path-policy.ts` | Validate allowed inbound media paths |
