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

// Buffer.alloc(8) creates a new Buffer instance in Node.js with an allocated memory size of 8 bytes. The allocated memory is initialized to zero. 
// The returned buffer can be used to store binary data such as raw bytes, integers, or floats.

// The Buffer class is a built-in Node.js class that allows you to work with binary data in a way similar to working with arrays in JavaScript. 
// The alloc() method is a static method of the Buffer class that creates a new buffer instance with a specified length.

// In the example, Buffer.alloc(8) creates a new buffer with 8 bytes of memory. 
// If you were to log the result of this call to the console, it would output a Buffer object with 8 bytes of memory, all initialized to zero. 
// Here is an example:
// const buf = Buffer.alloc(8);
// console.log(buf); // <Buffer 00 00 00 00 00 00 00 00>
// This buffer can then be used to read from and write to binary data in a Node.js application. 
// For example, you could use the writeInt32LE() method of the buffer to write a 32-bit integer value to the first four bytes of the buffer:

// buf.writeInt32LE(42, 0);
// console.log(buf); // <Buffer 2a 00 00 00 00 00 00 00>
// This would write the value 42 as a little-endian 32-bit integer to the first four bytes of the buffer, 
// resulting in the buffer containing the hexadecimal bytes 2a 00 00 00 00 00 00 00.