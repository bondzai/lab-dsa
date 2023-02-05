const axios = require('axios');

exports.getUser = (async (req, res) => {
    try {
        const response = await axios.get('https://script.google.com/macros/s/AKfycbzuE4HExzdaIyEqETAcenS9VVaQ1YIbxQjBzAo93lQeSf8ZbhTujHXNCUd1CQcMFp2-DQ/exec?action=readData');
        return response.data;
    } catch (error) {
        console.error(error);
        res.status(500).send('server internal error');
    };
});

exports.approve = (async (req, res) => {
    try {
        await axios.post('https://script.google.com/macros/s/AKfycbzuE4HExzdaIyEqETAcenS9VVaQ1YIbxQjBzAo93lQeSf8ZbhTujHXNCUd1CQcMFp2-DQ/exec?action=approve', {
        id: req.body.id
    })
        return true;
    } catch (error) {
        console.error(error);
        res.status(500).send('server internal error');
    };    

});

const nodemailer = require("nodemailer");
exports.sentEmail = async () => {
    const mailList = [
        'introbond.dev@gmail.com',
        'napat.s@swiftdynamics.co.th'];
    try {
        let transporter = nodemailer.createTransport({
            service: "gmail",
            auth: {
                user: "introbond.node@gmail.com",
                pass: process.env.NODE_MAILER_PASSWORD,
            }
    });
    
    let info = await transporter.sendMail({
        from: '"introbond.node@gmail.com',
        to: mailList,
        subject: "Backend interview: Puritat Chamart",
        text: `Dear Team (From nodejs)
            As a coding interview, here is the result of testing no 4.
            Thanks for reading.

            Regard,
            Puritat Chamart`
    });
    console.log("Message sent: %s", info.messageId);
    return true
    } catch (error) {
        console.log(error)
        return false
    }
};
