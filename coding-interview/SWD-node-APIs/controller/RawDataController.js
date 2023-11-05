const service = require('../util/service');

exports.getAllData = (async (req, res) => {
    try {
        const rawData = await service.getRawData();
        res.json({
            success: true,
            data: rawData,
        });
    } catch {
        res.status(500).send('internal server error');
    }
});

exports.getMinMaxAverage = (async (req, res) => {
    try {
        const rawData = await service.getRawData();
        const maxData = Math.max.apply(Math, rawData.map(a => a.data));
        const minData = Math.min.apply(Math, rawData.map(a => a.data));
        const averageData = rawData.reduce((total, current) => total + current.data, 0) / rawData.length;
        res.json({
            success: true,
            min: minData,
            max: maxData,
            average: averageData.toFixed(2),
        });
    } catch {
        res.status(500).send('internal server error');
    }
});

exports.getDataByRange = (async (req, res) => {
    try {
        const rangeIndex = Number(req.query.rangeIndex);
        const rawData = await service.getRawData();
        const maxData = Math.max.apply(Math, rawData.map(a => a.data));
        const minData = Math.min.apply(Math, rawData.map(a => a.data));
        const fullRange = maxData - minData;
        const subRangeNumber = Math.ceil(fullRange / 200);
        const subRangeValue = []
        let buffer = minData;
        for (let i = 0; i < subRangeNumber; i++) {
            buffer = buffer + 200;
            subRangeValue.push(buffer);
        };
        if (rangeIndex >= (subRangeValue.length - 1)) {
            res.status(400).send('no data in range')
        }
        const filter = rawData.filter((item) => {
            return (item.data >= subRangeValue[rangeIndex]) && (item.data <= subRangeValue[rangeIndex+1]);
        });
    
        res.json({
            success: true,
            filter: filter
        });
    } catch {
        res.status(500).send('internal server error');
    }
});