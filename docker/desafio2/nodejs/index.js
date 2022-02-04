const express = require("express");
const app = express();
app.use(express.urlencoded());
const port = 8080;
const fs = require("fs");

const config = {
  host: "mysql",
  user: "root",
  password: "secret",
  database: "nodejs",
};
const mysql = require("mysql2/promise");

let connection = null;
let file = null;

function readModuleFile(path, callback) {
  try {
    var filename = require.resolve(path);
    fs.readFile(filename, "utf8", callback);
  } catch (e) {
    callback(e);
  }
}

const getFromDB = async () => {
  const command = `SELECT name FROM nomes`;
  const result = await connection.query(command);
  return result[0].map((e) => e.name);
};

const saveToDB = (newName) => {
  const command = `INSERT INTO nomes(name) values('${newName}')`;
  connection.query(command);
};

const getContent = async () => {
  let names = await getFromDB();
  names = names.map((name) => `<li>${name}</li>`).join("<br>");
  return file.replace("<!-- INSERT LIST HERE -->", names);
};

app.get("/", async (req, res) => {
  console.log("get /");
  const content = await getContent();
  res.send(content);
});

app.post("/", async (req, res) => {
  console.log("post /");
  try {
    saveToDB(req.body.name);
    res.send(await getContent());
  } catch (e) {
    console.error("Error on db write");
    console.error(e);
    res.send(`error... <a href="/">Go back</a>`);
  }
});

app.listen(port, async () => {
  console.log("Awaiting DB load...");
  try {
    connection = await mysql.createConnection(config);
  } catch (err) {
    console.error("!!!!!! DB LOAD FAILED!");
    console.error(err);
    return;
  }
  await new Promise((res) => {
    readModuleFile("./index.html", function (err, readFile) {
      if (err) {
        console.log("!!!!!! HTML FAILED READ");
      }
      file = readFile;
      res();
    });
  });
  console.log("DB loaded");
  console.log(`Example app listening at http://localhost:${port}`);
});
