from socket import gethostname
from flask import Flask, render_template
app = Flask(__name__)

@app.route("/")
def hello_goose():
    hostname = gethostname()
    return render_template('index.html.j2', hostname=hostname)
