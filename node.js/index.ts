import * as express from "express"
import {Request, Response} from "express";
import {IncomingForm} from "formidable";

const PORT = 3579
const FORM_IMAGE_KEY = "image"

const app = express()

app.post("/upload", (req: Request, res: Response) => {
  const form = new IncomingForm()
  form.parse(req, (err, _, files) => {
    if (err || !files[FORM_IMAGE_KEY] || files[FORM_IMAGE_KEY].size === 0) {
      res.status(400).send({message: 'missing file'})
      return
    }
    handle_file(files[FORM_IMAGE_KEY])
    res.sendStatus(204)
  })
})

function handle_file(file) {
  // do something with the file
}

app.listen(PORT, function () {
  console.log(`listening on http://localhost:${PORT}`)
})

