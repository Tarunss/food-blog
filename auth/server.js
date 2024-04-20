const bcrypt = require("bcrypt")
const bodyParser = require("body-parser");
const server = require('express')();
var cors = require('cors');
const saltRounds = 10
var username = "ssohla"
// This server is for authorization of the admin account, which will be able to edit posts on the blog site.

// We also use bcrypt in order to hash our password
server.use("/admin", cors(), bodyParser.json(), (req, res) => {
    console.log(req.body)

    // Use a parser to parse the body of the request, and check if the username and password are correct.
    res.send({
        token: 'test123'
    });
});

server.listen(8081, () => console.log("authorization API running on localhost:8081/admin"))