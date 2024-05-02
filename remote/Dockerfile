


# Use an official Python runtime as a parent image
FROM python:3.12-slim as builder

WORKDIR /app

FROM python:3.12-slim as builder

WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    libsndfile1 \
    ffmpeg \
    libdbus-1-dev \
    pkg-config \
    meson \
    ninja-build \
    # Add this line to install GLib development libraries
    libglib2.0-dev \
    dbus-x11 && \
    rm -rf /var/lib/apt/lists/*

COPY requirements.txt .

RUN pip install --no-cache-dir -r requirements.txt

COPY . .

EXPOSE 5000

CMD ["python", "transcribe_server.py"]

