# Gosound-api
### API Features
* **Audio to Text Transcription**: This API allows for the conversion of audio files into accurate text transcriptions, suitable for a variety of audio content including podcasts, meetings, and interviews.
* **Text to Audio Conversion**: It also offers the capability to convert text into natural-sounding audio. This is ideal for producing audiobooks, voiceovers for videos, and other audio content, with a selection of voices and languages available.
* **All-in-One Platform**: Combining both transcription and text-to-audio conversion, the API provides a streamlined solution for switching between audio and text formats, facilitating content creation and accessibility.
* **Speed**: The API is optimized for fast processing, ensuring that conversions are completed swiftly to enhance user productivity.
* **Privacy**: With a commitment to user privacy, the API ensures that all data is handled with the utmost confidentiality and security, protecting user information during each interaction.
* **Unlimited Usage and No Fees per Query**: The API imposes no usage limits and does not charge fees per query, allowing users to freely access its features as often as needed without additional costs.
### Project Support Features
* Users can signup and login to their accounts
* Public (non-authenticated) users can access all causes on the platform
* Authenticated users can access all causes as well as create a new cause, edit their created cause and also delete what they've created.
### Installation Local
* Clone this repositorie.
* Install a model you will want [here](https://alphacephei.com/vosk/models) and put inside 'remote/' directory
### Usage
* Run:
```shell
docker-compose up -d
```
* Connect to the API using a request with browser or Insomnia/Postman on port 4000.
### API Endpoints
| HTTP Verbs | Endpoints | Action | Format
| --- | --- | --- | --- |
| GET | /v1/gosoundapi | To retrive status of API | GET
| POST | /v1/gosoundapi/tts | Send a json text to receive an audio download | {"text": "something"} |
| POST | /v1/gosoundapi/tts | Send an audio and receive a json transcript | Multipart form: name -> audio ; value -> mp3 audio |



### Technologies Used
* [GO](https://go.dev/) 
* [Python](https://www.python.org/)

### Marketplace
soon!

### Authors
* [Marcos Vinícius](https://github.com/odmrs)
