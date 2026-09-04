#!/usr/bin/env python3

import os
import json

ROOT = os.path.expanduser("~/storage/shared")
OUTPUT = os.path.expanduser("./audio_files.json")

AUDIO_EXTENSIONS = {
    ".mp3",
    ".m4a",
    ".flac",
    ".wav",
    ".ogg",
    ".opus",
    ".aac",
    ".wma",
    ".amr",
    ".aiff",
    ".aif",
}

audio_files = []

for root, dirs, files in os.walk(ROOT):
    for filename in files:
        ext = os.path.splitext(filename)[1].lower()

        if ext in AUDIO_EXTENSIONS:
#            full_path = "/storage/emulated/0" + os.path.relpath(os.path.join(root, filename),ROOT)
            full_path = os.path.join(root, filename)
            audio_files.append(full_path)
            print(f"found: {full_path}")

audio_files.sort()

with open(OUTPUT, "w", encoding="utf-8") as f:
    json.dump(audio_files, f, indent=2, ensure_ascii=False)

print(f"Found {len(audio_files)} audio files")
print(f"Saved to: {OUTPUT}")
