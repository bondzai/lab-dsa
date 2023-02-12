const sql = require('mssql');

const config = {
    user: 'sa',
    password: '12345678',
    server: 'localhost', 
    database: 'master',
    port: 1433,
    options: {
        trustServerCertificate: true
    }
};

sql.connect(config, function (err) {
    if (err) {
        console.log(err);
        return;
    }

    console.log("Successfully connected to SQL server!");
});
