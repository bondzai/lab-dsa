const express = require('express');
const router = express.Router();

const { getAllData, getDataByRange, getMinMaxAverage } = require('../controller/RawDataController');
router.route('/getAllData').get(getAllData);
router.route('/getMinMaxAverage').get(getMinMaxAverage);
router.route('/getDataByRange').get(getDataByRange);


module.exports = router;