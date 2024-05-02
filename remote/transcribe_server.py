from flask import Flask, request, jsonify
from vosk import Model, KaldiRecognizer
from pydub import AudioSegment
import json

app = Flask(__name__)
model = Model("./models")

@app.route('/transcribe', methods=['POST'])
def transcribe():
    if 'file' not in request.files:
        return "No file part", 400
    file = request.files['file']
    if file.filename == '':
        return "No selected file", 400

    try:
        sound = AudioSegment.from_file(file, format="mp3")

        rec = KaldiRecognizer(model, sound.frame_rate) 

        for chunk in sound[::1000]:  
            rec.AcceptWaveform(chunk.raw_data)  

        result = json.loads(rec.FinalResult())
        return jsonify(result)
    except Exception as e:
        return str(e), 500

@app.route('/checkhealth', methods=['GET'])
def check_health():
    return jsonify({
        "httpStatusCode": 200,
        "status": "Online",
        "message": "API is running smoothly"
    })
if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
