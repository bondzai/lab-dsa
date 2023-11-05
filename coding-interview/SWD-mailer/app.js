require('dotenv').config();
const express = require('express');
const app = express();
const morgan = require('morgan');
const service = require('./util/service');
const bcrypt = require('bcryptjs');

app.set('view engine', 'ejs');
app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use(morgan('tiny'));

app.get('/', (req, res) => {
    res.status(200).render('index');
})

app.post("/login", async (req, res) => {
    const { id, password } = req.body;
    const users = await service.getUser();
    const existingUser = users.find(user => user.id === id);
    if (!existingUser) {
        res.status(400).send('no user in database');
    };
    if (!await bcrypt.compare(password, existingUser.password)) {
        res.status(400).send('incorrect password');
    };
    res.render('approve', { loggedinID: existingUser.id, loggedinName: existingUser.name, loggedinApproveStatus: existingUser.approve_status });
});

app.post("/approve", async (req, res) => {
    const { id, name, approve_status } = req.body;
    await service.approve(req, res);
    const users = await service.getUser();
    const notApprove = await users.find(user => user.approve_status == false);
    console.log(typeof notApprove);
    if (!(typeof notApprove === "undefined")) {
        return res.send("wait for another approval");
    };
    await service.sentEmail();
    res.status(201).send("sent email successfully");
});

module.exports = app;