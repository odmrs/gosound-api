#!/bin/bash
echo "Installing vosk's models"
curl -s -o models.zip https://alphacephei.com/vosk/models/vosk-model-pt-fb-v0.1.1-20220516_2113.zip
echo "Download completed!"

unzip models.zip -d models/
echo "Unzip completed!"

rm -rf models.zip
echo "Cleanup completed!"

mv ./models/ ./remote/

cd remote/models/vosk-model-* 
mv * ../
rm -rf remote/models/vosk-model-*

clear
echo "Done, now run the docker-compose"
