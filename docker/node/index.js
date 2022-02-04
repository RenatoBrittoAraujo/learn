const express = require("express");
const app = express();
const port = 3000;

app.get("/", (req, res) => {
  res.send("<h1>I HAVE NOT just shit and cum </h1>");
});

app.listen(port, () => {
  console.log("SERVIDOR RODANDO NA PORTA", port);
});
