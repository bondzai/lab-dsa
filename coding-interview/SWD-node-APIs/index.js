const app = require('./app');
const PORT = (process.env.PORT || 4001);

app.listen(PORT, (req, res) => {
    console.log(`server is running at ${PORT}`);
});