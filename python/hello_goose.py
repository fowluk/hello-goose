from socket import gethostname
from flask import Flask, render_template
import os
app = Flask(__name__)

@app.route("/")
def hello_goose():

    instanceId = os.environ['CF_INSTANCE_GUID']
    instanceIndex = os.environ['CF_INSTANCE_INDEX']
    return render_template('index.html.j2', instance_id=instanceId, instance_index=instanceIndex)
