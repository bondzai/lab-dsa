require('dotenv').config();
const express = require('express');
const morgan = require('morgan');
const app = express();

app.set('view engine', 'ejs');
app.use(express.json());
app.use(morgan('tiny'));

const data = require('./route/RawData');
app.use('/api/v1', data);

app.get('/', (req, res) => {
    res.status(200).render('index');
});

module.exports = app;