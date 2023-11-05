const axios = require('axios');
exports.getRawData = (async (req, res) => {
    try {
        const response = await axios.get('http://3.1.189.234:8091/data/ttntest');
        return response.data;
    } catch (error) {
        console.error(error);
        res.status(500).send('server internal error');
    };
});