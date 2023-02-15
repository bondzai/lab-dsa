//write in node-red function

// Input message contains an array of 16-bit Modbus registers
const registers = msg.payload;

// Combine 4 registers into a 64-bit buffer (big-endian byte order)
const buffer = Buffer.alloc(8);
buffer.writeUInt16BE(registers[0], 0);
buffer.writeUInt16BE(registers[1], 2);
buffer.writeUInt16BE(registers[2], 4);
buffer.writeUInt16BE(registers[3], 6);

// Convert buffer to double-precision floating point value (big-endian byte order)
const value = buffer.readDoubleBE(0);

// Output message contains the converted value
msg.payload = value;
return msg;
