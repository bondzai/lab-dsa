const url = 'https://notify-api.line.me/api/notify'
const token = '8EQylhgR1lzr6Is7cgKExGeo1gerETSW45ihp3LrngQ';

msg.headers = {
    'content-type': 'application/x-www-form-urlencoded',
    'Authorization': 'Bearer ' + token
};
msg.payload = {
    "message": "Hello from node-red"
};

return msg;