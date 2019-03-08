const express = require('express')
const instance_id = process.env.INSTANCE_GUID || "No instance ID"
const instance_index = process.env.INSTANCE_INDEX || "No instance INDEX"
const port = process.env.PORT || 3000

let app = express();

app.use(express.static('views/static'))
app.engine('pug', require('pug').__express)
app.set("view engine", "pug");
app.set("views", "./views");

app.get('/', function (req, res) {
    res.render('index.pug',{instance_id: instance_id, instance_index: instance_index})
})

app.listen(port, () => console.log(`Hello-Goose app listening on port ${port}!`))
