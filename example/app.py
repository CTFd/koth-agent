from flask import Flask, request, redirect

app = Flask(__name__)


@app.route("/", methods=["GET", "POST"])
def index():
    if request.method == "GET":
        with open("owner.txt", "r+") as f:
            contents = f.read()
            return f"""
            <p>Contents of owner.txt: {contents}</p>
            <form method="POST">
                <label>Update owner.txt</label><br>
                <input type="text" name="identifier">
                <input type="submit">
            </form>
            """
    elif request.method == "POST":
        with open("owner.txt", "w+") as f:
            identifier = request.form["identifier"]
            f.write(identifier)
            return redirect("/")


if __name__ == "__main__":
    app.run(debug=True, host="0.0.0.0", threaded=True)
